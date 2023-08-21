import React, { useEffect, useState } from 'react'
import { handleAPIRequest } from '../../controllers/Api';
import { getCookie } from '../../controllers/SetUserContextAndCookie';
import Comment, { CommentProps } from '../Comment/Comment';
import styles from './Post.module.css'

interface PostCommentsProps {
  post_id: string
}

const PostComments: React.FC<PostCommentsProps> = ({
  post_id
}) => {
  const [postComments, setPostComments] = useState<
    CommentProps[] | null
  >(null);
  const [postCommentsLoading, setPostCommentsLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);



  useEffect(() => {
    const fetchData = async () => {
      setPostCommentsLoading(true);

      let url

      { post_id ? url = `/comment?post_id=${encodeURIComponent(post_id)}` : url = "/comments" }

      const options = {
        method: "GET",
        headers: {
          Authorization: "Bearer " + getCookie("sessionToken"),
          "Content-Type": "application/json",
        },
      };
      try {
        const response = await handleAPIRequest(url, options);
        console.log(response.data);

        const newpostcomments = response.data.Comments

        setPostComments(newpostcomments);

      } catch (error) {
        if (error instanceof Error) {
          setError(error.message);
        } else {
          setError("An unexpected error occurred.");
        }
      }
      setPostCommentsLoading(false);
    };

    fetchData(); // Call the async function
  }, [post_id]);

  if (postCommentsLoading) { return <p>Loading...</p> }

  return (
    <>
      {postComments
        ? postComments.map((commentprops) => (
          <div
            key={commentprops.comment_id}>
            <Comment
              key={commentprops.comment_id}
              comment_id={commentprops.comment_id}
              user_id={commentprops.user_id}
              content={commentprops.content}
              image={commentprops.image}
              likes={commentprops.likes}
              dislikes={commentprops.dislikes}
              creation_date={commentprops.creation_date} />
          </div>
        )) : null}
    </>
  )
}

export default PostComments