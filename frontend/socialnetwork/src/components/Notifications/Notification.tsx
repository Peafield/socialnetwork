import React, { useEffect, useState } from 'react'
import { NotificationProps } from './NotificationsTable'
import styles from './Notification.module.css'
import { getUserByUserID } from '../../controllers/GetUser'
import { group } from 'console'
import { ProfileProps } from '../Profile/Profile'

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
            {notificationMessage}
        </div>
    )
}

function composeNotificationMessage(props: NotificationProps, senderName: string) {
    let message = ""

    message += senderName + " "

    if (props.reaction_type != "") {
        message += props.reaction_type + "d your "
        message += notificationTypeForLikingOrDisliking(props.post_id, props.event_id, props.comment_id)
    } else {

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