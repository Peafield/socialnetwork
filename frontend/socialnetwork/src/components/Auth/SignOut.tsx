import React, { useContext, useState } from "react";
import { handleAPIRequest } from "../../controllers/Api";
import { UserContext } from "../../context/AuthContext";
import { MdOutlineLogout } from "react-icons/md";
import { useNavigate } from "react-router-dom";

export default function LogoutButton() {
  const navigate = useNavigate()
  const userContext = useContext(UserContext);
  const [isHovered, setIsHovered] = useState(false);
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

  const buttonStyle = {
    width: 'auto',
    height: 'auto',
    backgroundColor: 'rgba(211, 211, 211, 0)',
    fontSize: 'x-large',
    fontWeight: 'bold',
    color: isHovered ? '#fa8fa1' : '#fa4d6a', // Change color on hover
    transition: 'background-color 0.3s', // Add a transition for smooth color change
  };

  return (
    <button
      onClick={signOut}
      style={buttonStyle}
      onMouseEnter={() => setIsHovered(true)}
      onMouseLeave={() => setIsHovered(false)}>
      <MdOutlineLogout />
    </button>
  )
}
