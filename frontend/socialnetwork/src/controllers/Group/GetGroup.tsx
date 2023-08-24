import { handleAPIRequest } from "../Api";
import { getCookie } from "../SetUserContextAndCookie";

export const getGroupByName = async (group_id: string) => {
    const url = `/group?group_id=${encodeURIComponent(group_id)}`

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

        const group = response.data

        return group

    } catch (error) {
        return error
    }
}