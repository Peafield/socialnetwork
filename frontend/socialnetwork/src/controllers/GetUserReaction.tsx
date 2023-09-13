import { handleAPIRequest } from "./Api";
import { getCookie } from "./SetUserContextAndCookie";

export const GetUserPostReaction = async (postId: string) => {
    const url = `/reaction?post_id=${encodeURIComponent(postId)}`

    const options = {
        method: "GET",
        headers: {
            Authorization: "Bearer " + getCookie("sessionToken"),
            "Content-Type": "application/json",
        },
    };
    try {
        const response = await handleAPIRequest(url, options);

        if (response.status !== "success") {
            throw new Error(response.message);
        }

        return response.data

    } catch (error) {
        if (error instanceof Error) {
            throw new Error("Failed to fetch user by id: " + error.message);
        } else {
            throw new Error("An unexpected error occurred while fetching user by id.");
        }
    }
}

export const GetUserCommentReaction = async (commentId: string) => {
    const url = `/reaction?comment_id=${encodeURIComponent(commentId)}`

    const options = {
        method: "GET",
        headers: {
            Authorization: "Bearer " + getCookie("sessionToken"),
            "Content-Type": "application/json",
        },
    };
    try {
        const response = await handleAPIRequest(url, options);

        if (response.status !== "success") {
            throw new Error(response.message);
        }

        return response.data

    } catch (error) {
        if (error instanceof Error) {
            throw new Error("Failed to fetch user by id: " + error.message);
        } else {
            throw new Error("An unexpected error occurred while fetching user by id.");
        }
    }
}