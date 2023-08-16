import React from 'react'
import Container from '../Containers/Container'

interface ProfileBioProps {
    bio: string
}

const ProfileBio: React.FC<ProfileBioProps> = ({
    bio
}) => {
  return (
    <Container>
        <div>
            {bio}
        </div>
    </Container>
  )
}

export default ProfileBio