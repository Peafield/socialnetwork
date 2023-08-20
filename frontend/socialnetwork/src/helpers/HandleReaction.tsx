import { handleAPIRequest } from "../controllers/Api";

const REACTION_ENDPOINT = "/reaction"

export async function HandleReaction(reactionOn: string, postORCommentId: string, reactionType: 'like' | 'dislike', action: 'add' | 'remove') {
    const payload = {
      reactionOn,
        postORCommentId,
        type: reactionType,
        action,
    };

    const options = {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(payload),
    }

  try {
    const response = await handleAPIRequest(REACTION_ENDPOINT, options)
    if (response && response.status === "success") {
        return "success"
    }
  } catch (error) {
    console.log("Error in handleLike:", error)
    throw error
  }
}