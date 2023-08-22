import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { handleAPIRequest } from '../../controllers/Api'
import { getCookie } from '../../controllers/SetUserContextAndCookie'
import Container from '../Containers/Container'
import ProfileHeader from './ProfileHeader'
import ProfilePostsGrid from './ProfilePostsGrid'
import styles from './Profile.module.css'
import { getUserByDisplayName } from '../../controllers/GetUser'

export interface ProfileProps {
    user_id: string,
    display_name: string,
    avatar: string
    first_name: string,
    last_name: string,
    dob: string,
    about_me: string
}

const Profile: React.FC = () => {
    const [profile, setProfile] = useState<ProfileProps | null>(null)
    const [profileLoading, setProfileLoading] = useState<boolean>(false);
    const [error, setError] = useState<string | null>(null);
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
                  } else {
                    setError("An unexpected error occurred.");
                  }
            }
            setProfileLoading(false);
        };

        fetchData(); // Call the async function
    }, [username]);

    if (profileLoading) {return <p>Loading...</p>}

  return (
    <>
    {profile? <div className={styles.profilecontainer}>
        <ProfileHeader first_name={profile.first_name} last_name={profile.last_name} display_name={profile.display_name} avatar={profile.avatar} num_of_posts={0} followers={0} following={0} about_me={profile.about_me}/>
        <ProfilePostsGrid user_id={profile.user_id}/>
        </div> : null}
          
    </>
  )
}

export default Profile