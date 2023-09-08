import React, { CSSProperties, useContext, useEffect, useState } from 'react'
import { AiOutlineClose } from 'react-icons/ai';
import { UserContext } from '../../context/AuthContext';
import { WebSocketReadMessage, WebSocketWriteMessage } from '../../Socket';
import styles from './Chat.module.css'
import Conversation from './Conversation';
import UserChatDisplay from './UserChatDisplay';

interface FriendsMessagingListProps {
    message: WebSocketWriteMessage | null,
    sendMessage: (messegeToSend: WebSocketReadMessage) => void
}

const FriendsMessagingList: React.FC<FriendsMessagingListProps> = ({
    message,
    sendMessage
}) => {
    const userContext = useContext(UserContext)
    const [messageToSend, setMessageToSend] = useState<WebSocketReadMessage>({
        type: "",
        info: ""
    });
    const [currentUserChat, setCurrentUserChat] = useState<string | null>(null)
    const [currentUserChatDisplayName, setCurrentUserChatDisplayName] = useState<string | null>(null)
    const [messagableUsers, setMessagableUsers] = useState<any[]>([])

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
            setMessagableUsers(message.data.messagableUsers)
        } else if (message?.type == "online_user" && message.data) {
            console.log(message);

            const newState = messagableUsers.map((user) => {
                if (user.Name == message.data.username) {
                    return { ...user, LoggedInStatus: message.data.online ? 1 : 0 }
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
                        receiverID={currentUserChat} />
                </div>
                : messagableUsers ? (
                    <div
                        className={styles.messagableUsersContainer}>
                        {messagableUsers.map((user: any) => (
                            <div
                                key={user.UUID}
                                onClick={() => {
                                    console.log("clicked");
                                    setCurrentUserChatDisplayName(user.Name)
                                    setCurrentUserChat(user.UUID)
                                    setMessageToSend({
                                        type: "open_chat",
                                        info: {
                                            receiver: user.UUID,
                                        },
                                    });
                                }}
                                className={styles.userChatContainer}>
                                <UserChatDisplay
                                    follower_id={userContext.user ? userContext.user.userId : ""}
                                    followee_id={user.UUID}
                                    last_message={user.LastMessage}
                                    is_logged_in={user.LoggedInStatus} />
                            </div>
                        ))}
                    </div>
                ) : null}

        </>
    )
}

export default FriendsMessagingList 