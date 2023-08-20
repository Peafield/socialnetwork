import React from "react";
import { useState } from "react";
import styles from "./Post.module.css";
import { AiOutlineLike, AiOutlineDislike } from "react-icons/ai";
import { GoComment } from "react-icons/go";
import { HandleReaction } from "../../helpers/HandleReaction";

interface PostActionsProps {
  postId: string;
  likes: number;
  dislikes: number;
  numOfComments: number;
}

const PostActions: React.FC<PostActionsProps> = ({
  postId,
  likes,
  dislikes,
  numOfComments,
}) => {
  const [hasLiked, setHasLiked] = useState(false);
  const [hasDisliked, setHasDisliked] = useState(false);

  const handleLike = async () => {
    if (hasLiked) {
      await HandleReaction("post", postId, "like", "remove");
      setHasLiked(false);
      likes--
    } else {
      if (hasDisliked) {
        await HandleReaction("post", postId, "dislike", "remove");
        setHasDisliked(false);
        dislikes--
      }
      await HandleReaction("post", postId, "like", "add");
      setHasLiked(true);
      likes++
    }
  };

  const handleDislike = async () => {
    if (hasDisliked) {
      await HandleReaction("post", postId, "dislike", "remove");
      setHasDisliked(false);
      dislikes--
    } else {
      if (hasLiked) {
        await HandleReaction("post", postId, "like", "remove");
        setHasLiked(false);
        likes--
      }
      await HandleReaction("post", postId, "dislike", "add");
      setHasDisliked(true);
      dislikes++
    }
  };

  return (
    <div className={styles.postactionscontainer}>
      <div className={styles.postactionsubcontainer}>
        <div className={styles.postactioninfo}>
          <span className={styles.postactionspan}>
            <p>
              {likes} <AiOutlineLike />
            </p>
            <p>
              {dislikes} <AiOutlineDislike />
            </p>
          </span>
          <p>{numOfComments} comments</p>
        </div>
      </div>
      <form>
        <div
          className={`${styles.postactionsubcontainer} ${styles.postactionsubcontainerbottom}`}
        >
          <button onClick={handleLike}>
            <AiOutlineLike /> Like
          </button>
          <button onClick={handleDislike}>
            <AiOutlineDislike /> Dislike
          </button>
          <button>
            <GoComment /> Comment
          </button>
        </div>
      </form>
    </div>
  );
};

export default PostActions;
