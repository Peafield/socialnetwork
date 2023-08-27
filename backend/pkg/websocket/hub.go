// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients to be sent to everyone.
	broadcast chan []byte

	// Inbound messages from the clients to be sent to users who follow respective clients.
	broadcastToFollowers chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:            make(chan []byte),
		broadcastToFollowers: make(chan []byte),
		register:             make(chan *Client),
		unregister:           make(chan *Client),
		clients:              make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

// getClientByUsername retrieves the client object based on the username
func (h *Hub) GetClientByUsername(username string) *Client {
	for client := range h.clients {
		if client.Username == username {
			return client
		}
	}
	return nil
}

// getClientByUsername retrieves the client object based on the userID
func (h *Hub) GetClientByID(userID string) *Client {
	for client := range h.clients {
		if client.UserID == userID {
			return client
		}
	}
	return nil
}
