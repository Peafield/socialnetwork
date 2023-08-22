import { handleAPIRequest } from "../Api";
import { getCookie } from "../SetUserContextAndCookie";

export const NewFollower = async (followee_id: string) => {
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
        console.log(options.body);
  
        const response = await handleAPIRequest("/follow", options);
        if (response && response.status === "success") {
          return response
        }
      } catch (error) {
        return error
      }
}