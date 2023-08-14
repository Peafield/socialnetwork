import React from 'react'

interface PostHeaderProps {
  headerText: string,
  creatorId: string
  //size: number
}

const PostHeader: React.FC<PostHeaderProps> = ({
  headerText,
  creatorId
}) => {
  return (
    <>
      <div>
        <h1>{headerText}</h1>
        <h2>{creatorId}</h2>
      </div>
    </>
  )
}

export default PostHeader