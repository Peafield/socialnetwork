import React, { CSSProperties, useEffect, useState } from 'react'
import { Link } from 'react-router-dom'
import Container from '../Containers/Container'
import { FaUserCircle } from 'react-icons/fa'
import styles from './Post.module.css'

interface PostHeaderProps {
  creatorId: string
  creatorDisplayName: string,
  creatorAvatar: string,
  creationDate: number
}

const PostHeader: React.FC<PostHeaderProps> = ({
  creatorId,
  creatorDisplayName,
  creatorAvatar,
  creationDate
}) => {
  const [profilePicUrl, setProfilePicUrl] = useState<string | null>(null)

  useEffect(() => {
    if (creatorAvatar) {
      const decodedAvatar = atob(creatorAvatar); // Decode base64-encoded avatar data
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

  }, [creatorAvatar])

  return (
    <>
      <div className={styles.postheadercontainer}>
        {(profilePicUrl && <img src={profilePicUrl} alt='Profile pic' className={styles.avatar}/>) || (
          <span className={styles.profileIconStyle}>
            <FaUserCircle />
          </span>
        )}
        <Link to={"/dashboard/user/" + creatorDisplayName} >{creatorDisplayName}</Link>
        <p>{creationDate}</p>
      </div>
    </>
  )
}

export default PostHeader