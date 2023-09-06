import React, { useEffect, useState } from 'react'
import { WebSocketReadMessage, WebSocketWriteMessage } from '../../Socket'
import styles from './Chat.module.css'

interface LastMessageProps {
  follower_id: string
  followee_id: string
  last_message: string
}

const LastMessage: React.FC<LastMessageProps> = ({
  follower_id,
  followee_id,
  last_message
}) => {
  return (
    <div
      className={styles.lastmessagecontainer}>
      {last_message}
    </div>
  )
}

export default LastMessage