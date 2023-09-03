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
    <>
      <div className={styles.postcontentcontainer}>
        <p>{text}</p>
      </div>
      {postURL ?
        <div>
          <img
          style={{width: '300px', height: '300px'}}
            src={postURL}
            alt="Post Pic"
            className={styles.postpic}
          />
        </div>
        : null
      }
    </>
  )
}

export default PostContent