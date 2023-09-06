import React, { useEffect, useState } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import { handleAPIRequest } from '../../controllers/Api'
import { getGroupByName } from '../../controllers/Group/GetGroup'
import { getCookie } from '../../controllers/SetUserContextAndCookie'
import { PostProps } from '../Post/Post'
import GroupEvents from './GroupEvents'
import GroupHeader from './GroupHeader'
import GroupMembers from './GroupMembers'
import GroupPosts from './GroupPosts'

export interface GroupProps {
    group_id: string,
    title: string,
    description: string,
    creator_id: string,
    creation_date: string
}

const Group: React.FC = () => {
    const navigate = useNavigate();
    const [group, setGroup] = useState<GroupProps | null>(null)
    const [groupLoading, setGroupLoading] = useState(false)
    const [groupPosts, setGroupPosts] = useState<PostProps[]>([])
    const [groupPostsLoading, setGroupPostsLoading] = useState(false)
    const [error, setError] = useState<string | null>(null)
    const { groupname } = useParams();

    useEffect(() => {
        const fetchData = async () => {
            setGroupLoading(true);

            try {
                if (groupname) {
                    const newgroup = await getGroupByName(groupname)
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

            { group ? url = `/post?group_id=${encodeURIComponent(group.group_id)}` : url = "/post" }

            const options = {
                method: "GET",
                headers: {
                    Authorization: "Bearer " + getCookie("sessionToken"),
                    "Content-Type": "application/json",
                },
            };
            try {
                const response = await handleAPIRequest(url, options);
                setGroupPosts(response.data.Posts);
                console.log(response.data.Posts);

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

    return (
        <>
            {group ?
                <div>
                    <GroupHeader title={group.title} description={group.description} creator_id={group.creator_id} />
                    <GroupMembers />
                    <GroupEvents />
                    <GroupPosts posts={groupPosts} />
                </div>
                : null}
        </>
    )
}

export default Group