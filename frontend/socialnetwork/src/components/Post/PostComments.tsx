import React from 'react'

interface PostCommentsProps {
  post_id: string
}

const PostComments: React.FC<PostCommentsProps> = ({
  post_id
}) => {
  return (
    <div>PostComments</div>
  )
}

export default PostComments