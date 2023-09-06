import React from 'react'
import { FollowerProps } from './ProfileHeader'
import styles from './Profile.module.css'


interface ListFolloweesProps {
    followees: FollowerProps[]
}

const ListFollowees: React.FC<ListFolloweesProps> = ({
    followees
}) => {


    return (
        <div
            className={styles.listUsers}>
            {followees.map((follower) => (
                <div style={{ width: 'auto' }}>
                    {follower.followee_id}
                </div>
            ))}
        </div>
    )
}

export default ListFollowees