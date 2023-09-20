import { group } from "console";
import React, { useContext, useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { UserContext } from "../../context/AuthContext";
import { useWebSocketContext } from "../../context/WebSocketContext";
import { handleAPIRequest } from "../../controllers/Api";
import {
  getFollowers,
  getFollowees,
} from "../../controllers/Follower/GetFollower";
import { getUserByUserID } from "../../controllers/GetUser";
import { GetGroupMembers } from "../../controllers/Group/GetGroupMembers";
import { getCookie } from "../../controllers/SetUserContextAndCookie";
import { WebSocketReadMessage } from "../../Socket";
import Snackbar from "../feedback/Snackbar";
import { ProfileProps } from "../Profile/Profile";
import { FollowerProps } from "../Profile/ProfileHeader";
import styles from "./Group.module.css";

export interface GroupMemberProps {
  group_id: string;
  member_id: string;
  member_name: string;
  request_pending: number;
  permission_level: number;
  creation_date: string;
}

interface MembersProps {
  members: GroupMemberProps[] | null;
  isUserMember: boolean
}

interface InviteMemberFormData {
  inviter_id: string;
  invitees_ids: string[];
  group_id: string;
}

const GroupMembers: React.FC<MembersProps> = ({ members, isUserMember }) => {
  const userContext = useContext(UserContext);
  const navigate = useNavigate();
  const { message, sendMessage } = useWebSocketContext();
  let messageToSend: WebSocketReadMessage = {
    type: "",
    info: ""
  }
  const [formData, setFormData] = useState<InviteMemberFormData>({
    inviter_id: "",
    invitees_ids: [],
    group_id: members ? members[0].group_id : "",
  });
  const [selectableProfiles, setSelectableProfiles] = useState<ProfileProps[]>(
    []
  );
  const [updateTrigger, setUpdateTrigger] = useState<number>(0)
  const [error, setError] = useState<string | null>(null);
  const [snackbarOpen, setSnackbarOpen] = useState<boolean>(false);
  const [snackbarType, setSnackbarType] = useState<
    "success" | "error" | "warning"
  >("error");
  const { groupname } = useParams();

  useEffect(() => {
    const fetchData = async () => {
      try {
        if (userContext.user) {
          const followdataFollowers = await getFollowers(
            userContext.user?.userId
          );
          const followdataFollowees = await getFollowees(
            userContext.user?.userId
          );

          let allMembers: GroupMemberProps[]

          if (members && members[0]) {
            allMembers = await GetGroupMembers(members[0].group_id)
          }

          const followerUsersPromises = followdataFollowers.Followers.map(
            async (follower: FollowerProps) => {
              const user: ProfileProps = await getUserByUserID(
                follower.follower_id
              );
              return user;
            }
          );

          const followeeUsersPromises = followdataFollowees.Followers.map(
            async (follower: FollowerProps) => {
              const user: ProfileProps = await getUserByUserID(
                follower.followee_id
              );
              return user;
            }
          );

          // Use Promise.all to await all promises and get resolved users
          const followerUsers = await Promise.all(followerUsersPromises);
          const followeeUsers = await Promise.all(followeeUsersPromises);

          const profiles = [...followerUsers, ...followeeUsers];

          // Function to filter out duplicates based on user_id
          const uniqueProfiles = (array: any[]) => {
            const seen = new Set();
            return array.filter((item) => {
              if (seen.has(item.user_id)) {
                return false;
              }
              seen.add(item.user_id);
              return true;
            });
          };

          const mergedProfiles = uniqueProfiles(profiles);

          const finalProfiles = mergedProfiles.filter((profile) => {
            if (allMembers.some((member) => member.member_id === profile.user_id)) {
              return false
            }
            return true
          })

          setSelectableProfiles(finalProfiles);
        }
      } catch (error) {
        if (error instanceof Error) {
          setError(error.message);
          if (error.cause === 401) {
            navigate("/signin");
          }
        } else {
          setError("An unexpected error occurred.");
        }
      }
    };

    fetchData(); // Call the async function
  }, [updateTrigger]);

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const { name, value } = e.target;
    if (e.target.type === "file") {
      const file = (e.target as HTMLInputElement)?.files?.[0] || null;
      if (file) {
        const reader = new FileReader();
        reader.onloadend = () => {
          setFormData((prevState) => ({
            ...prevState,
            image_path: reader.result as string,
          }));
        };
        reader.readAsDataURL(file);
      }
    } else {
      let sp = formData.invitees_ids;
      sp.includes(value) ? sp.splice(sp.indexOf(value), 1) : sp.push(value);
      setFormData((prevState) => ({
        ...prevState,
        invitees_ids: sp,
      }));
    }
  };

  const handleSubmit = async (e: { preventDefault: () => void }) => {
    e.preventDefault();
    const data = { data: formData };
    const options = {
      method: "POST",
      headers: {
        Authorization: "Bearer " + getCookie("sessionToken"),
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    };
    try {
      const response = await handleAPIRequest("/groupmembers", options);
      if (response && response.status === "success") {
        setSnackbarType("success");
        setSnackbarOpen(true);
        formData.invitees_ids.forEach((invitee) => {
          messageToSend = {
            type: "notification",
            info: {
              receiver: invitee,
              group_id: formData.group_id,
              action_type: "invite"
            }
          }
          sendMessage(messageToSend)
          setUpdateTrigger((prev: number) => {
            return prev + 1
          })
        })
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
    <div className={styles.groupmemberscontainer}>
      {isUserMember ?
        <div className={styles.invitemembercontainer}>
          <form onSubmit={handleSubmit}>
            {selectableProfiles ? (
              <div className={styles.selectableprofilescontainer}>
                {selectableProfiles.map((profile) => (
                  <div key={profile.display_name} className={styles.checkbox}>
                    <input
                      type="checkbox"
                      id={profile.display_name}
                      name="selected_profiles"
                      onChange={handleChange}
                      value={profile.user_id}
                    />
                    <label htmlFor={profile.display_name} style={{ textTransform: 'none' }}>
                      {profile.display_name}{" "}
                      {`(${profile.first_name} ${profile.last_name})`}
                    </label>
                  </div>
                ))}
              </div>
            ) : null}
            <div className={styles.submit}>
              <button type="submit">Invite Users</button>
            </div>
          </form>
        </div>
        :
        null}
      <div className={styles.groupmemberslist}>
        {members
          ? members.map((member) => (
            <div key={member.member_id} className={styles.groupmember}>
              {member.member_name}
            </div>
          ))
          : null}
      </div>
      <Snackbar
        open={snackbarOpen}
        onClose={() => {
          setSnackbarOpen(false);
          setError(null);
        }}
        message={error ? error : "User/s Invited!"}
        type={snackbarType}
      />
    </div>
  );
};

export default GroupMembers;
