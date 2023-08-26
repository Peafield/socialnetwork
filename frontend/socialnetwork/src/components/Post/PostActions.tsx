import React from "react";
import { useState, useEffect, useRef } from "react";
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

  const currentTimeout = useRef<ReturnType<typeof setTimeout> | null>(null);

  // handleLikeDislike delays the sending of reaction for 5 seconds to make sure user decision is final.
  const handleLikeDislike = (reactionType: "like" | "dislike", e: React.MouseEvent) => {
    e.preventDefault();
    e.stopPropagation();

    if (currentTimeout.current) {
      clearTimeout(currentTimeout.current);
    }


    if (reactionType === "like") {
      if (hasLiked) {
        setNumOfLikes((prev) => prev - 1);
        setHasLiked(false);
      } else {
        if (hasDisliked) {
          setNumOfDisikes((prev) => prev - 1);
          setHasDisliked(false);
        }
        setNumOfLikes((prev) => prev + 1);
        setHasLiked(true);
      }
    }

    if (reactionType === "dislike") {
      if (hasDisliked) {
        setNumOfDisikes((prev) => prev - 1);
        setHasDisliked(false);
      } else {
        if (hasLiked) {
          setNumOfLikes((prev) => prev - 1);
          setHasLiked(false);
        }
        setNumOfDisikes((prev) => prev + 1);
        setHasDisliked(true);
      }
    }

    currentTimeout.current = setTimeout(async () => {
      await HandleReaction(creatorId, "post", postId, reactionType);
    }, 5000);
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
          <button onClick={(e) => handleLikeDislike("like", e)}>
            <AiOutlineLike /> Like
          </button>
          <button onClick={(e) => handleLikeDislike("dislike", e)}>
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
