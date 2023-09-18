import React, { ChangeEvent, useEffect, useState } from 'react'
import { AiOutlineArrowRight } from 'react-icons/ai'
import { Link, useNavigate } from 'react-router-dom'
import { handleAPIRequest } from '../../controllers/Api'
import { getUserByUserID } from '../../controllers/GetUser'
import { GetAllGroups } from '../../controllers/Group/GetGroup'
import { ProfileProps } from '../Profile/Profile'
import CreateGroup, { GroupProps } from './CreateGroup'
import styles from './Group.module.css'



const GroupsList = () => {
    const navigate = useNavigate()
    const [groups, setGroups] = useState<GroupProps[] | null>(null)
    const [error, setError] = useState<string | null>(null)

    useEffect(() => {
        const fetchData = async () => {
            try {
                const groupsData = await GetAllGroups()

                setGroups(groupsData)
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
        }

        fetchData()
    }, [])

    useEffect(() => {
        const fetchData = async () => {
            if (groups) {
                const groupsWithNames = await Promise.all(groups.map(async (group) => {
                    const user: ProfileProps = await getUserByUserID(group.creator_id)
                    group.creator_name = user.display_name
                    return group
                }))

                setGroups(groupsWithNames)
            }

        }

        fetchData();
    }, [groups])

    return (
        <div
            className={styles.groupslistcontainer}>
            <CreateGroup />
            <div
                className={styles.listofgroups}>
                {groups ?
                    groups.map((group) => (
                        <div
                            key={group.group_id}
                            className={styles.grouplink}>
                            <div>Group Name: {group.title}</div>
                            <div>Created By: {group.creator_name}</div>
                            <Link to={`/dashboard/group/${group.title}`}>
                                <span style={{ color: "#fa4d6a" }}>
                                    <AiOutlineArrowRight />
                                </span>
                            </Link>
                        </div>
                    ))
                    :
                    null
                }
            </div>
        </div>
    )
}

export default GroupsList