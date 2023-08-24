import React, { useEffect, useState } from 'react'
import { FaComment, FaCommentMedical } from 'react-icons/fa';
import { handleAPIRequest } from '../../controllers/Api';
import { getCookie } from '../../controllers/SetUserContextAndCookie';
import Comment, { CommentProps } from '../Comment/Comment';
import styles from './Post.module.css'

interface PostCommentsProps {
  post_id: string
}

interface CommentFormData {
  post_id: string,
  content: string,
  image: string
}

const PostComments: React.FC<PostCommentsProps> = ({
  post_id
}) => {
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
      <div
      className={styles.postcommentscontainer}>
        <div>Comments <FaComment /></div>
        <div
        className={styles.commentformcontainer}>
          <form onSubmit={handleSubmit}>
            <div>
              <textarea
                required
                maxLength={100}
                placeholder="Write a comment..."
                id="content"
                name="content"
                value={commentFormData.content}
                onChange={handleChange}
              />
            </div>
            <div>
              <button type="submit">
                <FaCommentMedical />
              </button>
            </div>
          </form>
        </div>
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
      </div>
    </>
  )
}

export default PostComments