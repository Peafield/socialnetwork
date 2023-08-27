import React from 'react'
import ChatHeader from './ChatHeader'
import LastMessage from './LastMessage'
import styles from './Chat.module.css'

interface UserChatProps {
    follower_id: string
    followee_id: string
}

const UserChatDisplay: React.FC<UserChatProps> = ({
    follower_id,
    followee_id
}) => {
    return (
        <>
            <ChatHeader user_id={followee_id} />
            <LastMessage follower_id={follower_id} followee_id={followee_id} />
        </>
    )
}

export default UserChatDisplay