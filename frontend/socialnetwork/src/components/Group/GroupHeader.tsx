import React, { useContext, useState } from 'react'
import { UserContext } from '../../context/AuthContext'
import { useWebSocketContext } from '../../context/WebSocketContext'
import { handleAPIRequest } from '../../controllers/Api'
import { getCookie } from '../../controllers/SetUserContextAndCookie'
import { WebSocketReadMessage } from '../../Socket'
import Snackbar from '../feedback/Snackbar'
import styles from './Group.module.css'

interface GroupHeaderProps {
    group_id: string
    title: string,
    description: string,
    creator_id: string
    creator_name: string
    isUserMember: boolean
    groupTab: string
    getGroupTab: (tab: string) => void;
}

const GroupHeader: React.FC<GroupHeaderProps> = ({
    group_id,
    title,
    description,
    creator_id,
    creator_name,
    isUserMember,
    groupTab,
    getGroupTab
}) => {
    const userContext = useContext(UserContext)
    const { message, sendMessage } = useWebSocketContext();
    let messageToSend: WebSocketReadMessage = {
        type: "",
        info: ""
    }
    const [error, setError] = useState<string | null>(null);
    const [snackbarOpen, setSnackbarOpen] = useState<boolean>(false);
    const [snackbarType, setSnackbarType] = useState<
        "success" | "error" | "warning"
    >("error");

    const handleTabChange = (e: React.MouseEvent<HTMLButtonElement>) => {
        getGroupTab(e.currentTarget.value)
    }

    const handleRequest = async (e: { preventDefault: () => void }) => {
        e.preventDefault();
        const data = {
            data: {
                requester_id: userContext.user?.userId,
                group_id: group_id
            }
        };
        const options = {
            method: "POST",
            headers: {
                Authorization: "Bearer " + getCookie("sessionToken"),
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data),
        };
        try {
            const response = await handleAPIRequest(`/groupmembers?has_requested=${encodeURIComponent("true")}`, options);
            if (response && response.status === "success") {
                setSnackbarType("success");
                setSnackbarOpen(true);
                messageToSend = {
                    type: "notification",
                    info: {
                        receiver: creator_id,
                        group_id: group_id,
                        action_type: "request"
                    }
                }
                sendMessage(messageToSend)

            }
        } catch (error) {
            if (error instanceof Error) {
                setError("Could not create post");
                setSnackbarType("error");
                setSnackbarOpen(true);
            } else {
                setError("An unexpected error occurred");
                setSnackbarType("error");
                setSnackbarOpen(true);
            }
        }
    };

    return (
        <div
            className={styles.groupheadercontainer}>
            <div className={styles.toprow}>
                <div>
                    Name of Group: {title}
                </div>
                <div>
                    Created By: {creator_name}
                </div>
            </div>
            <div className={styles.middlerow}>
                <div>
                    Description: {description}
                </div>
            </div>
            <div className={styles.bottomrow}>
                {!isUserMember ?
                    <div style={{ display: 'flex', alignItems: 'center' }}>
                        <button onClick={handleRequest} style={{ color: '#fa4d6a' }}>
                            Request Membership
                        </button>
                    </div>
                    :
                    null}
                <div style={{ display: 'flex', alignItems: 'center' }}><button onClick={handleTabChange} value="posts" style={groupTab == "posts" ? { textDecorationLine: 'underline' } : undefined}>Posts</button></div>
                <div style={{ display: 'flex', alignItems: 'center' }}><button onClick={handleTabChange} value="members" style={groupTab == "members" ? { textDecorationLine: 'underline' } : undefined}>Members</button></div>
                <div style={{ display: 'flex', alignItems: 'center' }}><button onClick={handleTabChange} value="events" style={groupTab == "events" ? { textDecorationLine: 'underline' } : undefined}>Events</button></div>
            </div>
            <Snackbar
                open={snackbarOpen}
                onClose={() => {
                    setSnackbarOpen(false);
                    setError(null);
                }}
                message={error ? error : "Request Sent!"}
                type={snackbarType}
            />
        </div>
    )
}

export default GroupHeader