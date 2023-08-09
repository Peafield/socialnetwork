import React, { useState } from "react";
import { useAuth } from "../../hooks/useAuth";
import { handleAPIRequest } from "../../controllers/Api";

const LogoutButton = () => {
  const { user, clearUser } = useAuth();
  const [error, setError] = useState<string | null>(null)

  const signOut = async () => {
    const options = {
      method: "POST",
      headers: {
        Authorization: "Bearer" + user?.authToken,
        "Content-Type": "application/json",
      },
    };
    try {
      await handleAPIRequest("/signout", options);
      clearUser();

    } catch (error) {
      setError(error.message);
    }
  };

  return <button onClick={signOut}>Logout</button>;
};

export default LogoutButton;
