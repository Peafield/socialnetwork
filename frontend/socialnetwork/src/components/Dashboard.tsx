import React, { useEffect, useState } from 'react'
import { handleAPIRequest } from '../controllers/Api';
import { getCookie } from '../controllers/SetUserContextAndCookie';
import Container from "./Containers/Container";
import Post, { PostProps } from './Post/Post';



export default function Dashboard() {
    const [userViewablePosts, setUserViewablePosts] = useState<PostProps[] | null>(null)
    const [postsLoading, setPostsLoading] = useState<boolean>(false)

    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchData = async () => {
            setPostsLoading(true);
            console.log(getCookie("sessionToken"));
            
            const options = {
                method: "GET",
                headers: {
                    'Authorization': 'Bearer ' + getCookie("sessionToken"),
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

        <Container>
            <div>
            {userViewablePosts
                    ? userViewablePosts.map((postProps) => (
                          <Post 
                          key={postProps.post_id}
                          post_id={postProps.post_id}
                          group_id={postProps.group_id}
                          creator_id={postProps.creator_id}
                          title={postProps.title}
                          image_path={postProps.image_path}
                          content={postProps.content}
                          num_of_comments={postProps.num_of_comments}
                          privacy_level={postProps.privacy_level}
                          likes={postProps.likes}
                          dislikes={postProps.dislikes}
                          creation_date={postProps.creation_date}
                           />
                      ))
                    : null}
            </div>
        </Container>
    )
}
