import { handleAPIRequest } from "../Api";
import { getCookie } from "../SetUserContextAndCookie";

export const GetGroupMembers = async (groupId: string) => {
    const url = `/groupmembers?group_id=${encodeURIComponent(groupId)}`

    const options = {
        method: "GET",
        headers: {
            Authorization: "Bearer " + getCookie("sessionToken"),
            "Content-Type": "application/json",
        },
    };
    try {
        const response = await handleAPIRequest(url, options);

        const groupMembers = response.data.GroupMembers

        return groupMembers

    } catch (error) {
        if (error instanceof Error) {
            throw new Error("Failed to fetch group members: " + error.message);
        } else {
            throw new Error("An unexpected error occurred while fetching group members.");
        }
    }
}