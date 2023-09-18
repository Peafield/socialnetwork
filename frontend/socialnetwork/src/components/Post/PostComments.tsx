import React, { useEffect, useState } from 'react'
import { FaComment, FaCommentMedical } from 'react-icons/fa';
import { useWebSocketContext } from '../../context/WebSocketContext';
import { handleAPIRequest } from '../../controllers/Api';
import { getCookie } from '../../controllers/SetUserContextAndCookie';
import { WebSocketReadMessage } from '../../Socket';
import Comment, { CommentProps } from '../Comment/Comment';
import styles from './Post.module.css'

interface PostCommentsProps {
  post_id: string
  creator_id: string
}

interface CommentFormData {
  post_id: string,
  content: string,
  image: string
}

const PostComments: React.FC<PostCommentsProps> = ({
  post_id,
  creator_id
}) => {
  const { message, sendMessage } = useWebSocketContext();
  const [postComments, setPostComments] = useState<
    CommentProps[] | null
  >(null);
  const [commentFormData, setCommentFormData] = useState<CommentFormData>({
    post_id: post_id,
    content: "",
    image: ""
  })
  const [postCommentsLoading, setPostCommentsLoading] = useState<boolean>(false);
  const [updateTrigger, setUpdateTrigger] = useState<number>(0)
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

        const commentData = response.data.Comments.map((comment: any) => {
          const newcomment = comment.CommentInfo
          newcomment.image = comment.CommentPicture
          return newcomment
        })

        setPostComments(commentData);

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
  }, [post_id, updateTrigger]);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target;

    if (e.target.type === "file") {
      const file = (e.target as HTMLInputElement)?.files?.[0] || null;
      if (file) {
        const reader = new FileReader();
        reader.onloadend = () => {
          setCommentFormData((prevState) => ({
            ...prevState,
            image: reader.result as string,
          }));
        }
        reader.readAsDataURL(file)
      }

    } else {
      setCommentFormData((prevState) => ({
        ...prevState,
        [name]: value,
      }));
    }
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const data = { data: commentFormData };

    const options = {
      method: "POST",
      headers: {
        Authorization: "Bearer " + getCookie("sessionToken"),
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    };
    try {
      const response = await handleAPIRequest("/comment", options);
      if (response && response.status === "success") {
        console.log("comment submit success");
        setUpdateTrigger(prevTigger => prevTigger + 1)
        const messageToSend: WebSocketReadMessage = {
          type: "notification",
          info: {
            receiver: creator_id,
            post_id: commentFormData.post_id,
            action_type: "comment"
          }
        }
        sendMessage(messageToSend)
      }
    } catch (error) {
      if (error instanceof Error) {
        setError(error.message);
      } else {
        setError("An unexpected error occurred.");
      }
    }
  };

  if (postCommentsLoading) { return <p>Loading...</p> }

  return (
    <>
      <div className={styles.addcommentcontainer}>
        <div>Comments <FaComment /></div>
        <form onSubmit={handleSubmit}>
          <div
            className={styles.commentformcontainer}>
            <div style={{ width: '80%' }}>
              <textarea
                required
                maxLength={100}
                placeholder="Write a comment..."
                id="content"
                name="content"
                value={commentFormData.content}
                onChange={handleChange}
              />
              <input
                type="file"
                id="image_path"
                name="image_path"
                onChange={handleChange}
              />
            </div>
            <div style={{ width: '10%' }}>
              <button type="submit" style={{ width: 'auto' }}>
                <FaCommentMedical />
              </button>
            </div>
          </div>
        </form>
      </div>
      {postComments
        ? postComments.map((commentprops) => (
          <Comment
            key={commentprops.comment_id}
            comment_id={commentprops.comment_id}
            user_id={commentprops.user_id}
            content={commentprops.content}
            image={commentprops.image}
            likes={commentprops.likes}
            dislikes={commentprops.dislikes}
            creation_date={commentprops.creation_date} />
        )) : null}
    </>
  )
}

export default PostComments