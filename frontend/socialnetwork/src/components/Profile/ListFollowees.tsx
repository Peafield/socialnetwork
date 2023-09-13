import React, { useEffect, useState } from 'react'
import { FollowerProps } from './ProfileHeader'
import styles from './Profile.module.css'
import { Link } from 'react-router-dom'
import { getUserByUserID } from '../../controllers/GetUser'
import { ProfileProps } from './Profile'


interface ListFolloweesProps {
    followees: FollowerProps[]
}

const ListFollowees: React.FC<ListFolloweesProps> = ({
    followees
}) => {

    const [followeeNames, setFolloweeNames] = useState<string[]>([])

    useEffect(() => {
        const fetchData = async () => {
            const names = await Promise.all(followees.map(async (followee) => {
                const user: ProfileProps = await getUserByUserID(followee.followee_id)
                return user.display_name
            }))

            setFolloweeNames(names)
        }

        fetchData();
    }, [followees])

    return (
        <div
            className={styles.listUsers}>
            {followeeNames.map((name) => (
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

export default ListFollowees