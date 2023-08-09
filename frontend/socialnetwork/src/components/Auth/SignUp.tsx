import React, { useState, FormEvent } from "react";
import { handleAPIRequest } from "../../controllers/Api";
import { useAuth } from "../../hooks/useAuth";

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
  const { login } = useAuth();
  const [error, setError] = useState(null);
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

  const handleChange = (e: { target: { type?: any; files?: any; name?: any; value?: any; }; }) => {
    const { name, value } = e.target;

    if (name === "password" || name === "confirmPassword") {
      if (formData.password !== formData.confirmPassword) {
        alert("passwords do not match!");
        return;
      }
    }

    if (e.target.type === "file") {
      const file = e.target.files ? e.target.files[0] : null;
      setFormData(prevState => ({
          ...prevState,
          avatar_path: file
      }));

    }

    setFormData(prevState => ({
      ...prevState,
      [name]: value,
    }));
  };

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();
    const data = formData;
    const options = {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(data)
    };
    try {
      const response = await handleAPIRequest("/signup", options);
      const user = {
        usernameEmail: data.email,
        authToken: response.Data.token,
      };

      login(user);
    } catch (error) {
      setError(error.message);
    }
  };

  return (
    <>
      <div className="auth-container">
        <h2>Sign Up</h2>
        <form onSubmit={handleSubmit}>
          <div className="input-group">
            <label htmlFor="email">Email:</label>
            <input
              required
              type="text"
              id="email"
              name="email"
              value={formData.email}
              onChange={handleChange}
            />
          </div>
          <div className="input-group">
            <label htmlFor="display_name">Display Name:</label>
            <input
              required
              type="text"
              id="display_name"
              name="display_name"
              value={formData.display_name}
              onChange={handleChange}
            />
          </div>
          <div className="input-group">
            <label htmlFor="password">Password:</label>
            <input
              required
              type="text"
              id="password"
              name="password"
              value={formData.password}
              onChange={handleChange}
            />
            <label htmlFor="confirmPassword">Confirm Password:</label>
            <input
              required
              type="text"
              id="confirmPassword"
              name="confirmPassword"
              value={formData.confirmPassword}
              onChange={handleChange}
            />
          </div>
          <div className="input-group">
            <label htmlFor="first_name">First Name:</label>
            <input
              required
              type="text"
              id="first_name"
              name="first_name"
              value={formData.first_name}
              onChange={handleChange}
            />
          </div>
          <div className="input-group">
            <label htmlFor="last_name">Last Name:</label>
            <input
              required
              type="text"
              id="last_name"
              name="last_name"
              value={formData.last_name}
              onChange={handleChange}
            />
          </div>
          <div className="input-group">
            <label htmlFor="dob">Date of Birth:</label>
            <input
              required
              type="date"
              id="dob"
              name="dob"
              value={formData.dob}
              onChange={handleChange}
            />
          </div>
          <div className="input-group">
            <label htmlFor="avatar_path">Date of Birth:</label>
            <input
              required
              type="date"
              id="dob"
              name="dob"
              value={formData.dob}
              onChange={handleChange}
            />
          </div>
          <div className="input-group">
            <label htmlFor="avatar_path">Profile Picture:</label>
            <input
              type="file"
              id="avatar_path"
              name="avatar_path"
              onChange={handleChange}
            />
          </div>
          <div className="input-group">
            <label htmlFor="about_me">About Me:</label>
            <textarea
              required
              maxLength={100}
              placeholder="Tell us about yourself..."
              id="about_me"
              name="about_me"
              value={formData.about_me}
              onChange={handleChange}
            />
          </div>
          <button type="submit">Sign Up</button>
        </form>
      </div>
    </>
  );
}
