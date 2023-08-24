import React, { useState } from 'react'
import { WebSocketMessage } from '../../Socket';

interface FriendsMessagingListProps {
    message: WebSocketMessage | null,
    sendMessage: (messegeToSend: string) => void
}

const FriendsMessagingList: React.FC<FriendsMessagingListProps> = ({
    message,
    sendMessage
}) => {
    const [messageToSend, setMessageToSend] = useState<string>("");

    const handleSendMessage = () => {
        if (messageToSend) {
            sendMessage(messageToSend);
        }
    };

    // Parse the received data if it's not undefined
    let messagableUsers = [];
    if (message && message.data) {
        messagableUsers = message.data.messagableUsers;
    }

    return (
        <>
            {messagableUsers ? (
                <div>
                    {messagableUsers.map((user: any) => (
                        <div key={user.UUID}>{user.Name}</div>
                    ))}
                </div>
            ) : null}
        </>
    )
}

export default FriendsMessagingList 