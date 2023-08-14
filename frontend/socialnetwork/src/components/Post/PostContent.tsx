import React from 'react'

interface PostContentProps {
    text: string
}

const PostContent: React.FC<PostContentProps> = ({
    text
}) => {
  return (
    <>
    <p>{text}</p>
    </>
  )
}

export default PostContent