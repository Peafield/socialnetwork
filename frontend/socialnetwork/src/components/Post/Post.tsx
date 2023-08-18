import React from 'react'
import PostHeader from './PostHeader'
import PostContent from './PostContent'
import PostActions from './PostActions'


export interface PostProps {
    post_id: string,
    group_id: string,
    creator_id: string,
    creator_display_name: string,
    image_path: string,
    content: string,
    num_of_comments: number,
    privacy_level: number,
    likes: number,
    dislikes: number,
    creation_date: number
}

const Post: React.FC<PostProps> = ({
    post_id,
    group_id,
    creator_id,
    creator_display_name,
    image_path,
    content,
    num_of_comments,
    privacy_level,
    likes,
    dislikes,
    creation_date
}) => {
    return (
        <>
            <PostHeader creatorDisplayName={creator_display_name} creatorId={creator_id} creationDate={creation_date}/>
            <PostContent text={content}/>
            <PostActions likes={likes} dislikes={dislikes} numOfComments={num_of_comments} />
        </>

    )
}

export default Post