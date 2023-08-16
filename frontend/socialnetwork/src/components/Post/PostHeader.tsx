import React from 'react'
import { Link } from 'react-router-dom'

interface PostHeaderProps {
  headerText: string,
  creatorId: string
  creatorDisplayName: string,
  //size: number
}

const PostHeader: React.FC<PostHeaderProps> = ({
  headerText,
  creatorId,
  creatorDisplayName
}) => {
  return (
    <>
      <div>
        <h1>{headerText}</h1>
        <Link to={"/dashboard/user/" + creatorDisplayName} >{creatorDisplayName}</Link>
      </div>
    </>
  )
}

export default PostHeader