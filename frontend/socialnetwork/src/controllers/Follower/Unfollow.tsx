import { handleAPIRequest } from "../Api";
import { getCookie } from "../SetUserContextAndCookie";

export const Unfollow = async (followee_id: string, follower_id: string) => {
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
        console.log(options.body);
  
        const response = await handleAPIRequest("/follow", options);
        if (response && response.status === "success") {
          return response
        }
      } catch (error) {
        return error
      }
}