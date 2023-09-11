import React from "react";
import { useState, useEffect, useRef } from "react";
import styles from "./Post.module.css";
import { AiOutlineLike, AiOutlineDislike } from "react-icons/ai";
import { GoComment } from "react-icons/go";
import { HandleReaction } from "../../helpers/HandleReaction";
import { useWebSocketContext } from "../../context/WebSocketContext";
import { WebSocketReadMessage } from "../../Socket";

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
  const { message, sendMessage } = useWebSocketContext();
  let messageToSend: WebSocketReadMessage = {
    type: "",
    info: ""
  }
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

      messageToSend = {
        type: "notification",
        info: {
          receiver: creatorId,
          post_id: postId,
          reaction_type: reactionType
        }
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

        messageToSend = {
          type: "notification",
          info: {
            receiver: creatorId,
            post_id: postId,
            reaction_type: reactionType
          }
        }
      }
    }

    currentTimeout.current = setTimeout(async () => {
      await HandleReaction(creatorId, "post", postId, reactionType);

      sendMessage(messageToSend)

      messageToSend = {
        type: "",
        info: ""
      }
    }, 2000);
  };

  return (
    <div className={styles.postactionscontainer}>
      <div className={styles.postactionsubcontainer}>
        <div className={styles.postactioninfo}>
          <span className={styles.postactionspan}>
            <div className={styles.postaction}>
              <div style={{ paddingBottom: '20%' }}>
                {numOfLikes}
              </div>
              <div style={{ color: '#fa4d6a' }} onClick={(e) => handleLikeDislike("like", e)}><AiOutlineLike /></div>
            </div>
            <div className={styles.postaction}>
              <div style={{ paddingBottom: '20%' }}>
                {numOfDislikes}
              </div>
              <div style={{ color: '#fa4d6a' }} onClick={(e) => handleLikeDislike("dislike", e)}><AiOutlineDislike /></div>
            </div>
            <div className={styles.postaction}>{numOfComments} <GoComment style={{ color: '#fa4d6a' }} /></div>
          </span>
        </div>
      </div>
    </div>
  );
};

export default PostActions;
