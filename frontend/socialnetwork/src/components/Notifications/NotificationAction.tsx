import React, { useState } from 'react'
import { NotificationProps } from './NotificationsTable'
import { TiArrowRightThick, TiMediaRecord, TiTick, TiTimes, TiTimesOutline, TiUser } from 'react-icons/ti'
import { useNavigate } from 'react-router-dom'
import { updateFollowRequest } from '../../controllers/Follower/UpdateFollowRequest'
import styles from './Notification.module.css'
import { DeleteNotification } from '../../controllers/Notifications/DeleteNotification'
import { UpdateGroupRequestPending } from '../../controllers/Group/UpdateGroupMembers'

const NotificationAction: React.FC<NotificationProps> = (props) => {
    const navigate = useNavigate();
    const action = getNotificationAction(props)
    const [decisionMade, setDecisionMade] = useState(false)

    const handleAccept = () => {
        if (props.action_type === "request") {
            if (props.group_id !== "") {
                UpdateGroupRequestPending(props.sender_id, props.group_id, true)
            } else {
                updateFollowRequest(props.sender_id, 1)
            }
        } else if (props.action_type === "invite" && props.group_id !== "") {
            UpdateGroupRequestPending(props.receiver_id, props.group_id, true)
        }
        setDecisionMade(true)
        DeleteNotification(props.notification_id)
    }

    const handleDecline = () => {
        if (props.action_type === "request") {
            if (props.group_id !== "") {
                UpdateGroupRequestPending(props.sender_id, props.group_id, false)
            } else {
                updateFollowRequest(props.sender_id, 0)
            }
        } else if (props.action_type === "invite" && props.group_id !== "") {
            UpdateGroupRequestPending(props.receiver_id, props.group_id, false)
        }
        setDecisionMade(true)
        DeleteNotification(props.notification_id)
    }

    const handleClick = () => {
        setDecisionMade(true)
        DeleteNotification(props.notification_id)
    }
    return (
        <div className={styles.notificationaction}>
            {action === "goto" ?
                !decisionMade ?
                    <span style={{ color: '#fa4d6a' }} onClick={handleClick}>
                        <TiTimesOutline />
                    </span>
                    : null
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