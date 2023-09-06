import { useContext } from "react";
import { UserContext } from "../context/AuthContext";

export function useSetUserContextAndCookie() {
  const userContext = useContext(UserContext);

  return (token: string) => {
    try {
      const base64Url = token.split(".")[1];
      const base64 = base64Url.replace(/-/g, "+").replace(/_/g, "/");
      const payload = JSON.parse(
        decodeURIComponent(
          atob(base64)
            .split("")
            .map((c) => {
              return "%" + ("00" + c.charCodeAt(0).toString(16)).slice(-2);
            })
            .join("")
        )
      );      

      const user = {
        userId: payload.user_id,
        displayName: payload.display_name,
        firstName: payload.first_name,
        lastName: payload.last_name,
        isLoggedIn: payload.is_logged_in,
        role: payload.role,
        exp: payload.exp,
        authToken: token,
      };

      userContext.setUser(user);
      setCookie("sessionToken", token, user.exp);
    } catch (error) {
      return error;
    }
  };
}

function setCookie(name: string, value: string, expiryTimestamp: number): void {
  const date = new Date(expiryTimestamp * 1000);
  const expires = "; expires=" + date.toUTCString();
  document.cookie = name + "=" + (value || "") + expires + "; path=/";
}

export function getCookie(name: string) {
  return document.cookie.split(name+"=")[1].split(";")[0]
}
