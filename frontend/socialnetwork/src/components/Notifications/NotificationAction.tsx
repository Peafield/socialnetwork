import React, { useState } from 'react'
import { NotificationProps } from './NotificationsTable'
import { TiArrowRightThick, TiTick, TiTimes, TiUser } from 'react-icons/ti'
import { useNavigate } from 'react-router-dom'
import { updateFollowRequest } from '../../controllers/Follower/UpdateFollowRequest'
import styles from './Notification.module.css'
import { DeleteNotification } from '../../controllers/Notifications/DeleteNotification'

const NotificationAction: React.FC<NotificationProps> = (props) => {
    const navigate = useNavigate();
    const action = getNotificationAction(props)
    const [decisionMade, setDecisionMade] = useState(false)

    const handleAccept = () => {
        if (props.action_type === "request") {
            updateFollowRequest(props.sender_id, 1)
        }
        setDecisionMade(true)
        DeleteNotification(props.notification_id)
    }

    const handleDecline = () => {
        if (props.action_type === "request") {
            updateFollowRequest(props.sender_id, 0)
        }
        setDecisionMade(true)
        DeleteNotification(props.notification_id)
    }
    return (
        <div className={styles.notificationaction}>
            {action === "goto" ?
                <span onClick={() => { navigate("/dashboard") }} style={{ color: '#fa4d6a' }}>
                    <TiArrowRightThick />
                </span>
                :
                !decisionMade ?
                    <div className={styles.yesorno}>
                        <span style={{ cursor: 'pointer', color: 'lightgreen' }} onClick={handleAccept}>
                            <TiTick />
                        </span>
                        <span style={{ cursor: 'pointer', color: 'darkred' }} onClick={handleDecline}>
                            <TiTimes />
                        </span>
                    </div>
                    :
                    <div className={styles.yesorno}>
                        <span>
                            <TiUser />
                        </span>
                    </div>

            }
        </div>
    )
}

function getNotificationAction(props: NotificationProps) {
    switch (props.action_type) {
        case "like":
            return "goto"
        case "dislike":
            return "goto"
        case "follow":
            return "goto"
        case "request":
            return "yesorno"
        case "comment":
            return "goto"
        case "invite":
            return "yesorno"
    }
}

export default NotificationAction