import { handleAPIRequest } from "../Api";
import { getCookie } from "../SetUserContextAndCookie";

export const UpdateGroupRequestPending = async (memberId: string, groupId: string, accept: boolean) => {
    const updateGroupMember = {
        member_id: memberId,
        group_id: groupId,
        accepted: accept
    }
    const data = { data: updateGroupMember };

    const options = {
        method: "PUT",
        headers: {
            Authorization: "Bearer " + getCookie("sessionToken"),
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    };
    try {
        const response = await handleAPIRequest("/groupmembers", options);
        if (response && response.status === "success") {
            console.log("update success");
        }
    } catch (error) {
        if (error instanceof Error) {
            throw new Error("Failed to fetch update group member: " + error.message);
        } else {
            throw new Error("An unexpected error occurred while fetching update group member.");
        }
    }
}