import React, { CSSProperties, useContext, useEffect, useState } from 'react'
import { AiOutlineClose } from 'react-icons/ai';
import { UserContext } from '../../context/AuthContext';
import { useWebSocketContext } from '../../context/WebSocketContext';
import { WebSocketReadMessage, WebSocketWriteMessage } from '../../Socket';
import styles from './Chat.module.css'
import Conversation from './Conversation';
import UserChatDisplay from './UserChatDisplay';

export interface ChatProps {
    chat_id: string
    sender_id: string
    receiver_id: string
    group_id: string
    creation_date: string
}

export interface ChatInfo {
    uuid: string
    name: string
    logged_in_status: number
    last_message: string
    last_message_time: string
    is_group: boolean
}

const FriendsMessagingList: React.FC = () => {
    const userContext = useContext(UserContext);
    const { message, sendMessage } = useWebSocketContext();
    const [messageToSend, setMessageToSend] = useState<WebSocketReadMessage>({
        type: "",
        info: ""
    });
    const [currentUserChat, setCurrentUserChat] = useState<string | null>(null)
    const [isChatGroup, setIsChatGroup] = useState(false)
    const [currentUserChatDisplayName, setCurrentUserChatDisplayName] = useState<string | null>(null)
    const [messagableUsers, setMessagableUsers] = useState<ChatInfo[]>([])

    const handleSendMessage = () => {
        if (messageToSend) {
            console.log(messageToSend);
            setTimeout(() => {
                sendMessage(messageToSend)
            }, 200)
        }
    };

    useEffect(() => {
        setMessageToSend({
            type: "messagable_users",
            info: {
                receiver: userContext.user?.userId,
            },
        })
    }, [currentUserChat])

    useEffect(() => {
        if (messageToSend) {
            handleSendMessage();
        }
    }, [messageToSend]);

    useEffect(() => {
        if (message?.type == "messagable_users" && message.data) {
            console.log(message.data);

            setMessagableUsers(message.data.messagableUsers)
        } else if (message?.type == "online_user" && message.data) {
            const newState = messagableUsers.map((user) => {
                if (user.name === message.data.username) {
                    return { ...user, logged_in_status: message.data.online ? 1 : 0 }
                }
                return user
            })

            setMessagableUsers(newState)
        } else if (message?.type == "private_message" && message.data) {
            setMessageToSend({
                type: "messagable_users",
                info: {
                    receiver: userContext.user?.userId,
                },
            })
        }
    }, [message])

    const closeStyle: CSSProperties = {
        position: 'sticky',
        top: '10px',
        left: '10px',
        margin: "10px",
        verticalAlign: "middle",
        color: "red",
        borderRadius: "50%",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        height: "3%",
        width: "3%",
        zIndex: '2'
    };

    return (
        <>
            {currentUserChat && currentUserChatDisplayName ?
                <div
                    className={styles.messagingcontainer}>
                    <div className={styles.conversationHeader}>
                        <span style={closeStyle} onClick={() => {
                            setCurrentUserChat(null)
                        }}>
                            <AiOutlineClose />
                        </span>
                        <h3>{currentUserChatDisplayName}</h3>
                    </div>
                    <Conversation
                        message={message}
                        sendMessage={sendMessage}
                        receiverName={currentUserChatDisplayName}
                        receiverID={currentUserChat}
                        isGroup={isChatGroup} />
                </div>
                : messagableUsers ? (
                    <div
                        className={styles.messagableUsersContainer}>
                        {messagableUsers.map((user: ChatInfo) => (
                            <div
                                key={user.uuid}
                                onClick={() => {
                                    console.log("clicked");
                                    setCurrentUserChatDisplayName(user.name)
                                    setCurrentUserChat(user.uuid)
                                    setIsChatGroup(user.is_group)
                                    setMessageToSend({
                                        type: "open_chat",
                                        info: user.is_group ?
                                            {
                                                group_id: user.uuid
                                            }
                                            :
                                            {
                                                receiver: user.uuid,
                                            },
                                    });
                                }}
                                className={styles.userChatContainer}>
                                <UserChatDisplay
                                    follower_id={userContext.user ? userContext.user.userId : ""}
                                    followee_id={user.uuid}
                                    recipient_name={user.name}
                                    last_message={user.last_message}
                                    is_logged_in={user.logged_in_status}
                                    is_group={user.is_group} />
                            </div>
                        ))}
                    </div>
                ) : null}

        </>
    )
}

export default FriendsMessagingList 