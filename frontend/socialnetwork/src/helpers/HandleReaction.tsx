import { handleAPIRequest } from "../controllers/Api";
import { getCookie } from '../controllers/SetUserContextAndCookie';

const REACTION_ENDPOINT = "/reaction";

// HandleReaction handles a reaction for a post or comment
export async function HandleReaction(
  creatorId: string,
  reactionOn: string,
  reactionOnId: string,
  reactionType: string,
) {
  const payload = {
    creatorId,
    reactionOn,
    reactionOnId,
    reactionType,
  };

  const data = { data: payload };
  const options = {
    method: "POST",
    headers: {
      Authorization: "Bearer " + getCookie("sessionToken"),
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  };

  try {
    const response = await handleAPIRequest(REACTION_ENDPOINT, options);
    if (response.error) {
      console.error("Error response:", response.error);
    }
   
  } catch (error) {
    console.log("Error in handleLike:", error);
    throw error;
  }
}
