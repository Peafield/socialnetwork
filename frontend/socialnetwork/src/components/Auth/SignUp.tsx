import React, { useState, FormEvent, useContext } from "react";
import { Link, redirect, useNavigate } from "react-router-dom";
import { handleAPIRequest } from "../../controllers/Api";
import { UserContext } from "../../context/AuthContext";
import styles from "./Auth.module.css";
import Container from "../Containers/Container";

interface FormData {
  email: string;
  display_name: string;
  password: string;
  confirmPassword: string;
  first_name: string;
  last_name: string;
  dob: string;
  avatar_path: File | null;
  about_me: string;
}

export default function SignUp() {
  const navigate = useNavigate();
  const userContext = useContext(UserContext);
  const [error, setError] = useState<string | null>(null);
  const [formData, setFormData] = useState<FormData>({
    email: "",
    display_name: "",
    password: "",
    confirmPassword: "",
    first_name: "",
    last_name: "",
    dob: "",
    avatar_path: null,
    about_me: "",
  });

  const handleChange = (e: {
    target: { type?: any; files?: any; name?: any; value?: any };
  }) => {
    const { name, value } = e.target;

    if (e.target.type === "file") {
      const file = e.target.files ? e.target.files[0] : null;
      setFormData((prevState) => ({
        ...prevState,
        avatar_path: file,
      }));
    } else {
      setFormData((prevState) => ({
        ...prevState,
        [name]: value,
      }));
    }
  };

  const HandleSubmit = async (e: FormEvent) => {
    e.preventDefault();

    if (formData.password !== formData.confirmPassword) {
      alert("passwords do not match!");
      return;
    }

    const data = formData;

    const options = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    };
    try {
      const response = await handleAPIRequest("/signup", options);
      const user = {
        usernameEmail: data.email,
        authToken: response.Data.token,
      };

      userContext.setUser(user);

      navigate('/dashboard')
    } catch (error) {
      if (error instanceof Error) {
        setError(error.message);
      } else {
        setError("An unexpected error occurred.");
      }
    }
  };

  return (
    <Container>
      <div className={styles.authcontainer}>
        <div className={styles.formwrapper}>
          <form onSubmit={HandleSubmit}>
            <h2 className={styles.h2}>Sign Up</h2>
            <div className={styles.inputgrouprow}>
              <label htmlFor="first_name">
                First Name:
                <input
                  className={styles.input}
                  required
                  type="text"
                  id="first_name"
                  name="first_name"
                  value={formData.first_name}
                  onChange={handleChange}
                />
              </label>

              <label htmlFor="last_name">
                Last Name:
                <input
                  className={styles.input}
                  required
                  type="text"
                  id="last_name"
                  name="last_name"
                  value={formData.last_name}
                  onChange={handleChange}
                />
              </label>
            </div>
            <div className={styles.inputgroup}>
              <label className={styles.label} htmlFor="email">
                Email:
                <input
                  className={styles.input}
                  required
                  type="text"
                  id="email"
                  name="email"
                  value={formData.email}
                  onChange={handleChange}
                />
              </label>
            </div>
            <div className={styles.inputgroup}>
              <label htmlFor="display_name">
                Display Name:
                <input
                  className={styles.input}
                  required
                  type="text"
                  id="display_name"
                  name="display_name"
                  value={formData.display_name}
                  onChange={handleChange}
                />
              </label>
            </div>

            <div className={styles.inputgroup}>
              <label htmlFor="password">
                Password:
                <input
                  className={styles.input}
                  required
                  type="password"
                  id="password"
                  name="password"
                  value={formData.password}
                  onChange={handleChange}
                />
              </label>
            </div>
            <div className={styles.inputgroup}>
              <label htmlFor="confirmPassword">
                Confirm Password:
                <input
                  className={styles.input}
                  required
                  type="password"
                  id="confirmPassword"
                  name="confirmPassword"
                  value={formData.confirmPassword}
                  onChange={handleChange}
                />
              </label>
            </div>

            <div className={styles.inputgrouprow}>
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
              <label htmlFor="avatar_path">
                Profile Picture:
                <input
                  className={styles.customfileupload}
                  type="file"
                  id="avatar_path"
                  name="avatar_path"
                  onChange={handleChange}
                />
              </label>
            </div>
            <div className={styles.inputgroup}>
              <label htmlFor="about_me">
                About Me:
                <textarea
                  required
                  maxLength={100}
                  placeholder="Tell us about yourself..."
                  id="about_me"
                  name="about_me"
                  value={formData.about_me}
                  onChange={handleChange}
                />
              </label>
            </div>
            <div className={styles.inputgroup}>
              <button className={styles.button} type="submit">
                Sign Up
              </button>
            </div>
          </form>
          <Link to="/signin">Already have an account? Sign in</Link>
        </div>
      </div>
    </Container>
  );
}
