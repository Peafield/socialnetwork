import React, { useEffect, useState } from 'react'
import { createImageURL } from '../../controllers/ImageURL'
import styles from './Post.module.css'

interface PostContentProps {
  text: string
  image_path: string
}

const PostContent: React.FC<PostContentProps> = ({
  text,
  image_path
}) => {
  const [postURL, setPostURL] = useState<string | null>(null)

  useEffect(() => {
    if (image_path) {
      const url = createImageURL(image_path)

      setPostURL(url);

      // Clean up the Blob URL when the component unmounts
      return () => {
        URL.revokeObjectURL(url);
      };
    }
  }, [image_path]);

  return (
    <div className={styles.postcontentcontainer}>
      <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'left', width: '100%' }}>
        <p>{text}</p>
      </div>
      {postURL ?
        <div
          className={styles.postpic}>
          <img
            style={{ maxWidth: '300px', maxHeight: '200px' }}
            src={postURL}
            alt="Post Pic"

          />
        </div>
        : null
      }
    </div>
  )
}

export default PostContent;


