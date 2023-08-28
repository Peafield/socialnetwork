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
            sendMessage(messageToSend);
        }
    };

    useEffect(() => {
        if (currentUserChat) {
            handleSendMessage();
        }
    }, [currentUserChat]);

    useEffect(() => {
        if (message?.type == "messagable_users" && message.data) {
            setMessagableUsers(message.data.messagableUsers)
        }
    }, [message])

    const closeStyle: CSSProperties = {
        margin: "10px",
        verticalAlign: "middle",
        color: "red",
        borderRadius: "50%",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        height: "3%",
        width: "3%",
    };

    return (
        <>
            <div
                className={styles.messagingcontainer}>
                {currentUserChat && currentUserChatDisplayName ?
                    <div>
                        <span style={closeStyle} onClick={() => {
                            setCurrentUserChat(null)
                        }}>
                            <AiOutlineClose />
                        </span>
                        <div>
                            <Conversation
                                message={message}
                                sendMessage={sendMessage}
                                receiverName={currentUserChatDisplayName}
                                receiverID={currentUserChat} />
                        </div>
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
                                    <UserChatDisplay follower_id={userContext.user ? userContext.user.userId : ""} followee_id={user.UUID} last_message={user.LastMessage} />
                                </div>
                            ))}
                        </div>
                    ) : null}
            </div>
        </>
    )
}

export default FriendsMessagingList 