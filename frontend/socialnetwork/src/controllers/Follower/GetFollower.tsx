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
        console.log(response);
        

        const follower = response.data

        return follower

    } catch (error) {
        return error
    }
}