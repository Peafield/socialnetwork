import React, { useEffect, useState } from 'react'
import { IoPeopleCircle } from 'react-icons/io5';
import { Link } from 'react-router-dom';
import { getUserByUserID } from '../../controllers/GetUser';
import { createImageURL } from '../../controllers/ImageURL';
import FormatDate from '../../helpers/DateConversion';
import { ProfileProps } from '../Profile/Profile';
import styles from './Comment.module.css'

interface CommentHeaderProps {
    user_id: string,
    creation_date: string,
}

const CommentHeader: React.FC<CommentHeaderProps> = ({
    user_id,
    creation_date
}) => {
    const [userData, setUserData] = useState<ProfileProps | null>(null);
    const [profilePicUrl, setProfilePicUrl] = useState<string | null>(null);
    const [userLoading, setUserLoading] = useState<boolean>(false);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchData = async () => {
            setUserLoading(true);

            try {
                const newUserData = await getUserByUserID(user_id)
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
    }, [user_id]);

    useEffect(() => {
        if (userData?.avatar) {
          const url = createImageURL(userData.avatar)
          setProfilePicUrl(url);
    
          // Clean up the Blob URL when the component unmounts
          return () => {
            URL.revokeObjectURL(url);
          };
        }
      }, [userData?.avatar]);
      const formattedDate = FormatDate(creation_date);

    return (
        <div className={styles.commmentheadercontainer}>
            <div className={styles.commentHeaderInfoContainer}>
                <div>
                    {(profilePicUrl && (
                        <img
                            src={profilePicUrl}
                            alt="Profile pic"
                            className={styles.avatar}
                        />
                    )) || (
                            <span className={styles.profileIconStyle}>
                                <IoPeopleCircle />
                            </span>
                        )}
                </div>
                <div>
                    <Link to={"/dashboard/user/" + userData?.display_name}>
                        {userData?.display_name}
                    </Link>
                    <p>{formattedDate}</p>
                </div>
            </div>
        </div>
    )
}

export default CommentHeader