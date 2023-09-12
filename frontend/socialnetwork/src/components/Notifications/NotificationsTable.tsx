import React, { useEffect, useState } from 'react'
import { useWebSocketContext } from '../../context/WebSocketContext';
import Notification from './Notification';
import styles from './Notification.module.css'

export interface NotificationProps {
  notification_id: string
  sender_id: string
  receiver_id: string
  group_id: string
  post_id: string
  event_id: string
  comment_id: string
  chat_id: string
  action_type: string
  read_status: number
  creation_date: string
}

const NotificationsTable = () => {
  const { message, sendMessage } = useWebSocketContext();
  const [notifications, setNotifications] = useState<NotificationProps[] | null>(null)

  useEffect(() => {
    if (message?.type == "notification" || message?.type == "open_notifications") {
      console.log(message.data);
      setNotifications(message.data)
    }
  }, [message])

  return (
    <>
      <div
        className={styles.allNotificationsContainer}>
        {notifications?.map((notification) => (
          <Notification
            key={notification.notification_id}
            notification_id={notification.notification_id}
            sender_id={notification.sender_id}
            receiver_id={notification.receiver_id}
            group_id={notification.group_id}
            post_id={notification.post_id}
            event_id={notification.event_id}
            comment_id={notification.comment_id}
            chat_id={notification.chat_id}
            action_type={notification.action_type}
            read_status={notification.read_status}
            creation_date={notification.creation_date} />
        ))}
      </div>
    </>
  )
}

export default NotificationsTable