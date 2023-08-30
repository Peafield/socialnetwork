import React, { useContext, useEffect, useState } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import { handleAPIRequest } from '../../controllers/Api'
import { getCookie } from '../../controllers/SetUserContextAndCookie'
import ProfileHeader, { FollowerProps } from './ProfileHeader'
import ProfilePostsGrid from './ProfilePostsGrid'
import styles from './Profile.module.css'
import { getUserByDisplayName } from '../../controllers/GetUser'
import { PostProps } from '../Post/Post'
import { getFollowees, getFollowerData, getFollowers } from '../../controllers/Follower/GetFollower'
import { UserContext } from '../../context/AuthContext'

export interface ProfileProps {
    user_id: string,
    display_name: string,
    avatar: string
    first_name: string,
    last_name: string,
    dob: string,
    about_me: string,
    is_private: string
}

const Profile: React.FC = () => {
    const navigate = useNavigate();
    const userContext = useContext(UserContext)
    const [profile, setProfile] = useState<ProfileProps | null>(null)
    const [profileLoading, setProfileLoading] = useState<boolean>(false);
    const [profilePosts, setProfilePosts] = useState<PostProps[]>([]);
    const [postsLoading, setPostsLoading] = useState<boolean>(false);
    const [followerData, setFollowerData] = useState<FollowerProps>({
        follower_id: "",
        followee_id: "",
        following_status: 0,
        request_pending: 0,
        creation_date: ""
    })
    const [followers, setFollowers] = useState<FollowerProps[]>([])
    const [followees, setFollowees] = useState<FollowerProps[]>([])
    const [error, setError] = useState<string | null>(null);
    const { username } = useParams();

    useEffect(() => {
        const fetchData = async () => {
            try {
                const following = await getFollowerData(profile ? profile.user_id : "")
                setFollowerData(following)

            } catch (error) {
                if (error instanceof Error) {
                    setError(error.message);
                } else {
                    setError("An unexpected error occurred.");
                }
            }
        }
        if (profile && userContext.user?.userId != profile.user_id) {
            fetchData()
        }
    }, [profile])

    useEffect(() => {
        const fetchData = async () => {
            try {
                const following = await getFollowers(profile ? profile.user_id : "")
                
                setFollowers(following.Followers)

            } catch (error) {
                if (error instanceof Error) {
                    setError(error.message);
                } else {
                    setError("An unexpected error occurred.");
                }
            }
        }
        if (profile) {
            fetchData()
        }
    }, [profile])

    useEffect(() => {
        const fetchData = async () => {
            try {
                const following = await getFollowees(profile ? profile.user_id : "")             

                setFollowees(following.Followers)

            } catch (error) {
                if (error instanceof Error) {
                    setError(error.message);
                } else {
                    setError("An unexpected error occurred.");
                }
            }
        }
        if (profile) {
            fetchData()
        }
    }, [profile])

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
        const fetchUserPostData = async () => {
            setPostsLoading(true);

            let url

            if (profile) {
                url = `/post?user_id=${encodeURIComponent(profile.user_id)}`
            } else {
                return
            }

            const options = {
                method: "GET",
                headers: {
                    Authorization: "Bearer " + getCookie("sessionToken"),
                    "Content-Type": "application/json",
                },
            };
            try {
                const response = await handleAPIRequest(url, options);
                if (response.data.Posts) {
                    setProfilePosts(response.data.Posts);
                }
            } catch (error) {
                if (error instanceof Error) {
                    setError(error.message);
                } else {
                    setError("An unexpected error occurred.");
                }
            }
            setPostsLoading(false);
        };

        fetchUserPostData(); // Call the async function
    }, [profile]);    

    if (profileLoading) { return <p>Loading...</p> }

    return (
        <>
            {profile ? <div className={styles.profilecontainer}>
                <ProfileHeader
                    profile_id={profile.user_id}
                    first_name={profile.first_name}
                    last_name={profile.last_name}
                    display_name={profile.display_name}
                    avatar={profile.avatar}
                    num_of_posts={profilePosts.length}
                    followers={followers.length}
                    following={followees.length}
                    about_me={profile.about_me} 
                    is_private={followerData.following_status == 1 || !profile.is_private ? false : true}
                    is_own_profile={userContext.user?.userId == profile.user_id ? true : false}/>
                <ProfilePostsGrid
                    user_id={profile.user_id}
                    posts={profilePosts}
                    is_private={followerData.following_status == 1 || userContext.user?.userId == profile.user_id || !profile.is_private ? false : true} />
            </div>
                : null}

        </>
    )
}

export default Profile