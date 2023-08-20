import React from 'react';
import styles from './Post.module.css';

interface PostContentProps {
  text: string;
}

const PostContent: React.FC<PostContentProps> = ({ text }) => {
  return (
    <>
      {text && (
        <div className={styles.postcontentcontainer}>
          <p>{text}</p>
        </div>
      )}
    </>
  );
}

export default PostContent;


