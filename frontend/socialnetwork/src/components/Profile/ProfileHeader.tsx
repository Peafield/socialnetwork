import React from 'react'
import Container from '../Containers/Container'

interface ProfileHeaderProps {
    display_name: string,
    num_of_posts: number,
    followers: number,
    following: number
}

const ProfileHeader: React.FC<ProfileHeaderProps> = ({
    display_name,
    num_of_posts,
    followers,
    following
}) => {
  return (
    <Container>
        <div>
            <div>
                {display_name}
            </div>
            <div>
                {num_of_posts}
            </div>
            <div>
                {followers}
            </div>
            <div>
                {following}
            </div>
        </div>
    </Container>
  )
}

export default ProfileHeader