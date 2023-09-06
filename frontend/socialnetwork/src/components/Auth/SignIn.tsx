import React, { ChangeEvent, useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { handleAPIRequest } from "../../controllers/Api";
import Container from "../Containers/Container";
import styles from "./Signin.module.css";
import { useSetUserContextAndCookie } from "../../controllers/SetUserContextAndCookie";
import Snackbar from "../feedback/Snackbar";
import { IoShareSocial } from "react-icons/io5";

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
        setError("Invalid username/email or password");
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
      <div className={styles.signinpage}>
        <div className={`${styles.logo} ${styles.flip}`}>
          <IoShareSocial />
        </div>
        <h2
          style={{
            position: "absolute",
            top: "10%",
            left: "66%",
            color: "#fa4d6a",
          }}
        >
          Social Network
        </h2>
        <div className={styles.authcontainer}>
          <div className={styles.formwrapper}>
            <form onSubmit={handleSubmit}>
              <div className={styles.inputgroup}>
                <input
                  className={styles.input}
                  type="text"
                  placeholder="Username/Email"
                  value={formData.username_email}
                  name="username_email"
                  onChange={handleChange}
                />
              </div>
              <div className={styles.inputgroup}>
                <input
                  className={styles.input}
                  type="password"
                  placeholder="Password"
                  value={formData.password}
                  name="password"
                  onChange={handleChange}
                />
              </div>
              <div className={styles.inputgroup}>
                <button
                  className={styles.button}
                  type="submit"
                  onClick={() => {
                    setSnackbarOpen(false);
                    setError(null);
                  }}
                >
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
      </div>
    </Container>
  );
}
