import React, { useEffect, useState } from 'react'
import { FollowerProps } from './ProfileHeader'
import styles from './Profile.module.css'
import { getUserByUserID } from '../../controllers/GetUser'
import { ProfileProps } from './Profile'
import { Link } from 'react-router-dom'

interface ListFollowersProps {
    followers: FollowerProps[]
}

const ListFollowers: React.FC<ListFollowersProps> = ({
    followers
}) => {
    const [followerNames, setFollowerNames] = useState<string[]>([])

    useEffect(() => {
        const fetchData = async () => {
            const names = await Promise.all(followers.map(async (follower) => {
                const user: ProfileProps = await getUserByUserID(follower.follower_id)
                return user.display_name
            }))

            setFollowerNames(names)
        }

        fetchData();
    }, [followers])

    return (
        <div
            className={styles.listUsers}>
            {followerNames.map((name) => (
                <Link
                    key={name}
                    to={"/dashboard/user/" + name}
                    style={{ width: 'auto' }}>
                    {name}
                </Link>
            ))}
        </div>
    )
}

export default ListFollowers