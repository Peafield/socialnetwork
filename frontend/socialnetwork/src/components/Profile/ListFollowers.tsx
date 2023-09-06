import React from 'react'
import { FollowerProps } from './ProfileHeader'
import styles from './Profile.module.css'

interface ListFollowersProps {
    followers: FollowerProps[]
}

const ListFollowers: React.FC<ListFollowersProps> = ({
    followers
}) => {


    return (
        <div
            className={styles.listUsers}>
            {followers.map((follower) => (
                <div>
                    {follower.follower_id}
                </div>
            ))}
        </div>
    )
}

export default ListFollowers