import React, { CSSProperties, useState } from 'react'
import { AiOutlineClose } from 'react-icons/ai';
import Container from '../Containers/Container';
import Modal from '../Containers/Modal';
import Post, { PostProps } from '../Post/Post';
import PostComments from '../Post/PostComments';
import styles from './Group.module.css'

interface GroupPostsProps {
    posts: PostProps[]
    isUserMember: boolean
}

const GroupPosts: React.FC<GroupPostsProps> = ({
    posts,
    isUserMember
}) => {
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [modalPost, setModalPost] = useState<PostProps | null>(null)

    const closeStyle: CSSProperties = {
        margin: "10px",
        verticalAlign: "middle",
        color: "red",
        borderRadius: "50%",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        height: "3%",
        width: "3%",
    };

    return (
        <div className={styles.grouppostscontainer}>
            <Modal open={isModalOpen} onClose={() => setIsModalOpen(false)}>
                <span style={closeStyle} onClick={() => setIsModalOpen(false)}>
                    <AiOutlineClose />
                </span>
                {modalPost ?
                    <div
                        className={styles.postmodalcontainer}>
                        <div
                            className={styles.postcontainer}>
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
                        </div>
                        <div
                            className={styles.postcommentscontainer}>
                            <PostComments post_id={modalPost.post_id} creator_id={modalPost.creator_id} />
                        </div>
                    </div>
                    : null
                }
            </Modal>
            {posts
                && posts.map((postProps) => (
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
                ))}
        </div>
    )
}

export default GroupPosts