import React from 'react'

interface PostTimestampProps {
    time: number
}

const PostTimestamp: React.FC<PostTimestampProps> = ({
    time
}) => {
    return (
        <>
            <p>{time}</p>
        </>
    )
}

export default PostTimestamp