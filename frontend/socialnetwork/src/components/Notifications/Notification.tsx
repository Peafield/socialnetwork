import React, { useEffect, useState } from 'react'
import { NotificationProps } from './NotificationsTable'
import styles from './Notification.module.css'
import { getUserByUserID } from '../../controllers/GetUser'
import { group } from 'console'
import { ProfileProps } from '../Profile/Profile'
import NotificationAction from './NotificationAction'

const Notification: React.FC<NotificationProps> = (props) => {
    const [notificationMessage, setNotificationMessage] = useState<string>("")
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchUserName = async () => {
            try {
                const user: ProfileProps = await getUserByUserID(props.sender_id)

                if (user) {
                    setNotificationMessage(composeNotificationMessage(props, user.display_name))
                }
            } catch (error) {
                if (error instanceof Error) {
                    setError(error.message);
                } else {
                    setError("An unexpected error occurred.");
                }
            }
        }

        fetchUserName();
    }, [])

    return (
        <div
            className={styles.notificationcontainer}>
            <div style={{ width: '60%' }}>
                {notificationMessage}
            </div>
            <NotificationAction
                notification_id={props.notification_id}
                sender_id={props.sender_id}
                receiver_id={props.receiver_id}
                group_id={props.group_id}
                post_id={props.post_id}
                event_id={props.event_id}
                comment_id={props.comment_id}
                chat_id={props.chat_id}
                action_type={props.action_type}
                read_status={props.read_status}
                creation_date={props.creation_date} />
        </div>
    )
}

function composeNotificationMessage(props: NotificationProps, senderName: string) {
    let message = ""

    message += senderName + " "

    switch (props.action_type) {
        case "like":
            message += props.action_type + "d your "
            message += notificationTypeForLikingOrDisliking(props.post_id, props.event_id, props.comment_id)
            break
        case "dislike":
            message += props.action_type + "d your "
            message += notificationTypeForLikingOrDisliking(props.post_id, props.event_id, props.comment_id)
            break
        case "follow":
            message += props.action_type + "ed you"
            break
        case "request":
            if (props.group_id !== "") {
                message += props.action_type + "ed to join your group - %Group Name%"
            } else {
                message += props.action_type + "ed to follow you"
            }
            break
        case "comment":
            message += props.action_type + "ed on your post"
            break
        case "invite":
            message += props.action_type + "d you to their group - %Group Name%"
            break
    }

    return message
}

function notificationTypeForLikingOrDisliking(
    postId: string,
    eventId: string,
    commentId: string,
) {
    if (postId != "") {
        return "post"
    }
    if (eventId != "") {
        return "event"
    }
    if (commentId != "") {
        return "comment"
    }
}

export default Notification