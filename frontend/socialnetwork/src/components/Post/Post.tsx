import React, { useEffect, useState } from "react";
import PostHeader from "./PostHeader";
import PostContent from "./PostContent";
import PostActions from "./PostActions";
import { ProfileProps } from "../Profile/Profile";
import { getUserByUserID } from "../../controllers/GetUser";

export interface PostProps {
  post_id: string;
  group_id: string;
  creator_id: string;
  image_path: string;
  content: string;
  num_of_comments: number;
  privacy_level: number;
  likes: number;
  dislikes: number;
  creation_date: string;
}

const Post: React.FC<PostProps> = ({
  post_id,
  group_id,
  creator_id,
  image_path,
  content,
  num_of_comments,
  privacy_level,
  likes,
  dislikes,
  creation_date,
}) => {
  const [userData, setUserData] = useState<ProfileProps | null>(null);
  const [userLoading, setUserLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchData = async () => {
      setUserLoading(true);

      try {
        const newUserData = await getUserByUserID(creator_id)
        setUserData(newUserData);
      } catch (error) {
        if (error instanceof Error) {
          setError(error.message);
        } else {
          setError("An unexpected error occurred.");
        }
      }
      setUserLoading(false);
    };

    fetchData(); // Call the async function
  }, [creator_id]);
  
  return (
    <>
      {userData ? (
        <>
          <PostHeader
            creatorDisplayName={userData.display_name}
            creatorId={creator_id}
            creationDate={creation_date}
            creatorAvatar={userData.avatar}
            postPrivacyLevel={privacy_level}
          />
          <PostContent text={content} />
          <PostActions
            likes={likes}
            dislikes={dislikes}
            numOfComments={num_of_comments}
          />
        </>
      ) : null}
    </>
  );
};

export default Post;
