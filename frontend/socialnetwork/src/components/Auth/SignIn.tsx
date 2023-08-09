import React, { ChangeEvent, useContext, useState } from "react";
import {Link} from "react-router-dom";
import { handleAPIRequest } from "../../controllers/Api";
import { UserContext } from "../../context/AuthContext";

interface SignInFormData {
  usernameEmail: string;
  password: string;
}

export default function SignIn() {
  const userContext = useContext(UserContext)
  const [formData, setFormData] = useState<SignInFormData>({
    usernameEmail: "",
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

  const handleSubmit = async (e: { preventDefault: () => void; }) => {
    e.preventDefault();
    const data = formData;
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
        usernameEmail: data.usernameEmail,
        authToken: response.Data.token,
      };

      userContext.setUser(user);
    } catch (error) {
        if (error instanceof Error) {
            setError(error.message);
        } else {
            setError('An unexpected error occurred.');
        }
    }
  };

  return (
    <>
      <div>
        <form onSubmit={handleSubmit}>
          <label>
            Username/Email:
            <input
              type="text"
              value={formData.usernameEmail}
              name="usernameEmail"
              onChange={handleChange}
            />
          </label>
          <label>
            Password:
            <input
              type="password"
              value={formData.password}
              name="password"
              onChange={handleChange}
            />
          </label>
          <label>
            Submit
            <button type="submit" />
          </label>
        </form>
      </div>
      <div>
        <Link to="/signup">Don't have an account? Sign up</Link>
      </div>
    </>
  );
}