import { handleAPIRequest } from "../Api";
import { getCookie } from "../SetUserContextAndCookie";

export const newFollower = async (followee_id: string) => {
  const newFollower = {
    followee_id: followee_id
  }
  const data = { data: newFollower };

  const options = {
    method: "POST",
    headers: {
      Authorization: "Bearer " + getCookie("sessionToken"),
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  };
  try {
    const response = await handleAPIRequest("/follow", options);
    if (response && response.status === "success") {
      return response
    }
  } catch (error) {
    if (error instanceof Error) {
      throw new Error("Failed to fetch new follower: " + error.message);
    } else {
      throw new Error("An unexpected error occurred while fetching new follower.");
    }
  }
}