import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { handleAPIRequest } from "../../controllers/Api";
import styles from "./Signup.module.css";
import Container from "../Containers/Container";
import { useSetUserContextAndCookie } from "../../controllers/SetUserContextAndCookie";
import { IoShareSocial } from "react-icons/io5";
import Snackbar from "../feedback/Snackbar";

interface SignUpFormData {
  email: string;
  display_name: string;
  password: string;
  confirmPassword: string;
  first_name: string;
  last_name: string;
  dob: string;
  avatar_path: string;
  about_me: string;
}

export default function SignUp() {
  const setUserContextAndCookie = useSetUserContextAndCookie();
  const navigate = useNavigate();
  const [error, setError] = useState<string | null>(null);
  const [formData, setFormData] = useState<SignUpFormData>({
    email: "",
    display_name: "",
    password: "",
    confirmPassword: "",
    first_name: "",
    last_name: "",
    dob: "",
    avatar_path: "",
    about_me: "",
  });
  const [snackbarOpen, setSnackbarOpen] = useState<boolean>(false);
  const [snackbarType, setSnackbarType] = useState<
    "success" | "error" | "warning"
  >("error");

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target;

    if (e.target.type === "file") {
      const file = (e.target as HTMLInputElement)?.files?.[0] || null;
      if (file) {
        const reader = new FileReader();
        reader.onloadend = () => {
          setFormData((prevState) => ({
            ...prevState,
            avatar_path: reader.result as string,
          }));
        }
        reader.readAsDataURL(file)
      }

    } else {
      setFormData((prevState) => ({
        ...prevState,
        [name]: value,
      }));
    }
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const data = { data: formData };

    if (formData.password != formData.confirmPassword) {
      setError("passwords do not match")
      return
    }

    const options = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    };
    try {
      const response = await handleAPIRequest("/signup", options);
      if (response && response.status === "success") {
        setUserContextAndCookie(response.data);
        navigate("/dashboard");
      }
    } catch (error) {
      if (error instanceof Error) {
        if (error.cause && typeof error.cause === 'string') {
          const causeString: string = error.cause;
          if (causeString == '400') {
            setError("Email or username already taken");
          }
        } else {
          setError(error.message)
        }
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
      <div
        className={styles.signuppage}>
        <div className={`${styles.logo} ${styles.flip}`}>
          <IoShareSocial />
        </div>
        <div className={styles.suauthcontainer}>
          <div className={styles.formwrapper}>
            <form className={styles.form} onSubmit={handleSubmit}>
              <div className={styles.inputgrouprow}>
                <input
                  className={styles.input}
                  placeholder="First name"
                  required
                  type="text"
                  id="first_name"
                  name="first_name"
                  value={formData.first_name}
                  onChange={handleChange}
                />
                <input
                  className={styles.input}
                  placeholder="Last name"
                  required
                  type="text"
                  id="last_name"
                  name="last_name"
                  value={formData.last_name}
                  onChange={handleChange}
                />
              </div>
              <div className={styles.inputgroup}>
                <input
                  className={styles.input}
                  placeholder="Email"
                  required
                  type="text"
                  id="email"
                  name="email"
                  value={formData.email}
                  onChange={handleChange}
                />
              </div>
              <div className={styles.inputgroup}>
                <input
                  className={styles.input}
                  placeholder="Display Name"
                  required
                  type="text"
                  id="display_name"
                  name="display_name"
                  value={formData.display_name}
                  onChange={handleChange}
                />
              </div>

              <div className={styles.inputgrouprow}>
                <input
                  className={styles.input}
                  placeholder="Password"
                  required
                  type="password"
                  id="password"
                  name="password"
                  value={formData.password}
                  onChange={handleChange}
                />
                <input
                  className={styles.input}
                  required
                  placeholder="Confirm Password"
                  type="password"
                  id="confirmPassword"
                  name="confirmPassword"
                  value={formData.confirmPassword}
                  onChange={handleChange}
                />
              </div>
              <div className={styles.inputgroup}>
                <label htmlFor="dob">
                  Date of Birth:
                  <input
                    className={styles.input}
                    required
                    type="date"
                    id="dob"
                    name="dob"
                    value={formData.dob}
                    onChange={handleChange}
                  />
                </label>
              </div>
              <div className={styles.inputgroup}>
                <label htmlFor="avatar_path">
                  Profile Picture:
                  <input
                    className={styles.input}
                    type="file"
                    id="avatar_path"
                    name="avatar_path"
                    onChange={handleChange}
                  />
                </label>
              </div>
              <div className={styles.inputgroup}>
                <textarea
                  maxLength={100}
                  placeholder="Tell us about yourself..."
                  id="about_me"
                  name="about_me"
                  value={formData.about_me}
                  onChange={handleChange}
                />
              </div>
              <div className={styles.inputgroup}>
                <button className={styles.button} type="submit">
                  Sign Up
                </button>
              </div>
            </form>
            <div style={{ height: '5%' }}>
              <Link className={styles.a} to="/signin">
                Already have an account? Sign in
              </Link>
            </div>
          </div>
        </div>
        <Snackbar
          open={snackbarOpen}
          onClose={() => {
            setSnackbarOpen(false);
            setError(null);
          }}
          message={error ? error : "Signed up successfully!"}
          type={snackbarType}
        />
      </div>
    </Container>
  );
}
