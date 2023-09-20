import { GroupProps } from "../../components/Group/Group";
import { handleAPIRequest } from "../Api";
import { getUserByUserID } from "../GetUser";
import { getCookie } from "../SetUserContextAndCookie";

export const GetAllGroups = async () => {
    const url = `/group`

    const options = {
        method: "GET",
        headers: {
            Authorization: "Bearer " + getCookie("sessionToken"),
            "Content-Type": "application/json",
        },
    };
    try {
        const response = await handleAPIRequest(url, options);

        const groups = response.data.Groups

        return groups

    } catch (error) {
        if (error instanceof Error) {
            throw new Error("Failed to fetch all groups: " + error.message);
        } else {
            throw new Error("An unexpected error occurred while fetching all groups.");
        }
    }
}

export const getGroupByName = async (name: string) => {
    const url = `/group?group_title=${encodeURIComponent(name)}`

    const options = {
        method: "GET",
        headers: {
            Authorization: "Bearer " + getCookie("sessionToken"),
            "Content-Type": "application/json",
        },
    };
    try {
        const response = await handleAPIRequest(url, options);

        const group = response.data

        return group

    } catch (error) {
        if (error instanceof Error) {
            throw new Error("Failed to fetch group: " + error.message);
        } else {
            throw new Error("An unexpected error occurred while fetching group.");
        }
    }
}

export const getGroupByID = async (groupId: string) => {
    const url = `/group?group_id=${encodeURIComponent(groupId)}`

    const options = {
        method: "GET",
        headers: {
            Authorization: "Bearer " + getCookie("sessionToken"),
            "Content-Type": "application/json",
        },
    };
    try {
        const response = await handleAPIRequest(url, options);

        const group = response.data

        return group

    } catch (error) {
        if (error instanceof Error) {
            throw new Error("Failed to fetch group: " + error.message);
        } else {
            throw new Error("An unexpected error occurred while fetching group.");
        }
    }
}

export const GetUserGroups = async (userId: string) => {
    const url = `/group?user_id=${encodeURIComponent(userId)}`

    const options = {
        method: "GET",
        headers: {
            Authorization: "Bearer " + getCookie("sessionToken"),
            "Content-Type": "application/json",
        },
    };
    try {
        const response = await handleAPIRequest(url, options);

        const groups = response.data.Groups

        return groups

    } catch (error) {
        if (error instanceof Error) {
            throw new Error("Failed to fetch user groups: " + error.message);
        } else {
            throw new Error("An unexpected error occurred while fetching user groups.");
        }
    }
}