import { handleAPIRequest } from "../Api";
import { getCookie } from "../SetUserContextAndCookie";

export const unfollow = async (followee_id: string, follower_id: string) => {
  const deleteFollower = {
    followee_id: followee_id,
    follower_id: follower_id
  }
  const data = { data: deleteFollower };

  const options = {
    method: "DELETE",
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
      throw new Error("Failed to fetch unfollow: " + error.message);
    } else {
      throw new Error("An unexpected error occurred while fetching unfollow.");
    }
  }
}