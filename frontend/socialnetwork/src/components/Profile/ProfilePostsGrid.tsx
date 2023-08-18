import React, { useEffect, useState } from 'react'
import { handleAPIRequest } from '../../controllers/Api';
import { getCookie } from '../../controllers/SetUserContextAndCookie';
import Container from '../Containers/Container';
import Modal from '../Containers/Modal';
import Post, { PostProps } from '../Post/Post'
import PostComments from '../Post/PostComments';
import styles from './Profile.module.css'

interface ProfilePostsGridProps {
    user_id: string
}

const ProfilePostsGrid: React.FC<ProfilePostsGridProps> = ({
    user_id
}) => {
    const [profilePosts, setProfilePosts] = useState<PostProps[] | null>(null)
    const [postsLoading, setPostsLoading] = useState<boolean>(false);
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [modalPost, setModalPost] = useState<PostProps | null>(null)
    const [error, setError] = useState<string | null>(null);


    useEffect(() => {
        const fetchData = async () => {
            setPostsLoading(true);

            let url

            {user_id ? url = `/post?user_id=${encodeURIComponent(user_id)}` : url = "/post"}

            const options = {
                method: "GET",
                headers: {
                    Authorization: "Bearer " + getCookie("sessionToken"),
                    "Content-Type": "application/json",
                },
            };
            try {
                const response = await handleAPIRequest(url, options);
                setProfilePosts(response.data.Posts);
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

  return (
    <Container>
            <div className={styles.profilepostsgridcontainer}>
                <Modal open={isModalOpen} onClose={() => setIsModalOpen(false)}>
                    {modalPost ?
                        <>
                            <Post
                                key={modalPost.post_id}
                                post_id={modalPost.post_id}
                                group_id={modalPost.group_id}
                                creator_id={modalPost.creator_id}
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
                {profilePosts
                    ? profilePosts.map((postProps) => (
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

export default ProfilePostsGrid