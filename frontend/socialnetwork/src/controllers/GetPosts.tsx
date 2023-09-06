import { handleAPIRequest } from "./Api";
import { getCookie } from "./SetUserContextAndCookie";

export const getUserPosts = async (userId: string) => {
    const url = `/post?user_id=${encodeURIComponent(userId)}`;
    const token = getCookie("sessionToken");

    const options = {
        method: "GET",
        headers: {
            Authorization: "Bearer " + token,
            "Content-Type": "application/json",
        },
    };

    try {
        const response = await handleAPIRequest(url, options);

        if (response.status !== "success") {
            throw new Error(response.message);
        }

        return response.data.Posts;
    } catch (error) {
        if (error instanceof Error) {
            throw new Error("Failed to fetch user posts: " + error.message);
        } else {
            throw new Error("An unexpected error occurred while fetching user posts.");
        }
    }
};
