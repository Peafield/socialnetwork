import React, { ChangeEvent, useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { handleAPIRequest } from "../../controllers/Api";
import Container from "../Containers/Container";
import styles from "./Auth.module.css";
import { useSetUserContextAndCookie } from "../../controllers/SetUserContextAndCookie";
import Snackbar from "../feedback/Snackbar";

interface SignInFormData {
  username_email: string;
  password: string;
}

export default function SignIn() {
  const navigate = useNavigate();
  const [formData, setFormData] = useState<SignInFormData>({
    username_email: "",
    password: "",
  });
  const [error, setError] = useState<string | null>(null);
  const [snackbarOpen, setSnackbarOpen] = useState<boolean>(false);
  const [snackbarType, setSnackbarType] = useState<
    "success" | "error" | "warning"
  >("error");

  const setUserContextAndCookie = useSetUserContextAndCookie();

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
      const response = await handleAPIRequest("/signin", options);
      if (response && response.status === "success") {
        const errorResult = setUserContextAndCookie(response.data);
        if (errorResult) {
          if (typeof errorResult === "string") {
            setError(errorResult);
            setSnackbarType("error");
            setSnackbarOpen(true);
          } else if (errorResult instanceof Error) {
            setError(errorResult.message);
            setSnackbarType("error");
            setSnackbarOpen(true);
          } else {
            setError("An unexpected error occurred.");
          }
        } else {
          setSnackbarType("success");
          setSnackbarOpen(true);
          setTimeout(() => {
            navigate("/dashboard");
          }, 1000);
        }
      }
    } catch (error) {
      if (error instanceof Error) {
        setError(
          "Oops something went wrong. Please wait a minute before trying again."
        );
        setSnackbarType("error");
        setSnackbarOpen(true);
      } else {
        setError("An unexpected error occurred.");
        setSnackbarType("error");
        setSnackbarOpen(true);
      }
    }
  };

  return (
    <Container>
      <div className={styles.authcontainer}>
        <div className={styles.formwrapper}>
          <h2 className={styles.h2}>Sign In</h2>
          <form onSubmit={handleSubmit}>
            <div className={styles.inputgroup}>
              <label className={styles.label} htmlFor="username_email">
                Username/Email:
                <input
                  className={styles.input}
                  type="text"
                  value={formData.username_email}
                  name="username_email"
                  onChange={handleChange}
                />
              </label>
            </div>
            <div className={styles.inputgroup}>
              <label className={styles.label} htmlFor="password">
                Password:
                <input
                  className={styles.input}
                  type="password"
                  value={formData.password}
                  name="password"
                  onChange={handleChange}
                />
              </label>
            </div>
            <div className={styles.inputgroup}>
              <button className={styles.button} type="submit">
                Sign In
              </button>
            </div>
          </form>
          <Link to="/signup">Don't have an account? Sign up</Link>
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
}
