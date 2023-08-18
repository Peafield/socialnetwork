import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { handleAPIRequest } from '../../controllers/Api'
import { getCookie } from '../../controllers/SetUserContextAndCookie'
import Container from '../Containers/Container'
import ProfileBio from './ProfileBio'
import ProfileHeader from './ProfileHeader'
import ProfilePostsGrid from './ProfilePostsGrid'
import styles from './Profile.module.css'

interface ProfileProps {
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

            let url

            {username ? url = `/user?displayName=${encodeURIComponent(username)}` : url = "/user"}
            
            const options = {
                method: "GET",
                headers: {
                    Authorization: "Bearer " + getCookie("sessionToken"),
                    "Content-Type": "application/json",
                },
                params: {
                    displayName: username
                }
            };
            try {
                const response = await handleAPIRequest(url, options);
                console.log(response.data);

                const newprofile = response.data.UserInfo
                const avatar = response.data.ProfilePic

                newprofile.avatar = avatar
                
                console.log(newprofile);
                
                setProfile(newprofile);

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
    <Container>
    {profile? <div className={styles.profilecontainer}>
        <ProfileHeader display_name={profile.display_name} avatar={profile.avatar} num_of_posts={0} followers={0} following={0}/>
        <ProfileBio bio={profile.about_me}/>  
        <ProfilePostsGrid user_id={profile.user_id}/>
        </div> : null}
          
    </Container>
  )
}

export default Profile