import React, { useContext, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { UserContext } from "../../context/AuthContext";
import { handleAPIRequest } from "../../controllers/Api";
import {
  getFollowers,
  getFollowees,
} from "../../controllers/Follower/GetFollower";
import { getUserByUserID } from "../../controllers/GetUser";
import { GetUserGroups } from "../../controllers/Group/GetGroup";
import { getCookie } from "../../controllers/SetUserContextAndCookie";
import Container from "../Containers/Container";
import Snackbar from "../feedback/Snackbar";
import { GroupProps } from "../Group/Group";
import { ProfileProps } from "../Profile/Profile";
import { FollowerProps } from "../Profile/ProfileHeader";
import styles from "./Post.module.css";

interface CreatePostFormData {
  content: string;
  group_id: string
  image_path: string;
  privacy_level: number;
  selected_profiles: string[];
}

const CreatePost: React.FC = () => {
  const navigate = useNavigate();
  const userContext = useContext(UserContext);
  const [formData, setFormData] = useState<CreatePostFormData>({
    content: "",
    group_id: "",
    image_path: "",
    privacy_level: 0,
    selected_profiles: [],
  });
  const [showGroups, setShowGroups] = useState(false);
  const [selectableGroups, setSelectableGroups] = useState<GroupProps[]>([])
  const [showFollowers, setShowFollowers] = useState(false);
  const [selectableProfiles, setSelectableProfiles] = useState<ProfileProps[]>(
    []
  );
  const [error, setError] = useState<string | null>(null);
  const [snackbarOpen, setSnackbarOpen] = useState<boolean>(false);
  const [snackbarType, setSnackbarType] = useState<
    "success" | "error" | "warning"
  >("error");

  useEffect(() => {
    const fetchData = async () => {
      try {
        if (userContext.user) {
          const profiles = await getSelectableProfiles(userContext.user?.userId)
          const groups = await getSelectableGroups(userContext.user?.userId)
          setSelectableProfiles(profiles);
          setSelectableGroups(groups);
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
  }, []);

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const { name, value } = e.target;
    console.log(name, value);

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
      switch (name) {
        case "group_id":
          setFormData((prevState) => ({
            ...prevState,
            group_id: value,
          }));
          break;
        case "privacy_level":
          setFormData((prevState) => ({
            ...prevState,
            [name]: Number(value),
          }));
          setShowFollowers(name === "privacy_level" && value === "2");
          setShowGroups(name === "privacy_level" && value === "3");
          break;
        case "selected_profiles":
          let sp = formData.selected_profiles;
          sp.includes(value) ? sp.splice(sp.indexOf(value), 1) : sp.push(value);
          setFormData((prevState) => ({
            ...prevState,
            selected_profiles: sp,
          }));
          break;
        default:
          setFormData((prevState) => ({
            ...prevState,
            [name]: value,
          }));
          break;
      }
    }
    console.log(formData);

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
      const response = await handleAPIRequest("/post", options);
      if (response && response.status === "success") {
        setSnackbarType("success");
        setSnackbarOpen(true);
        setTimeout(() => {
          navigate("/dashboard");
        }, 1000);
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
    <Container>
      <div className={styles.createpostcontainer}>
        <div className={styles.formwrapper}>
          <form onSubmit={handleSubmit}>
            <div className={styles.inputgroup}>
              <textarea
                placeholder="Write something for your post!"
                value={formData.content}
                name="content"
                onChange={handleChange}
              />
            </div>
            <div className={styles.inputgroup}>
              <label htmlFor="image_path">
                Include a picture?
                <input
                  type="file"
                  id="image_path"
                  name="image_path"
                  onChange={handleChange}
                />
              </label>
            </div>
            <div className={styles.inputgrouprow}>
              <input
                className={styles.input}
                id="public_privacy_level"
                type="radio"
                value={0}
                defaultChecked
                name="privacy_level"
                onChange={handleChange}
              />
              <label htmlFor="public_privacy_level">Public</label>
              <input
                className={styles.input}
                id="private_privacy_level"
                type="radio"
                value={1}
                name="privacy_level"
                onChange={handleChange}
              />
              <label htmlFor="private_privacy_level">Private</label>
              <input
                className={styles.input}
                id="selected_privacy_level"
                type="radio"
                value={2}
                name="privacy_level"
                onChange={handleChange}
              />
              <label htmlFor="selected_privacy_level">Selected</label>
              <input
                className={styles.input}
                type="radio"
                id="group_privacy_level"
                value={3}
                name="privacy_level"
                onChange={handleChange}
              />
              <label htmlFor="group_privacy_level">Group</label>
            </div>
            {showFollowers && (
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
                    <label htmlFor={profile.display_name}>
                      {profile.display_name}{" "}
                      {`(${profile.first_name} ${profile.last_name})`}
                    </label>
                  </div>
                ))}
              </div>
            )}
            {showGroups && (
              <div className={styles.selectablegroupscontainer}>
                {selectableGroups.map((group) => (
                  <div key={group.group_id} className={styles.checkbox}>
                    <input
                      type="radio"
                      id={group.group_id}
                      name="group_id"
                      checked={formData.group_id === group.group_id}
                      onChange={handleChange}
                      value={group.group_id}
                    />
                    <label htmlFor={group.group_id}>
                      {group.title}
                    </label>
                  </div>
                ))}
              </div>
            )}
            <div className={styles.inputgroup}>
              <button type="submit">Create Post</button>
            </div>
          </form>
        </div>
      </div >
      <Snackbar
        open={snackbarOpen}
        onClose={() => {
          setSnackbarOpen(false);
          setError(null);
        }}
        message={error ? error : "Post Sucessfully Created!"}
        type={snackbarType}
      />
    </Container >
  );
};

export default CreatePost;

async function getSelectableProfiles(userId: string) {
  const followdataFollowers = await getFollowers(
    userId
  );
  const followdataFollowees = await getFollowees(
    userId
  );

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

  return mergedProfiles

}

async function getSelectableGroups(userId: string) {
  const userGroups: GroupProps[] = await GetUserGroups(userId)

  return userGroups
}