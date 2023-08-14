import React from 'react'

interface PostActionsProps {
    likes: number,
    dislikes: number,
    numOfComments: number
}

const PostActions: React.FC<PostActionsProps> = ({
    likes,
    dislikes,
    numOfComments
}) => {
  return (
    <>
    <p>Likes: {likes}</p>
    <p>Dislikes: {dislikes}</p>
    <p>Number of Comments: {numOfComments}</p>
    </>
  )
}

export default PostActions