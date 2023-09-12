import React, { ChangeEvent, MouseEventHandler, useContext, useEffect, useState } from "react";
import { FaUserCircle, FaUserEdit } from "react-icons/fa";
import { Link } from "react-router-dom";
import { UserContext } from "../../context/AuthContext";
import { useWebSocketContext } from "../../context/WebSocketContext";
import { handleAPIRequest } from "../../controllers/Api";
import { getFollowerData } from "../../controllers/Follower/GetFollower";
import { newFollower } from "../../controllers/Follower/NewFollower";
import { unfollow } from "../../controllers/Follower/Unfollow";
import { createImageURL } from "../../controllers/ImageURL";
import { getCookie } from "../../controllers/SetUserContextAndCookie";
import { WebSocketReadMessage } from "../../Socket";
import Container from "../Containers/Container";
import styles from "./Profile.module.css";

interface ProfileHeaderProps {
  profile_id: string,
  first_name: string,
  last_name: string,
  email: string,
  display_name: string;
  avatar: string;
  num_of_posts: number;
  followers: number;
  following: number;
  dob: string
  about_me: string;
  is_private: boolean;
  is_own_profile: boolean;
  profileTab: string;
  getProfileTab: (tab: string) => void;
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
  email,
  display_name,
  avatar,
  num_of_posts,
  followers,
  following,
  dob,
  about_me,
  is_private,
  is_own_profile,
  profileTab,
  getProfileTab
}) => {
  const userContext = useContext(UserContext);
  const { message, sendMessage } = useWebSocketContext();
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
    if (!is_own_profile) {
      fetchData()
    }
  }, [updateTrigger])

  const handleFollow = async () => {
    try {
      const response = await newFollower(profile_id)
      if (response && response.status === "success") {
        console.log("follow submit success");
        setUpdateTrigger(prevTrigger => prevTrigger + 1)

        const action = is_private ? "request" : "follow"

        const messageToSend: WebSocketReadMessage = {
          type: "notification",
          info: {
            receiver: profile_id,
            action_type: action
          }
        }

        sendMessage(messageToSend)
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

  const handleTabChange = (e: React.MouseEvent<HTMLButtonElement>) => {
    getProfileTab(e.currentTarget.value)
  }

  return (
    <>
      <div className={styles.profileheadercontainer}>
        <div className={styles.displaypicturecontainer}>
          {(profilePicUrl && (!is_private || is_own_profile) && (
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
          {is_private && !is_own_profile ?
            <div>This profile is private</div>
            : <div>{first_name} {last_name}</div>}
          <div style={{ color: "gray" }}>{display_name}</div>
          <div style={{
            fontSize: "small",
            fontStyle: "italic",
            color: "gray",
            width: 'auto'
          }}>
            {!is_own_profile ?
              followerData.following_status == 1 ?
                <button onClick={handleUnfollow}>Unfollow</button>
                : followerData.request_pending == 1 ?
                  <button onClick={handleUnfollow}>Request Pending...</button>
                  : <button onClick={handleFollow}>Follow!</button>
              : <div>
                <span
                  style={{
                    fontSize: "300%",
                  }}>

                  <Link to={userContext.user ? "/dashboard/user/edit/" + userContext.user.displayName : ""}>
                    <FaUserEdit />
                  </Link>
                </span>
              </div>}
          </div>
        </div>
        <div className={styles.otherprofileinfocontainer}>
          <div className={styles.profilestatscontainer}>
            <div style={{ display: 'flex', alignItems: 'center' }}>{num_of_posts}<button onClick={handleTabChange} value="posts" style={profileTab == "posts" ? { textDecorationLine: 'underline' } : undefined}> Posts</button></div>
            <div style={{ display: 'flex', alignItems: 'center' }}>{followers}<button onClick={handleTabChange} value="followers" style={profileTab == "followers" ? { textDecorationLine: 'underline' } : undefined}> Followers</button></div>
            <div style={{ display: 'flex', alignItems: 'center' }}><button onClick={handleTabChange} value="followees" style={profileTab == "followees" ? { textDecorationLine: 'underline' } : undefined}>Following </button>{following}</div>

          </div>
          <div className={styles.aboutmecontainer}>
            <div>{!is_private || is_own_profile ? about_me : null}</div>
          </div>
          {is_own_profile ?
            <div className={styles.personaldetails}>
              <div>{email}</div>
              <div>{dob.split("T00:00:00Z")}</div>
              <div>{is_private ? "Private Profile" : "Public Profile"}</div>
            </div>
            : null}
        </div>
      </div>
    </>
  );
};

export default ProfileHeader