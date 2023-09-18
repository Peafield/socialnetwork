import React, { useEffect, useState } from 'react'
import { createImageURL } from '../../controllers/ImageURL';
import styles from './Comment.module.css'

interface CommentContentProps {
  content: string,
  image: string
}

const CommentContent: React.FC<CommentContentProps> = ({
  content,
  image
}) => {
  const [commentURL, setCommentURL] = useState<string | null>(null)

  useEffect(() => {
    if (image) {
      const url = createImageURL(image)

      setCommentURL(url);

      // Clean up the Blob URL when the component unmounts
      return () => {
        URL.revokeObjectURL(url);
      };
    }
  }, [image]);

  return (
    <div className={styles.commentcontentcontainer}>
      <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'left', width: '100%' }}>
        <p>{content}</p>
      </div>
      {commentURL ?
        <div
          className={styles.commentpic}>
          <img
            style={{ maxWidth: '180px', maxHeight: '120px' }}
            src={commentURL}
            alt="Post Pic"
          />
        </div>
        : null
      }
    </div>
  )
}

export default CommentContent