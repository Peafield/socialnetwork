import React, { useEffect, useState } from "react";
import { FaUserCircle } from "react-icons/fa";
import Container from "../Containers/Container";
import styles from "./Profile.module.css";

interface ProfileHeaderProps {
  first_name: string,
  last_name: string,
  display_name: string;
  avatar: string;
  num_of_posts: number;
  followers: number;
  following: number;
  about_me: string;
}

const ProfileHeader: React.FC<ProfileHeaderProps> = ({
  first_name,
  last_name,
  display_name,
  avatar,
  num_of_posts,
  followers,
  following,
  about_me,
}) => {
  const [profilePicUrl, setProfilePicUrl] = useState<string | null>(null);

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

      setProfilePicUrl(url);

      // Clean up the Blob URL when the component unmounts
      return () => {
        URL.revokeObjectURL(url);
      };
    }
  }, [avatar]);

  return (
    <Container>
      <div className={styles.profileheadercontainer}>
        <div className={styles.displaypicturecontainer}>
          {(profilePicUrl && (
            <img
              src={profilePicUrl}
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
          <div>{first_name} {last_name}</div>
          <div style={{ color: "gray" }}>{display_name}</div>
        </div>
        <div className={styles.otherprofileinfocontainer}>
          <div className={styles.profilestatscontainer}>
          <div>{num_of_posts} Posts</div>
          <div>{followers} Followers</div>
          <div>Following {following}</div>
          </div>
          <div className={styles.aboutmecontainer}>
            <div>{about_me}</div>
          </div>
        </div>
      </div>
    </Container>
  );
};

export default ProfileHeader