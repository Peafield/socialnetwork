import React, { useContext, useEffect, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { UserContext } from '../../context/AuthContext';
import { handleAPIRequest } from '../../controllers/Api';
import { getUserByDisplayName } from '../../controllers/GetUser';
import { getCookie, useSetUserContextAndCookie } from '../../controllers/SetUserContextAndCookie';
import styles from './Profile.module.css'

interface EditProfileFormData {
  email: string;
  display_name: string;
  new_password: string;
  confirm_password: string;
  first_name: string;
  last_name: string;
  dob: string;
  avatar_path: string;
  about_me: string;
  is_private: number
  old_password: string
}

interface OwnProfileProps {
  user_id: string,
  email: string,
  display_name: string,
  avatar: string
  first_name: string,
  last_name: string,
  dob: string,
  about_me: string,
  is_private: number
}

const EditProfile: React.FC = () => {
  const navigate = useNavigate();
  const setUserContextAndCookie = useSetUserContextAndCookie();
  const userContext = useContext(UserContext)
  const [profile, setProfile] = useState<OwnProfileProps | null>(null)
  const [profileLoading, setProfileLoading] = useState<boolean>(false);
  const [formData, setFormData] = useState<EditProfileFormData>({
    email: "",
    display_name: "",
    new_password: "",
    confirm_password: "",
    first_name: "",
    last_name: "",
    dob: "",
    avatar_path: "",
    about_me: "",
    is_private: 0,
    old_password: ""
  });
  const [error, setError] = useState<string | null>(null)
  const { username } = useParams();

  useEffect(() => {
    const fetchData = async () => {
      setProfileLoading(true);

      try {
        if (username) {
          const newprofile = await getUserByDisplayName(username)
          setProfile(newprofile)
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
      setProfileLoading(false);
    };

    fetchData(); // Call the async function
  }, [username]);

  useEffect(() => {
    if (profile) {
      setFormData({
        email: profile.email,
        display_name: profile.display_name,
        new_password: "",
        confirm_password: "",
        first_name: profile.first_name,
        last_name: profile.last_name,
        dob: profile.dob,
        avatar_path: profile.avatar,
        about_me: profile.about_me,
        is_private: profile.is_private,
        old_password: ""
      })
    }

  }, [profile])

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target;

    if (e.target.type === "file") {
      const file = (e.target as HTMLInputElement)?.files?.[0] || null;
      if (file) {
        const reader = new FileReader();
        reader.onloadend = () => {
          setFormData((prevState) => ({
            ...prevState,
            avatar_path: reader.result as string,
          }));
        }
        reader.readAsDataURL(file)
      }

    } else {
      setFormData((prevState) => ({
        ...prevState,
        [name]: value,
      }));
    }

    console.log(formData);

  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const data = { data: formData };
    console.log(formData);

    const options = {
      method: "PUT",
      headers: {
        Authorization: "Bearer " + getCookie("sessionToken"),
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    };
    try {
      const response = await handleAPIRequest("/user", options);
      if (response && response.status === "success") {
        setUserContextAndCookie(response.data);
        navigate("/dashboard/user/" + username);
      }
    } catch (error) {
      if (error instanceof Error) {
        setError(error.message);
      } else {
        setError("An unexpected error occurred.");
      }
    }
  };

  return (
    <div className={styles.editprofilecontainer}>
      <h2>Edit Profile</h2>
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor='isPrivate'>Private Profile?</label>
          <input
            type="checkbox"
            id='isPrivate'
            name='is_private'
            checked={formData.is_private === 1}
            onChange={(e) => {
              const newValue = e.target.checked ? 1 : 0;
              setFormData((prevState) => ({
                ...prevState,
                is_private: newValue,
              }));
            }}
          />

        </div>
        <div>
          <label htmlFor="firstName">First Name:</label>
          <input
            type="text"
            id="firstName"
            name='first_name'
            value={formData.first_name}
            onChange={handleChange}
          />
        </div>
        <div>
          <label htmlFor="lastName">Last Name:</label>
          <input
            type="text"
            id="lastName"
            name='last_name'
            value={formData.last_name}
            onChange={handleChange}
          />
        </div>
        <div>
          <label htmlFor="displayName">Display Name:</label>
          <input
            type="text"
            id="displayName"
            name='display_name'
            value={formData.display_name}
            onChange={handleChange}
          />
        </div>
        <div>
          <label htmlFor="email">Email:</label>
          <input
            type="text"
            id="email"
            name='email'
            value={formData.email}
            onChange={handleChange}
          />
        </div>
        <div>
          <label htmlFor="newPassword">New Password:</label>
          <input
            type="password"
            id="newPassword"
            name='new_password'
            value={formData.new_password}
            onChange={handleChange}
          />
        </div>
        <div>
          <label htmlFor="confirmPassword">Confirm Password:</label>
          <input
            type="password"
            id="confirmPassword"
            name='confirm_password'
            value={formData.confirm_password}
            onChange={handleChange}
          />
        </div>
        <div>
          <label htmlFor="dob">
            Date of Birth:
            <input
              type="date"
              id="dob"
              name="dob"
              value={formData.dob}
              onChange={handleChange}
            />
          </label>
        </div>
        <div>
          <label htmlFor="aboutMe">About Me:</label>
          <textarea
            id="aboutMe"
            name='about_me'
            value={formData.about_me}
            onChange={handleChange}
          />
        </div>
        <div>
          <label htmlFor="profilePicture">Profile Picture:</label>
          <input
            type="file"
            id="profilePicture"
            name='avatar_path'
            accept="image/*"
            onChange={handleChange}
          />
          {formData.avatar_path && (
            <img src={formData.avatar_path} alt="Profile" style={{
              maxWidth: '100px',
              verticalAlign: 'middle',
              width: '160px',
              height: '160px',
              borderRadius: '50%',
            }} />
          )}
        </div>
        <div>
          <label htmlFor="oldPassword">Current Password:</label>
          <input
            required
            type="password"
            id="oldPassword"
            name='old_password'
            value={formData.old_password}
            onChange={handleChange}
          />
        </div>
        <button type="submit">
          Save Changes
        </button>
      </form>
    </div>
  );
};

export default EditProfile;
