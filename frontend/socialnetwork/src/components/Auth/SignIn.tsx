import React, { ChangeEvent, useContext, useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { handleAPIRequest } from "../../controllers/Api";
import { UserContext } from "../../context/AuthContext";
import Container from "../Containers/Container";
import styles from "./Auth.module.css";

interface SignInFormData {
  username_email: string;
  password: string;
}

export default function SignIn() {
  const navigate = useNavigate()
  const userContext = useContext(UserContext);
  const [formData, setFormData] = useState<SignInFormData>({
    username_email: "",
    password: "",
  });
  const [error, setError] = useState<string | null>(null);

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
      const user = {
        usernameEmail: data.data.username_email,
        authToken: response.Data.token,
      };

      userContext.setUser(user);

      navigate("/dashboard")
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
    </Container>
  );
}
