import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { FaUserCircle, FaGlobeAfrica } from "react-icons/fa";
import {FaPeopleGroup } from "react-icons/fa6"
import { IoPeopleCircle } from "react-icons/io5"
import styles from "./Post.module.css";
import FormatDate from "../../helpers/DateConversion";

interface PostHeaderProps {
  creatorId: string;
  creatorDisplayName: string;
  creatorAvatar: string;
  creationDate: string;
  postPrivacyLevel: number;
}

const PostHeader: React.FC<PostHeaderProps> = ({
  creatorId,
  creatorDisplayName,
  creatorAvatar,
  creationDate,
  postPrivacyLevel,

}) => {
  const [profilePicUrl, setProfilePicUrl] = useState<string | null>(null);

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

      setProfilePicUrl(url);

      // Clean up the Blob URL when the component unmounts
      return () => {
        URL.revokeObjectURL(url);
      };
    }
  }, [creatorAvatar]);
  let formattedDate = FormatDate(creationDate);
  let PrivacyIconType = postPrivacyLevel === 0 ? FaGlobeAfrica : 
                      postPrivacyLevel === 1 ? FaPeopleGroup : IoPeopleCircle;

  return (
    <div className={styles.postheadercontainer}>
      <div className={styles.postHeaderInfoContainer}>
        <div>
          {(profilePicUrl && (
            <img
              src={profilePicUrl}
              alt="Profile pic"
              className={styles.avatar}
            />
          )) || (
            <span className={styles.profileIconStyle}>
              <IoPeopleCircle />
            </span>
          )}
        </div>
        <div>
          <Link to={"/dashboard/user/" + creatorDisplayName}>
            {creatorDisplayName}
          </Link>
          <p>{formattedDate} <PrivacyIconType /></p>
        </div>
      </div>
    </div>
  );
};

export default PostHeader;
