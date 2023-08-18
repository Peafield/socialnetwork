import React, { ChangeEvent, useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import { handleAPIRequest } from "../../controllers/Api";
import Container from "../Containers/Container";
import Snackbar from "../feedback/Snackbar";
import styles from "./Post.module.css";

interface CreatePostFormData {
  content: string;
  privacy_level: number;
}

const CreatePost: React.FC = () => {
  const navigate = useNavigate();
  const [formData, setFormData] = useState<CreatePostFormData>({
    content: "",
    privacy_level: 0,
  });
  const [error, setError] = useState<string | null>(null);
  const [snackbarOpen, setSnackbarOpen] = useState<boolean>(false);
  const [snackbarType, setSnackbarType] = useState<
    "success" | "error" | "warning"
  >("error");

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });
  };

  const handleSubmit = async (e: { preventDefault: () => void }) => {
    e.preventDefault();
    const data = { data: formData };
    const options = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    };
    try {
      const response = await handleAPIRequest("/post", options);
      if (response && response.status === "success") {
        setSnackbarType("success");
        setSnackbarOpen(true);
        setTimeout(() => {
          navigate("/dashboard");
        }, 1000);
      }
    } catch (error) {
      if (error instanceof Error) {
        setError("Could not create post");
        setSnackbarType("error");
        setSnackbarOpen(true);
      } else {
        setError("An unexpected error occurred");
        setSnackbarType("error");
        setSnackbarOpen(true);
      }
    }
  };

  return (
    <Container>
      <div className={styles.createpostcontainer}>
        <div className={styles.formwrapper}>
          <h2 className={styles.h2}>Create Post</h2>
          <form onSubmit={handleSubmit}>
            <div className={styles.inputgroup}>
              <input
                className={styles.input}
                type="text"
                placeholder="Write something for your post!"
                value={formData.content}
                name="content"
                onChange={handleChange}
              />
            </div>
            <div className={styles.inputgrouprow}>
              <input
                className={styles.input}
                id="public_privacy_level"
                type="radio"
                value={0}
                name="privacy_level"
                onChange={handleChange}
              />
              <label htmlFor="public_privacy_level">Public</label>
              <input
                className={styles.input}
                id="private_privacy_level"
                type="radio"
                value={1}
                name="privacy_level"
                onChange={handleChange}
              />
              <label htmlFor="private_privacy_level">Private</label>
              <input
                className={styles.input}
                id="selected_privacy_level"
                type="radio"
                value={2}
                name="privacy_level"
                onChange={handleChange}
              />
              <label htmlFor="selected_privacy_level">Selected</label>
            </div>
            <div className={styles.inputgroup}>
              <button className={styles.button} type="submit">
                Create Post
              </button>
            </div>
          </form>
        </div>
      </div>
      <Snackbar
        open={snackbarOpen}
        onClose={() => {
          setSnackbarOpen(false);
          setError(null);
        }}
        message={error ? error : "Signed in successfully!"}
        type={snackbarType}
      />
    </Container>
  );
};

export default CreatePost;
