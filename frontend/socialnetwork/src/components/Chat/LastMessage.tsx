import React from 'react'
import styles from './Chat.module.css'

interface LastMessageProps {
    follower_id: string
    followee_id: string
}

const LastMessage: React.FC<LastMessageProps> = () => {
  return (
    <div
    className={styles.lastmessagecontainer}>Need to return last message</div>
  )
}

export default LastMessage