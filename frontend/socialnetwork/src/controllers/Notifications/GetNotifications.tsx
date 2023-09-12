import { handleAPIRequest } from "../Api";
import { getCookie } from "../SetUserContextAndCookie";

export const GetNotifications = async () => {
    const url = `/notification`;
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

        return response.data.Notifications;
    } catch (error) {
        if (error instanceof Error) {
            throw new Error("Failed to fetch user notifications: " + error.message);
        } else {
            throw new Error("An unexpected error occurred while fetching user notifications.");
        }
    }
}

