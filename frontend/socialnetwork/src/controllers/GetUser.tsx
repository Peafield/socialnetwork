import { handleAPIRequest } from "./Api";
import { getCookie } from "./SetUserContextAndCookie";

export const getUserByDisplayName = async (display_name: string) => {
    const url = `/user?display_name=${encodeURIComponent(display_name)}`

    const options = {
        method: "GET",
        headers: {
            Authorization: "Bearer " + getCookie("sessionToken"),
            "Content-Type": "application/json",
        },
    };
    try {
        const response = await handleAPIRequest(url, options);

        const newprofile = response.data.UserInfo
        const avatar = response.data.ProfilePic

        newprofile.avatar = avatar

        return newprofile

    } catch (error) {
        return error
    }
}

export const getUserByUserID = async (user_id: string) => {
    const url = `/user?user_id=${encodeURIComponent(user_id)}`

    const options = {
        method: "GET",
        headers: {
            Authorization: "Bearer " + getCookie("sessionToken"),
            "Content-Type": "application/json",
        },
    };
    try {
        const response = await handleAPIRequest(url, options);

        const newprofile = response.data.UserInfo
        const avatar = response.data.ProfilePic

        newprofile.avatar = avatar

        return newprofile

    } catch (error) {
        return error
    }
}