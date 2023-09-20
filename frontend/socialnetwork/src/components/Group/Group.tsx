import React, { useContext, useEffect, useState } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import { handleAPIRequest } from '../../controllers/Api'
import { getGroupByName } from '../../controllers/Group/GetGroup'
import { getCookie } from '../../controllers/SetUserContextAndCookie'
import { PostProps } from '../Post/Post'
import GroupEvents from './GroupEvents'
import GroupHeader from './GroupHeader'
import GroupMembers, { GroupMemberProps } from './GroupMembers'
import GroupPosts from './GroupPosts'
import styles from './Group.module.css'
import { UserContext } from '../../context/AuthContext'
import { GetGroupMembers } from '../../controllers/Group/GetGroupMembers'
import { getUserByUserID } from '../../controllers/GetUser'
import { ProfileProps } from '../Profile/Profile'
import { PiMaskSadDuotone } from 'react-icons/pi'

export interface GroupProps {
    group_id: string,
    title: string,
    description: string,
    creator_id: string,
    creator_name: string
    creation_date: string
}

const Group: React.FC = () => {
    const navigate = useNavigate();
    const userContext = useContext(UserContext)
    const [group, setGroup] = useState<GroupProps | null>(null)
    const [isUserMember, setIsUserMember] = useState(false)
    const [groupLoading, setGroupLoading] = useState(false)
    const [groupPosts, setGroupPosts] = useState<PostProps[]>([])
    const [groupMembers, setGroupMembers] = useState<GroupMemberProps[] | null>(null)
    const [groupPostsLoading, setGroupPostsLoading] = useState(false)
    const [currentGroupTab, setCurrentGroupTab] = useState<string>("posts")
    const [error, setError] = useState<string | null>(null)
    const { groupname } = useParams();

    useEffect(() => {
        const fetchData = async () => {
            setGroupLoading(true);

            try {
                if (groupname) {
                    const newgroup = await getGroupByName(groupname)
                    const members = await GetGroupMembers(newgroup.group_id)
                    const filteredMembers = members.filter((member: GroupMemberProps) => member.permission_level > 0 && member.permission_level <= 2)
                    const membersWithNames = await Promise.all(filteredMembers.map(async (member: GroupMemberProps) => {
                        if (member.member_id === userContext.user?.userId) {
                            setIsUserMember(true)
                        }
                        const user: ProfileProps = await getUserByUserID(member.member_id)
                        member.member_name = user.display_name
                        return member
                    }))
                    setGroupMembers(membersWithNames)
                    const creator: ProfileProps = await getUserByUserID(newgroup.creator_id)
                    newgroup.creator_name = creator.display_name
                    setGroup(newgroup)
                } else {
                    setError("could not find group name")
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
            setGroupLoading(false);
        };

        fetchData(); // Call the async function
    }, [groupname]);

    useEffect(() => {
        const fetchUserPostData = async () => {
            setGroupPostsLoading(true);

            let url

            if (group) {
                url = `/post?group_id=${encodeURIComponent(group.group_id)}`
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

                const posts = response.data.Posts

                const postData = posts.map((post: any) => {
                    const newpost = post.PostInfo
                    newpost.image_path = post.PostPicture
                    return newpost
                })

                setGroupPosts(postData);

            } catch (error) {
                if (error instanceof Error) {
                    setError(error.message);
                } else {
                    setError("An unexpected error occurred.");
                }
            }
            setGroupPostsLoading(false);
        };

        fetchUserPostData(); // Call the async function
    }, [group]);

    const renderSwitch = () => {
        if (isUserMember) {
            switch (currentGroupTab) {
                case "posts":
                    return <GroupPosts posts={groupPosts} isUserMember={isUserMember} />
                case "members":
                    return <GroupMembers members={groupMembers} isUserMember={isUserMember} />
                case "events":
                    return <GroupEvents group_id={group ? group.group_id : ""} />
                default:
                    return <GroupPosts posts={groupPosts} isUserMember={isUserMember} />
                    break;
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

    return (
        <>
            {group ?
                <div className={styles.grouppagecontainer}>
                    <GroupHeader
                        group_id={group.group_id}
                        title={group.title}
                        description={group.description}
                        creator_id={group.creator_id}
                        creator_name={group.creator_name}
                        isUserMember={isUserMember}
                        groupTab={currentGroupTab}
                        getGroupTab={setCurrentGroupTab} />
                    {renderSwitch()}
                </div>
                : null}
        </>
    )
}

export default Group