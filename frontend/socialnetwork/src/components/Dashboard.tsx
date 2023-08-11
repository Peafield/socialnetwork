import React, { useEffect, useState } from 'react'
import { handleAPIRequest } from '../controllers/Api';
import Container from "./Containers/Container";

interface Post {
    post_id: string,
    group_id: string,
    creator_id: string,
    title: string,
    image_path: string,
    content: string,
    num_of_comments: number,
    privacy_level: number,
    likes: number,
    dislikes: number,
    creation_date: number
}

export default function Dashboard() {
    const [userViewablePosts, setUserViewablePosts] = useState<Post[] | null>(null)
    const [postsLoading, setPostsLoading] = useState<boolean>(false)

    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchData = async () => {
            setPostsLoading(true);
            const options = {
                method: "GET",
                headers: {
                    "Content-Type": "application/json",
                },
            };
            try {
                const response = await handleAPIRequest("/post", options);
                setUserViewablePosts(response.Data)
            } catch (error) {
                if (error instanceof Error) {
                    setError(error.message);
                } else {
                    setError("An unexpected error occurred.");
                }
            }
            setPostsLoading(false);
        };

        fetchData(); // Call the async function

    }, [userViewablePosts]);

    return (
        <>
            <Container>
                <div>
                    <h2>Posts</h2>
                </div>
            </Container>
        </>
    )
}
