import { handleAPIRequest } from "../Api";
import { getCookie } from "../SetUserContextAndCookie";

export const updateFollowRequest = async (followerId: string, followingStatus: number) => {
    const updateFollower = {
        follower_id: followerId,
        following_status: followingStatus
    }
    const data = { data: updateFollower };

    const options = {
        method: "PUT",
        headers: {
            Authorization: "Bearer " + getCookie("sessionToken"),
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    };
    try {
        const response = await handleAPIRequest("/follow", options);
        if (response && response.status === "success") {
            console.log("follow success");

        }
    } catch (error) {
        if (error instanceof Error) {
            throw new Error("Failed to fetch update follower: " + error.message);
        } else {
            throw new Error("An unexpected error occurred while fetching update follower.");
        }
    }
}