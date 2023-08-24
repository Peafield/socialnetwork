import React from 'react'

interface CommentContentProps {
    content: string,
    image: string
}

const CommentContent: React.FC<CommentContentProps> = ({
    content,
    image
}) => {
  return (
    <div>{content}</div>
  )
}

export default CommentContent