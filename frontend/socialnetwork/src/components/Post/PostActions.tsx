import React from 'react'
import styles from './Post.module.css'

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
    <div className={styles.postactionscontainer}>
      <p>Likes: {likes}</p>
      <p>Dislikes: {dislikes}</p>
      <p># of Comments: {numOfComments}</p>
    </div>
  )
}

export default PostActions