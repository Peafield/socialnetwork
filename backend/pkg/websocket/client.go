// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	chatcontrollers "socialnetwork/pkg/controllers/ChatControllers"
	followercontrollers "socialnetwork/pkg/controllers/FollowerControllers"
	usercontrollers "socialnetwork/pkg/controllers/UserControllers"
	"socialnetwork/pkg/db/dbstatements"
	"socialnetwork/pkg/db/dbutils"
	"socialnetwork/pkg/middleware"
	"socialnetwork/pkg/models/dbmodels"
	"socialnetwork/pkg/models/readwritemodels"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type ReadMessage struct {
	Type string                 `json:"type"`
	Info map[string]interface{} `json:"info"`
}
type WriteMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type BasicUserInfo struct {
	UUID            string
	Name            string
	LoggedInStatus  int
	LastMessageTime time.Time
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte

	UserID   string
	Username string
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))

		var msg ReadMessage
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Println(err)
			continue
		}

		switch msg.Type {
		case "open_chat":
			err := handleOpenChat(msg, c)
			if err != nil {
				log.Printf("error: %v", err)
			}
			break
		case "private_message":
			err := handlePrivateMessage(msg, c)
			if err != nil {
				log.Printf("error: %v", err)
			}
		default:
			log.Println(msg)
			break
		}

		//message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		//c.hub.broadcast <- message
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	onlineUsersTicker := time.NewTicker(1 * time.Second)
	defer func() {
		ticker.Stop()
		onlineUsersTicker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-onlineUsersTicker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))

			messagableUsers, err := getMessagableUsers(c.UserID)
			if err != nil {
				log.Println(err)
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			message := WriteMessage{
				Type: "messagableUsers",
				Data: map[string][]BasicUserInfo{
					"messagableUsers": messagableUsers,
				},
			}

			jsonMessage, _ := json.Marshal(message)
			w.Write(jsonMessage)
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	var userData readwritemodels.Payload

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Read WebSocket messages and handle headers
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			return
		}

		if messageType == websocket.TextMessage {
			var message struct {
				Type   string `json:"type"`
				Header string `json:"header"`
				Value  string `json:"value"`
			}
			if err := json.Unmarshal(p, &message); err != nil {
				log.Println("WebSocket message unmarshal error:", err)
				continue
			}

			if message.Type == "header" {
				// Validate headers, e.g., user authentication
				// Do something with message.Header and message.Value
				if message.Header == "Authorization" {
					payload, err := middleware.ValidateTokenWebSocket(message.Value)
					if err != nil {
						log.Println("Error validating websocket token:", err)
						continue
					}
					userData = *payload
					break
				}
			}
		}
	}

	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256), UserID: userData.UserId, Username: userData.DisplayName}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}

func getMessagableUsers(userId string) ([]BasicUserInfo, error) {
	var messagableUsers []BasicUserInfo

	followees, err := followercontrollers.SelectFolloweesOfSpecificUser(dbutils.DB, userId)
	if err != nil {
		return nil, fmt.Errorf("error retrieving followee's of users: %w", err)
	}

	chats, err := chatcontrollers.SelectAllChatsByUser(dbutils.DB, userId)
	if err != nil {
		return nil, fmt.Errorf("error retrieving chats by user: %w", err)
	}

	for _, f := range followees.Followers {
		fUser, err := usercontrollers.GetUser(dbutils.DB, "", dbstatements.SelectUserByID, f.FolloweeId)
		if err != nil {
			return nil, fmt.Errorf("error getting user from followee id: %w", err)
		}
		messagableUsers = append(messagableUsers, BasicUserInfo{
			UUID:            f.FolloweeId,
			Name:            fUser.UserInfo.DisplayName,
			LoggedInStatus:  fUser.UserInfo.IsLoggedIn,
			LastMessageTime: time.Now(),
		})
	}

	for _, c := range chats.Chats {
		cUser := &dbmodels.UserProfileData{}

		if userId == c.SenderId {
			cUser, err = usercontrollers.GetUser(dbutils.DB, "", dbstatements.SelectUserByID, c.ReceiverId)
		} else if userId == c.ReceiverId {
			cUser, err = usercontrollers.GetUser(dbutils.DB, "", dbstatements.SelectUserByID, c.SenderId)
		}

		if err != nil {
			return nil, fmt.Errorf("error getting user: %w", err)
		}

		if !containsUser(messagableUsers, cUser.UserInfo.UserId) {
			messagableUsers = append(messagableUsers, BasicUserInfo{
				UUID:            cUser.UserInfo.UserId,
				Name:            cUser.UserInfo.DisplayName,
				LoggedInStatus:  cUser.UserInfo.IsLoggedIn,
				LastMessageTime: time.Now(),
			})
		}

	}

	return messagableUsers, nil
}

func createMarshalledWriteMessage(typ string, data interface{}) []byte {
	var writeMessage WriteMessage
	writeMessage.Type = typ
	writeMessage.Data = data
	marshalledData, err := json.Marshal(writeMessage)
	if err != nil {
		log.Printf("error: %v", err)
	}
	return marshalledData
}

func containsUser(s []BasicUserInfo, userId string) bool {
	for _, v := range s {
		if v.UUID == userId {
			return true
		}
	}

	return false
}
