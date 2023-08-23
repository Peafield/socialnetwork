import React from 'react'
import { useEffect, useState } from "react"
import styles from "./Post.module.css"


interface PostImageProps {
    imageString: string
}

const PostImage: React.FC<PostImageProps> = ({imageString}) => {
  const [image, setIamge] = useState<string | null>(null);

  useEffect(() => {
    if (imageString) {
      const decodedAvatar = atob(imageString); // Decode base64-encoded avatar data
      const avatarBuffer = new ArrayBuffer(decodedAvatar.length);
      const avatarView = new Uint8Array(avatarBuffer);
      for (let i = 0; i < decodedAvatar.length; i++) {
        avatarView[i] = decodedAvatar.charCodeAt(i);
      }

      const blob = new Blob([avatarBuffer]);
      const url = URL.createObjectURL(blob);


      setIamge(url);

      // Clean up the Blob URL when the component unmounts
      return () => {
        URL.revokeObjectURL(url);
      };
    }
  }, [imageString]);

  return (
    <>
      <div className={styles.postimagecontainer}>
          {(image && (
            <img
              src={image}
              alt="User Image"
              className={styles.avatar}
            />
          ))}
        </div>
    </>
  )
}

export default PostImage