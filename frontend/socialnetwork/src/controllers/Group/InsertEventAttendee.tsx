import { handleAPIRequest } from "../Api";
import { getCookie } from "../SetUserContextAndCookie";

export const InsertEventAttendee = async (eventId: string, attendeeId: string, attending: boolean) => {
    const insertEventAttendee = {
        attendee_id: attendeeId,
        event_id: eventId,
        attending_status: attending ? 1 : 0
    }
    const data = { data: insertEventAttendee };

    const options = {
        method: "POST",
        headers: {
            Authorization: "Bearer " + getCookie("sessionToken"),
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    };
    try {
        const response = await handleAPIRequest("/eventattendees", options);
        if (response && response.status === "success") {
            console.log("insert success");
        }
    } catch (error) {
        if (error instanceof Error) {
            throw new Error("Failed to fetch insert event attendee: " + error.message);
        } else {
            throw new Error("An unexpected error occurred while fetching insert event attendee.");
        }
    }
}