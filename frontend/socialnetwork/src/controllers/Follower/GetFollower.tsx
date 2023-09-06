import { handleAPIRequest } from "../Api";
import { getCookie } from "../SetUserContextAndCookie";

export const getFollowerData = async (followee_id: string) => {
    const url = `/follow?followee_id=${encodeURIComponent(followee_id)}`

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

        const follower = response.data

        return follower

    } catch (error) {
        if (error instanceof Error) {
            throw new Error("Failed to fetch follower data: " + error.message);
        } else {
            throw new Error("An unexpected error occurred while fetching follower data.");
        }
    }
}

export const getFollowers = async (followee_id: string) => {
    const url = `/follow?followers_id=${encodeURIComponent(followee_id)}`

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

        const followers = response.data

        return followers

    } catch (error) {
        if (error instanceof Error) {
            throw new Error("Failed to fetch followers: " + error.message);
        } else {
            throw new Error("An unexpected error occurred while fetching followers.");
        }
    }
}

export const getFollowees = async (follower_id: string) => {
    const url = `/follow?followees_id=${encodeURIComponent(follower_id)}`

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

        const followees = response.data

        return followees

    } catch (error) {
        if (error instanceof Error) {
            throw new Error("Failed to fetch followees: " + error.message);
        } else {
            throw new Error("An unexpected error occurred while fetching followees.");
        }
    }
}