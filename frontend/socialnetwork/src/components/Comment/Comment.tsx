import React, { useContext, useEffect, useState } from 'react'
import CommentActions from './CommentActions'
import CommentContent from './CommentContent'
import CommentHeader from './CommentHeader'
import styles from './Comment.module.css'
import { UserContext } from '../../context/AuthContext'
import { GetUserCommentReaction } from '../../controllers/GetUserReaction'

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
    const userContext = useContext(UserContext)
    const [userReaction, setUserReaction] = useState<string | null>(null)
    const [error, setError] = useState<string | null>(null)

    useEffect(() => {
        const fetchData = async () => {
            try {
                if (userContext.user) {
                    const reactionData = await GetUserCommentReaction(comment_id)
                    if (reactionData) {
                        setUserReaction(reactionData.reaction)
                    }
                }
            } catch (error) {
                if (error instanceof Error) {
                    setError(error.message);
                } else {
                    setError("An unexpected error occurred.");
                }
            }
        };

        fetchData(); // Call the async function
    }, []);

    return (
        <div
            className={styles.commentcontainer}>
            <CommentHeader user_id={user_id} creation_date={creation_date} />
            <CommentContent content={content} image={image} />
            <CommentActions creatorId={user_id} commentId={comment_id} likes={likes} dislikes={dislikes} userReaction={userReaction} />
        </div>
    )
}

export default Comment