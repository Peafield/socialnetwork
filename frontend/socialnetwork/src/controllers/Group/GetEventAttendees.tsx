import { handleAPIRequest } from "../Api";
import { getCookie } from "../SetUserContextAndCookie";

export const GetEventAttendees = async (eventId: string) => {
    const url = `/eventattendees?event_id=${encodeURIComponent(eventId)}`

    const options = {
        method: "GET",
        headers: {
            Authorization: "Bearer " + getCookie("sessionToken"),
            "Content-Type": "application/json",
        },
    };
    try {
        const response = await handleAPIRequest(url, options);

        const attendees = response.data.GroupEventAttendees

        return attendees

    } catch (error) {
        if (error instanceof Error) {
            throw new Error("Failed to fetch group members: " + error.message);
        } else {
            throw new Error("An unexpected error occurred while fetching group members.");
        }
    }
}