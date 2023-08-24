import React, { useContext, useEffect, useState } from "react";
import { FaUserCircle } from "react-icons/fa";
import { UserContext } from "../../context/AuthContext";
import { handleAPIRequest } from "../../controllers/Api";
import { getFollowerData } from "../../controllers/Follower/GetFollower";
import { newFollower } from "../../controllers/Follower/NewFollower";
import { unfollow } from "../../controllers/Follower/Unfollow";
import { createImageURL } from "../../controllers/ImageURL";
import { getCookie } from "../../controllers/SetUserContextAndCookie";
import Container from "../Containers/Container";
import styles from "./Profile.module.css";

interface ProfileHeaderProps {
  profile_id: string,
  first_name: string,
  last_name: string,
  display_name: string;
  avatar: string;
  num_of_posts: number;
  followers: number;
  following: number;
  about_me: string;
}

export interface FollowerProps {
  follower_id: string,
  followee_id: string,
  following_status: number,
  request_pending: number,
  creation_date: string
}

const ProfileHeader: React.FC<ProfileHeaderProps> = ({
  profile_id,
  first_name,
  last_name,
  display_name,
  avatar,
  num_of_posts,
  followers,
  following,
  about_me
}) => {
  const userContext = useContext(UserContext)
  const [profilePicUrl, setProfilePicUrl] = useState<string | null>(null);
  const [followerData, setFollowerData] = useState<FollowerProps>({
    follower_id: "",
    followee_id: "",
    following_status: 0,
    request_pending: 0,
    creation_date: ""
  });
  const [updateTrigger, setUpdateTrigger] = useState<number>(0)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    if (avatar) {
      const url = createImageURL(avatar)

      setProfilePicUrl(url);

      // Clean up the Blob URL when the component unmounts
      return () => {
        URL.revokeObjectURL(url);
      };
    }
  }, [avatar]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const following = await getFollowerData(profile_id)
        setFollowerData(following)

      } catch (error) {
        if (error instanceof Error) {
          setError(error.message);
        } else {
          setError("An unexpected error occurred.");
        }
      }
    }
    fetchData()
  }, [updateTrigger])

  const handleFollow = async () => {
    try {
      const response = await newFollower(profile_id)
      if (response && response.status === "success") {
        console.log("follow submit success");
        setUpdateTrigger(prevTrigger => prevTrigger + 1)
      }
    } catch (error) {
      if (error instanceof Error) {
        setError(error.message);
      } else {
        setError("An unexpected error occurred.");
      }
    }
  }

  const handleUnfollow = async () => {
    try {
      const response = await unfollow(profile_id, userContext.user ? userContext.user.userId : "")
      if (response && response.status === "success") {
        console.log("unfollow success");
        setUpdateTrigger(prevTrigger => prevTrigger + 1)
      }
    } catch (error) {
      if (error instanceof Error) {
        setError(error.message);
      } else {
        setError("An unexpected error occurred.");
      }
    }
  }

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
          <div style={{
            fontSize: "small",
            fontStyle: "italic",
            color: "gray"
          }}>
            {followerData.following_status == 1 ?
              <button onClick={handleUnfollow}>Unfollow</button>
              : followerData.request_pending == 1 ?
                <button onClick={handleUnfollow}>Request Pending...</button>
                : <button onClick={handleFollow}>Follow!</button>}
          </div>
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