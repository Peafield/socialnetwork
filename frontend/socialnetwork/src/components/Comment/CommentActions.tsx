import React, { useEffect, useRef, useState } from 'react'
import { AiFillLike, AiOutlineLike, AiFillDislike, AiOutlineDislike } from 'react-icons/ai';
import { GoComment } from 'react-icons/go';
import { useWebSocketContext } from '../../context/WebSocketContext';
import { HandleReaction } from '../../helpers/HandleReaction';
import { WebSocketReadMessage } from '../../Socket';
import styles from './Comment.module.css'

interface CommentActionsProps {
  creatorId: string;
  commentId: string;
  likes: number,
  dislikes: number
  userReaction: string | null
}

const CommentActions: React.FC<CommentActionsProps> = ({
  creatorId,
  commentId,
  likes,
  dislikes,
  userReaction
}) => {
  const { message, sendMessage } = useWebSocketContext();
  let messageToSend: WebSocketReadMessage = {
    type: "",
    info: ""
  }
  const [numOfLikes, setNumOfLikes] = useState(likes);
  const [numOfDislikes, setNumOfDisikes] = useState(dislikes);

  const [hasLiked, setHasLiked] = useState(userReaction === "like");
  const [hasDisliked, setHasDisliked] = useState(userReaction === "dislike");

  useEffect(() => {
    setNumOfLikes(likes);
    setNumOfDisikes(dislikes);
    setHasLiked(userReaction === "like")
    setHasDisliked(userReaction === "dislike")
  }, [likes, dislikes, userReaction]);

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

        messageToSend = {
          type: "notification",
          info: {
            receiver: creatorId,
            comment_id: commentId,
            action_type: reactionType
          }
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
            comment_id: commentId,
            action_type: reactionType
          }
        }
      }
    }

    currentTimeout.current = setTimeout(async () => {
      await HandleReaction(creatorId, "comment", commentId, reactionType);

      sendMessage(messageToSend)

      messageToSend = {
        type: "",
        info: ""
      }
    }, 500);
  };

  return (
    <div className={styles.commentactionscontainer}>
      <div className={styles.commentactionsubcontainer}>
        <div className={styles.commentactioninfo}>
          <span className={styles.commentactionspan}>
            <div className={styles.commentaction}>
              <div style={{ paddingBottom: '20%' }}>
                {numOfLikes}
              </div>
              <div style={{ color: '#fa4d6a' }} onClick={(e) => handleLikeDislike("like", e)}>{hasLiked ? <AiFillLike /> : <AiOutlineLike />}</div>
            </div>
            <div className={styles.commentaction}>
              <div style={{ paddingBottom: '20%' }}>
                {numOfDislikes}
              </div>
              <div style={{ color: '#fa4d6a' }} onClick={(e) => handleLikeDislike("dislike", e)}>{hasDisliked ? <AiFillDislike /> : <AiOutlineDislike />}</div>
            </div>
            <div className={styles.commentaction}>{0} <GoComment style={{ color: '#fa4d6a' }} /></div>
          </span>
        </div>
      </div>
    </div>
  );
};

export default CommentActions 