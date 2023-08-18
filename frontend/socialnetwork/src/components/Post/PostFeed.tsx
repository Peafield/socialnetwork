import React, { useEffect, useState } from 'react'
import { handleAPIRequest } from '../../controllers/Api';
import { getCookie } from '../../controllers/SetUserContextAndCookie';
import Container from '../Containers/Container'
import Modal from '../Containers/Modal';
import Post, { PostProps } from './Post'
import styles from './Post.module.css'
import PostComments from './PostComments';

const PostFeed: React.FC = () => {
    const [userViewablePosts, setUserViewablePosts] = useState<
        PostProps[] | null
    >(null);
    const [postsLoading, setPostsLoading] = useState<boolean>(false);
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [modalPost, setModalPost] = useState<PostProps | null>(null)
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchData = async () => {
            setPostsLoading(true);

            const options = {
                method: "GET",
                headers: {
                    Authorization: "Bearer " + getCookie("sessionToken"),
                    "Content-Type": "application/json",
                },
            };
            try {
                const response = await handleAPIRequest("/post", options);
                setUserViewablePosts(response.data.Posts);
                console.log(response.data.Posts);

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
    }, []);

    if (postsLoading) return <p>Loading...</p>

    return (
        <Container>
            <div className={styles.postfeedcontainer}>
                <Modal open={isModalOpen} onClose={() => setIsModalOpen(false)}>
                    {modalPost ?
                        <>
                            <Post
                                key={modalPost.post_id}
                                post_id={modalPost.post_id}
                                group_id={modalPost.group_id}
                                creator_id={modalPost.creator_id}
                                creator_display_name={modalPost.creator_display_name}
                                image_path={modalPost.image_path}
                                content={modalPost.content}
                                num_of_comments={modalPost.num_of_comments}
                                privacy_level={modalPost.privacy_level}
                                likes={modalPost.likes}
                                dislikes={modalPost.dislikes}
                                creation_date={modalPost.creation_date}
                            />
                            <PostComments post_id={modalPost.post_id}/>
                        </> : null
                    }

                    <button onClick={() => setIsModalOpen(false)}>Close</button>
                </Modal>
                {userViewablePosts
                    ? userViewablePosts.map((postProps) => (
                        <div
                            className={styles.postcontainer}
                            key={postProps.post_id}
                            onClick={() => {
                                setIsModalOpen(true)
                                setModalPost(postProps)
                            }}>
                            <Post
                                key={postProps.post_id}
                                post_id={postProps.post_id}
                                group_id={postProps.group_id}
                                creator_id={postProps.creator_id}
                                creator_display_name={postProps.creator_display_name}
                                image_path={postProps.image_path}
                                content={postProps.content}
                                num_of_comments={postProps.num_of_comments}
                                privacy_level={postProps.privacy_level}
                                likes={postProps.likes}
                                dislikes={postProps.dislikes}
                                creation_date={postProps.creation_date}
                            />
                        </div>
                    ))
                    : null}
            </div>
        </Container>
    )
}

export default PostFeed