import React, { useContext, useState } from "react";
import { handleAPIRequest } from "../../controllers/Api";
import { UserContext } from "../../context/AuthContext";
import { MdOutlineLogout } from "react-icons/md";
import { useNavigate } from "react-router-dom";

export default function LogoutButton() {
  const navigate = useNavigate()
  const userContext = useContext(UserContext);
  const [error, setError] = useState<string | null>(null);

  const signOut = async () => {
    const options = {
      method: "POST",
      headers: {
        Authorization: "Bearer " + userContext.user?.authToken,
        "Content-Type": "application/json",
      },
    };
    try {      
      const response = await handleAPIRequest("/signout", options)
      console.log(response);
      
      navigate("/signin")
    } catch (error) {
      if (error instanceof Error) {
        setError(error.message);
      } else {
        setError("An unexpected error occurred.");
      }
    }
  };

  return (
    <button
      onClick={signOut}
      style={{
        width: "auto",
        height: "auto",
        color: "black",
        fontSize: "2rem",
        fontWeight: "bold",
        backgroundColor: "#fa4d6a00",
      }}>
      <MdOutlineLogout />
    </button>
  )
}
