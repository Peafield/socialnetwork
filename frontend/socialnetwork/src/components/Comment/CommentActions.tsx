import React from 'react'

interface CommentActionsProps {
    likes: number,
    dislikes: number
}

const CommentActions: React.FC<CommentActionsProps> = ({
    likes,
    dislikes
}) => {
  return (
    <div>CommentActions</div>
  )
}

export default CommentActions