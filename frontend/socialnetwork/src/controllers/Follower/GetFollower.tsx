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

        const follower = response.data

        return follower

    } catch (error) {
        return error
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

        const followers = response.data

        return followers

    } catch (error) {
        return error
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

        const followees = response.data

        return followees

    } catch (error) {
        return error
    }
}