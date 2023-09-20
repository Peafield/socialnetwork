import { handleAPIRequest } from "../Api";
import { getCookie } from "../SetUserContextAndCookie";

export const UpdateEventAttendeeStatus = async (eventId: string, attendeeId: string, attending: boolean) => {
    const updateEventAttendee = {
        attendee_id: attendeeId,
        event_id: eventId,
        attending_status: attending ? 1 : 0
    }
    const data = { data: updateEventAttendee };

    const options = {
        method: "PUT",
        headers: {
            Authorization: "Bearer " + getCookie("sessionToken"),
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    };
    try {
        const response = await handleAPIRequest("/eventattendees", options);
        if (response && response.status === "success") {
            console.log("update success");
        }
    } catch (error) {
        if (error instanceof Error) {
            throw new Error("Failed to fetch update event attendee status: " + error.message);
        } else {
            throw new Error("An unexpected error occurred while fetching update event attendee status.");
        }
    }
}