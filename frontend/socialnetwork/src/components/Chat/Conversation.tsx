import React, { useContext, useEffect, useRef, useState } from 'react'
import { AiOutlineSend } from 'react-icons/ai'
import { UserContext } from '../../context/AuthContext'
import { getUserByUserID } from '../../controllers/GetUser'
import { WebSocketReadMessage, WebSocketWriteMessage } from '../../Socket'
import { ProfileProps } from '../Profile/Profile'
import styles from './Chat.module.css'

export interface ChatMessageProps {
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
  isGroup: boolean
}

const Conversation: React.FC<ConversationProps> = ({
  message,
  sendMessage,
  receiverName,
  receiverID,
  isGroup
}) => {
  const userContext = useContext(UserContext)
  const [chats, setChats] = useState<ChatMessageProps[] | null>(null)
  const [chatID, setChatID] = useState<string | null>(null)
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
      sendMessage(messageToSend)
    }
  }

  const reloadChat = () => {
    sendMessage({
      type: "open_chat",
      info: isGroup ?
        {
          group_id: receiverID
        }
        :
        {
          receiver: receiverID
        }
    })
  }

  const mapMessageData = async (messageData: any) => {
    let newChats: ChatMessageProps[] = [];

    for (const chat of messageData) {
      let sender = "";

      if (userContext.user && chat.sender_id == userContext.user?.userId) {
        sender = userContext.user?.displayName;
      } else {
        if (!isGroup) {
          sender = receiverName;
        } else {
          const user: ProfileProps = await getUserByUserID(chat.sender_id);
          sender = user.display_name;
        }
      }

      newChats.push({
        messageId: chat.message_id,
        senderName: sender,
        timeSent: chat.creation_date,
        message: chat.message,
      });
    }

    console.log(newChats);
    setChats(newChats);
  };



  useEffect(() => {
    if (message?.type == "open_chat") {
      if (typeof message.data === 'string') {
        setChatID(message.data)
      } else {
        setChatID(message.data[0].chat_id)
        mapMessageData(message.data)
      }
    } else if (message?.type == "private_message") {
      if (message.data.every((chat: any) => {
        return chat.chat_id == chatID
      })) {
        reloadChat()
      }
    }
  }, [message])

  return (
    <>
      <div
        className={styles.conversationcontainer}>
        <div className={styles.chatmessagescontainer}>
          {chats &&
            chats.map((chat: ChatMessageProps, index: number) => (
              <div key={chat.messageId}
                className={
                  chat.senderName === userContext.user?.displayName
                    ? styles.userchatinfo
                    : styles.receiverchatinfo}>
                <div className={
                  chat.senderName === userContext.user?.displayName
                    ? styles.userchatinfo
                    : styles.receiverchatinfo}
                  style={{ color: 'gray' }}>
                  {`${chat.senderName} (${chat.timeSent.split("T")[1].split("Z")[0]} ${chat.timeSent.split("T")[0]})`}
                </div>
                <div
                  ref={index === chats.length - 1 ? lastMessageRef : null}
                  className={
                    chat.senderName === userContext.user?.displayName
                      ? styles.userchatmessagecontainer
                      : styles.receiverchatmessagecontainer
                  }
                >
                  <p>{chat.message}</p>
                </div>
              </div>
            ))}
        </div>
        <div
          className={styles.sendmessagecontainer}>
          <textarea placeholder='Type message...' value={messageToSend.info.message} onChange={(event) => setMessageToSend({
            type: "private_message",
            info: isGroup ? {
              message: event.target.value,
              group_id: receiverID
            }
              :
              {
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