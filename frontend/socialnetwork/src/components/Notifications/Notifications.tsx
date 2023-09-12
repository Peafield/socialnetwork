import React, { useContext, useEffect, useState } from "react";
import { IoMdNotifications } from "react-icons/io";
import styles from "./Notification.module.css";
import { useWebSocketContext } from "../../context/WebSocketContext";
import { WebSocketReadMessage } from "../../Socket";
import { UserContext } from "../../context/AuthContext";
import { NotificationProps } from "./NotificationsTable";

interface NotificationsProps {
  setIsModalOpen: React.Dispatch<React.SetStateAction<boolean>>;
  setSideModalDisplay: React.Dispatch<React.SetStateAction<string | null>>;
}

const Notifications: React.FC<NotificationsProps> = ({
  setIsModalOpen,
  setSideModalDisplay
}) => {
  const userContext = useContext(UserContext);
  const { message, sendMessage } = useWebSocketContext();
  const [notificationCount, setNotificationCount] = useState<number | null>(null);
  let messageToSend: WebSocketReadMessage

  const handleOpenNotifications = () => {
    messageToSend = {
      type: "open_notifications",
      info: {
        userId: userContext.user?.userId
      }
    }

    sendMessage(messageToSend)
  }

  useEffect(() => {
    if (message?.type == "notification" || message?.type == "open_notifications") {
      const unreadMessagesCount = message.data?.filter((notification: NotificationProps) => notification.read_status == 0).length
      setNotificationCount(unreadMessagesCount > 0 ? unreadMessagesCount : null)
    }
  }, [message])

  return (
    <div className={styles.notificationbutton} onClick={() => { setIsModalOpen(true); setSideModalDisplay("notifications"); handleOpenNotifications(); }}>
      <IoMdNotifications />
      {notificationCount && notificationCount > 0 && (
        <div className={styles.badge}>{notificationCount}</div>
      )}
    </div>
  );
};

export default Notifications;
