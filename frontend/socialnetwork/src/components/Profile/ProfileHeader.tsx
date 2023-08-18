import React, { useEffect, useState } from 'react'
import { FaUserCircle } from 'react-icons/fa'
import { defer } from 'react-router-dom'
import Container from '../Containers/Container'
import styles from './Profile.module.css'

interface ProfileHeaderProps {
    display_name: string,
    avatar: string,
    num_of_posts: number,
    followers: number,
    following: number
}

const ProfileHeader: React.FC<ProfileHeaderProps> = ({
    display_name,
    avatar,
    num_of_posts,
    followers,
    following
}) => {
    const [profilePicUrl, setProfilePicUrl] = useState<string | null>(null)

    useEffect(() => {
        if (avatar) {
            const decodedAvatar = atob(avatar); // Decode base64-encoded avatar data
            const avatarBuffer = new ArrayBuffer(decodedAvatar.length);
            const avatarView = new Uint8Array(avatarBuffer);
            for (let i = 0; i < decodedAvatar.length; i++) {
                avatarView[i] = decodedAvatar.charCodeAt(i);
            }

            const blob = new Blob([avatarBuffer]);
            const url = URL.createObjectURL(blob);
            console.log(url);

            setProfilePicUrl(url)

            // Clean up the Blob URL when the component unmounts
            return () => {
                URL.revokeObjectURL(url);
            };
        }
    }, [avatar])

    return (
        <Container>
            <div className={styles.profileheadercontainer}>
                <div>
                    {(profilePicUrl && <img src={profilePicUrl} alt='Profile pic' className={styles.avatar} />) || (
                        <span className={styles.profileIconStyle}>
                            <FaUserCircle />
                        </span>
                    )}
                </div>
                <div>
                    {display_name}
                </div>
                <div>
                    Posts: {num_of_posts}
                </div>
                <div>
                    Followers: {followers}
                </div>
                <div>
                    Following: {following}
                </div>
            </div>
        </Container>
    )
}

export default ProfileHeader