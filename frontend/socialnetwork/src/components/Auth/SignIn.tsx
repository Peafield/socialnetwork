import React, { ChangeEvent, useState } from "react";
import {Link} from "react-router-dom";
import { handleAPIRequest } from "../../controllers/Api";
import { useAuth } from "../../hooks/useAuth";

interface SignInFormData {
  usernameEmail: string;
  password: string;
}

export default function SignIn() {
  const { setUser } = useAuth();
  const [formData, setFormData] = useState<SignInFormData>({
    usernameEmail: "",
    password: "",
  });
  const [error, setError] = useState(null);

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

      setUser(user);
    } catch (error) {
      setError(error.message);
    }
  };

  return (
    <>
      <div>
        <form>
          <label>
            Username/Email:
            <input
              type="text"
              value=""
              name="usernameEmail"
              onChange={handleChange}
            />
          </label>
          <label>
            Password:
            <input
              type="text"
              value=""
              name="password"
              onChange={handleChange}
            />
          </label>
          <label>
            Submit
            <button type="submit" onSubmit={handleSubmit} />
          </label>
        </form>
      </div>
      <div>
        <Link to="/signup">Don't have an account? Sign up</Link>
      </div>
    </>
  );
}
