import React, { useState } from "react";
import { IoMdNotifications } from "react-icons/io";
import styles from "./Notification.module.css";
import SideModal from "../Containers/SideModal";
import NotificationsTable from "./NotificationsTable";

const Notification = () => {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [notificationCount, setNotificationCount] = useState<number | null>(null);

  return (
    <>
      <div className={styles.notificationbutton} onClick={() => {}}>
        <IoMdNotifications />
        {notificationCount && notificationCount > 0 && (
          <div className={styles.badge}>{notificationCount}</div>
        )}
      </div>
      <SideModal open={isModalOpen} onClose={() => setIsModalOpen(false)}>
        <NotificationsTable />
      </SideModal>
    </>
  );
};

export default Notification;
