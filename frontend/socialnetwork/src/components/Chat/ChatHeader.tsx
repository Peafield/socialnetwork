import React, { useEffect, useState } from 'react'
import { FaUserCircle } from 'react-icons/fa';
import { useNavigate } from 'react-router-dom';
import { getUserByDisplayName, getUserByUserID } from '../../controllers/GetUser';
import { createImageURL } from '../../controllers/ImageURL';
import { ProfileProps } from '../Profile/Profile';
import styles from './Chat.module.css'

interface ChatHeaderProps {
    user_id: string
}

const ChatHeader: React.FC<ChatHeaderProps> = ({
    user_id
}) => {
    const navigate = useNavigate()
    const [user, setUser] = useState<ProfileProps | null>(null)
    const [userAvatarURL, setUserAvatarURL] = useState<string | null>(null)
    const [userLoading, setUserLoading] = useState(false)
    const [error, setError] = useState<string | null>(null)
    useEffect(() => {
        const fetchData = async () => {
            setUserLoading(true);

            try {
                if (user_id) {
                    const newprofile = await getUserByUserID(user_id)
                    setUser(newprofile)
                    if (newprofile.avatar) {
                        const url = createImageURL(newprofile.avatar)
                        setUserAvatarURL(url);
                        console.log(newprofile.avatar);

                        // Clean up the Blob URL when the component unmounts
                        return () => {
                            URL.revokeObjectURL(url);
                        };
                    }
                } else {
                    setError("could not find profile username")
                }
            } catch (error) {
                if (error instanceof Error) {
                    setError(error.message);
                    if (error.cause == 401) {
                        navigate("/signin")
                    }
                } else {
                    setError("An unexpected error occurred.");
                }
            }
            setUserLoading(false);
        };

        fetchData(); // Call the async function
    }, [user_id]);

    return (
        <div
            className={styles.chatheadercontainer}>
            <div className={styles.displaypicturecontainer}>
                {(userAvatarURL && (
                    <img
                        src={userAvatarURL}
                        alt="Profile pic"
                        className={styles.avatar}
                    />
                )) || (
                        <span className={styles.profileIconStyle}>
                            <FaUserCircle />
                        </span>
                    )}
            </div>
            <div className={styles.nameinfocontainer}>
                <div style={{ color: "black", fontWeight: "bold" }}>{user?.display_name}</div>
            </div>
        </div >
    )
}

export default ChatHeader