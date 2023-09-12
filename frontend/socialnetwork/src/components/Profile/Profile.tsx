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
import { getUserPosts } from '../../controllers/GetPosts'
import ListFollowees from './ListFollowees'
import ListFollowers from './ListFollowers'
import { PiMaskSadDuotone } from 'react-icons/pi'

export interface ProfileProps {
    user_id: string,
    email: string
    display_name: string,
    avatar: string
    first_name: string,
    last_name: string,
    dob: string,
    about_me: string,
    is_private: number
}

const Profile: React.FC = () => {
    const navigate = useNavigate();
    const userContext = useContext(UserContext)
    const [profile, setProfile] = useState<ProfileProps | null>(null)
    const [profileLoading, setProfileLoading] = useState<boolean>(false);
    const [profilePosts, setProfilePosts] = useState<PostProps[]>([]);
    const [followerData, setFollowerData] = useState<FollowerProps>({
        follower_id: "",
        followee_id: "",
        following_status: 0,
        request_pending: 0,
        creation_date: ""
    })
    const [followers, setFollowers] = useState<FollowerProps[]>([])
    const [followees, setFollowees] = useState<FollowerProps[]>([])
    const [currentProfileTab, setCurrentProfileTab] = useState<string>("posts")
    const [isprivate, setIsPrivate] = useState(true)
    const [error, setError] = useState<string | null>(null);
    const { username } = useParams();

    useEffect(() => {
        const fetchData = async () => {
            setProfileLoading(true);

            try {
                if (username) {
                    const newprofile = await getUserByDisplayName(username)

                    setProfile(newprofile)

                    const followdata = await getFollowerData(newprofile.user_id)
                    const followdataFollowers = await getFollowers(newprofile.user_id)
                    const followdataFollowees = await getFollowees(newprofile.user_id)
                    const userPosts = await getUserPosts(newprofile.user_id)

                    const postData = userPosts.map((post: any) => {
                        const newpost = post.PostInfo
                        newpost.image_path = post.PostPicture
                        return newpost
                    })

                    setFollowerData(followdata)
                    setFollowers(followdataFollowers.Followers)
                    setFollowees(followdataFollowees.Followers)
                    setProfilePosts(postData)

                } else {
                    setError("could not find profile username")
                }
            } catch (error) {
                if (error instanceof Error) {
                    setError(error.message);
                    if (error.cause === 401) {
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
            setIsPrivate(isPrivate(
                followerData.following_status,
                userContext.user?.userId,
                profile.user_id,
                profile.is_private))
        }
    }, [profile, followerData])


    const renderSwitch = (tab: string, profile: ProfileProps) => {
        if (!isprivate) {
            switch (tab) {
                case "posts":
                    return (
                        <ProfilePostsGrid
                            user_id={profile.user_id}
                            posts={profilePosts}
                            is_private={isprivate} />
                    )
                case "followers":
                    return <ListFollowers followers={followers} />
                case "followees":
                    return <ListFollowees followees={followees} />
                default:
                    return (
                        <ProfilePostsGrid
                            user_id={profile.user_id}
                            posts={profilePosts}
                            is_private={isprivate} />
                    )
            }
        } else {
            return (
                <div
                    style={{
                        display: "flex",
                        flexDirection: "column",
                        justifyContent: "center",
                        alignItems: "center",
                        margin: "10px",
                        padding: "10px"
                    }}>
                    <span style={{ fontSize: "300%" }}>
                        <PiMaskSadDuotone />
                    </span>
                    <div>
                        Nothing to see here
                    </div>
                </div>
            )
        }

    }

    if (profileLoading) { return <p>Loading...</p> }

    return (
        <>
            {profile ? <div className={styles.profilecontainer}>
                <ProfileHeader
                    profile_id={profile.user_id}
                    first_name={profile.first_name}
                    last_name={profile.last_name}
                    email={profile.email}
                    display_name={profile.display_name}
                    avatar={profile.avatar}
                    num_of_posts={profilePosts.length}
                    followers={followers.length}
                    following={followees.length}
                    about_me={profile.about_me}
                    dob={profile.dob}
                    is_private={followerData.following_status === 1 || !profile.is_private ? false : true}
                    is_own_profile={userContext.user?.userId === profile.user_id ? true : false}
                    profileTab={currentProfileTab}
                    getProfileTab={setCurrentProfileTab} />
                {renderSwitch(currentProfileTab, profile)}
            </div>
                : null}

        </>
    )
}

function isPrivate(followStatus: number, userId: string | undefined, profileId: string, isprivate: number) {
    return followStatus === 1 || userId === profileId || isprivate === 0 ? false : true
}

export default Profile