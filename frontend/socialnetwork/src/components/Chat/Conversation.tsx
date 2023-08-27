import React, { useContext, useEffect, useRef, useState } from 'react'
import { AiOutlineSend } from 'react-icons/ai'
import { UserContext } from '../../context/AuthContext'
import { WebSocketReadMessage, WebSocketWriteMessage } from '../../Socket'
import styles from './Chat.module.css'

export interface ChatProps {
  messageId: string
  senderName: string
  timeSent: string
  message: string
}

interface ConversationProps {
  message: WebSocketWriteMessage | null,
  sendMessage: (messegeToSend: WebSocketReadMessage) => void
  receiverName: string
  receiverID: string
}

const Conversation: React.FC<ConversationProps> = ({
  message,
  sendMessage,
  receiverName,
  receiverID
}) => {
  const userContext = useContext(UserContext)
  const [chats, setChats] = useState<ChatProps[] | null>(null)
  const [messageToSend, setMessageToSend] = useState<WebSocketReadMessage>({
    type: "",
    info: "",
  })
  const lastMessageRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (lastMessageRef.current) {
      // Scroll to the last message element
      lastMessageRef.current.scrollIntoView({ behavior: "smooth" });
    }
  }, [chats]);

  const handleSendMessage = () => {
    if (messageToSend) {
      console.log(messageToSend);
      sendMessage(messageToSend)
    }
  }

  useEffect(() => {
    if (message?.type == "open_chat" || message?.type == "private_message") {
      let newChats: ChatProps[] = [];
      message.data.map((chat: any) => {
        let sender = ""
        if (userContext.user && chat.sender_id == userContext.user?.userId) {
          sender = userContext.user?.displayName
        } else {
          sender = receiverName
        }
        newChats.push({
          messageId: chat.message_id,
          senderName: sender,
          timeSent: "",
          message: chat.message
        })
      })
      console.log(message.data);
      setChats(newChats)
    }
  }, [message])

  return (
    <>
      <div
        className={styles.conversationcontainer}>
        <div className={styles.chatmessagescontainer}>
          {chats &&
            chats.map((chat: ChatProps, index: number) => (
              <div
                key={chat.messageId}
                ref={index === chats.length - 1 ? lastMessageRef : null}
                className={
                  chat.senderName === receiverName
                    ? styles.receiverchatmessagecontainer
                    : styles.userchatmessagecontainer
                }
              >
                <p>{chat.message}</p>
              </div>
            ))}
        </div>
        <div
          className={styles.sendmessagecontainer}>
          <textarea placeholder='Type message...' value={messageToSend.info.message} onChange={(event) => setMessageToSend({
            type: "private_message",
            info: {
              message: event.target.value,
              receiver: receiverID
            }
          })}
            onKeyDown={(event) => {
              if (event.key === 'Enter') {
                event.preventDefault(); // Prevent new line from being inserted
                handleSendMessage();
                setMessageToSend(prevMessage => ({
                  ...prevMessage,
                  info: {
                    ...prevMessage.info,
                    message: '' // Reset the message
                  }
                }));
              }
            }} />
          <button onClick={() => {
            handleSendMessage()
            setMessageToSend(prevMessage => ({
              ...prevMessage,
              info: {
                ...prevMessage.info,
                message: '' // Reset the message
              }
            }));
          }}>
            <AiOutlineSend />
          </button>
        </div>
      </div>
    </>
  )
}

export default Conversation