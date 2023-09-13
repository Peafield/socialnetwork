import { handleAPIRequest } from "../Api";
import { getCookie } from "../SetUserContextAndCookie";

export const DeleteNotification = async (notification_id: string) => {
    const deleteNotification = {
        notification_id: notification_id
    }
    const data = { data: deleteNotification };

    const url = `/notification`;
    const token = getCookie("sessionToken");

    const options = {
        method: "DELETE",
        headers: {
            Authorization: "Bearer " + token,
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    };

    try {
        const response = await handleAPIRequest(url, options);

        if (response.status !== "success") {
            throw new Error(response.message);
        }

        return response

    } catch (error) {
        if (error instanceof Error) {
            throw new Error("Failed to delete user notification: " + error.message);
        } else {
            throw new Error("An unexpected error occurred while fetching user notifications.");
        }
    }
}

