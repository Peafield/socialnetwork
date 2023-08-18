import React from 'react'
import { Link } from 'react-router-dom'
import Container from '../Containers/Container'
import styles from './Post.module.css'

interface PostHeaderProps {
  creatorId: string
  creatorDisplayName: string,
  creationDate: number
}

const PostHeader: React.FC<PostHeaderProps> = ({
  creatorId,
  creatorDisplayName,
  creationDate
}) => {
  return (
    <>
      <div className={styles.postheadercontainer}>
        <Link to={"/dashboard/user/" + creatorDisplayName} >{creatorDisplayName}</Link>
      </div>
    </>
  )
}

export default PostHeader