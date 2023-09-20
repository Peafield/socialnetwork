import React from 'react'
import ChatHeader from './ChatHeader'
import LastMessage from './LastMessage'
import styles from './Chat.module.css'
import { WebSocketReadMessage, WebSocketWriteMessage } from '../../Socket'

interface UserChatProps {
    follower_id: string
    followee_id: string
    recipient_name: string
    last_message: string
    is_logged_in: number
    is_group: boolean
}

const UserChatDisplay: React.FC<UserChatProps> = ({
    follower_id,
    followee_id,
    recipient_name,
    last_message,
    is_logged_in,
    is_group
}) => {
    return (
        <>
            <ChatHeader user_id={followee_id} is_logged_in={is_logged_in} user_name={recipient_name} />
            <LastMessage follower_id={follower_id} followee_id={followee_id} last_message={last_message} />
        </>
    )
}

export default UserChatDisplay