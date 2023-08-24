import React from 'react'
import CommentActions from './CommentActions'
import CommentContent from './CommentContent'
import CommentHeader from './CommentHeader'
import styles from './Comment.module.css'

export interface CommentProps {
    comment_id: string,
    user_id: string,
    content: string,
    image: string,
    likes: number,
    dislikes: number,
    creation_date: string,
}

const Comment: React.FC<CommentProps> = ({
    comment_id,
    user_id,
    content,
    image,
    likes,
    dislikes,
    creation_date
}) => {
    

    return (
        <div
        className={styles.commentcontainer}>
            <CommentHeader user_id={user_id} creation_date={creation_date}/>
            <CommentContent content={content} image={image}/>
            <CommentActions likes={likes} dislikes={dislikes}/>
        </div>
    )
}

export default Comment