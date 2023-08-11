import React, { useContext, useState } from "react";
import { handleAPIRequest } from "../../controllers/Api";
import { UserContext } from "../../context/AuthContext";

export default function LogoutButton() {
  const userContext = useContext(UserContext);
  const [error, setError] = useState<string | null>(null);

  const signOut = async () => {
    const options = {
      method: "POST",
      headers: {
        // Authorization: "Bearer" + userContext.user?.authToken,
        "Content-Type": "application/json",
      },
    };
    try {
      await handleAPIRequest("/signout", options);
    } catch (error) {
      if (error instanceof Error) {
        setError(error.message);
      } else {
        setError("An unexpected error occurred.");
      }
    }
  };

  return <button onClick={signOut}>Logout</button>;
}
