import React from "react";
import { useState, useEffect } from "react";
import styles from "./Post.module.css";
import { AiOutlineLike, AiOutlineDislike } from "react-icons/ai";
import { GoComment } from "react-icons/go";
import { HandleReaction } from "../../helpers/HandleReaction";

interface PostActionsProps {
  creatorId: string;
  postId: string;
  likes: number;
  dislikes: number;
  AmountOfComments: number;
}

const PostActions: React.FC<PostActionsProps> = ({
  creatorId,
  postId,
  likes,
  dislikes,
  AmountOfComments,
}) => {
  const [numOfLikes, setNumOfLikes] = useState(likes);
  const [numOfDislikes, setNumOfDisikes] = useState(dislikes);
  const [numOfComments, setNumOfComments] = useState(AmountOfComments);

  const [hasLiked, setHasLiked] = useState(false);
  const [hasDisliked, setHasDisliked] = useState(false);

  useEffect(() => {
    setNumOfLikes(likes);
    setNumOfDisikes(dislikes);
    setNumOfComments(AmountOfComments);
  }, [likes, dislikes, AmountOfComments]);

  const handleLike = async (e: React.MouseEvent) => {
    e.preventDefault();
    e.stopPropagation();
    if (hasLiked) {
      await HandleReaction(creatorId, "post", postId, "like");
      setHasLiked(false);
      setNumOfLikes(likes + 1);
    } else {
      if (hasDisliked) {
        await HandleReaction(creatorId, "post", postId, "dislike");
        setHasDisliked(false);
        setNumOfDisikes(dislikes - 1);
      }
      await HandleReaction(creatorId, "post", postId, "like");
      setHasLiked(true);
      setNumOfLikes(likes + 1);
    }
  };

  const handleDislike = async (e: React.MouseEvent) => {
    e.preventDefault();
    e.stopPropagation();
    if (hasDisliked) {
      await HandleReaction(creatorId, "post", postId, "dislike");
      setHasDisliked(false);
      setNumOfDisikes(dislikes - 1);
    } else {
      if (hasLiked) {
        await HandleReaction(creatorId, "post", postId, "like");
        setHasLiked(false);
        setNumOfLikes(likes + 1);
      }
      await HandleReaction(creatorId, "post", postId, "dislike");
      setHasDisliked(true);
      setNumOfDisikes(dislikes - 1);
    }
  };

  return (
    <div className={styles.postactionscontainer}>
      <div className={styles.postactionsubcontainer}>
        <div className={styles.postactioninfo}>
          <span className={styles.postactionspan}>
            <p>
              {numOfLikes} <AiOutlineLike />
            </p>
            <p>
              {numOfDislikes} <AiOutlineDislike />
            </p>
          </span>
          <p>{numOfComments} comments</p>
        </div>
      </div>
      <form>
        <div
          className={`${styles.postactionsubcontainer} ${styles.postactionsubcontainerbottom}`}
        >
          <button onClick={(e) => handleLike(e)}>
            <AiOutlineLike /> Like
          </button>
          <button onClick={(e) => handleDislike(e)}>
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
