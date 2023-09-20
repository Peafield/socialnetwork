import React, { useContext, useEffect, useState } from 'react'
import { TiTick, TiTimes } from 'react-icons/ti'
import { UserContext } from '../../context/AuthContext'
import { handleAPIRequest } from '../../controllers/Api'
import { getUserByUserID } from '../../controllers/GetUser'
import { GetEventAttendees } from '../../controllers/Group/GetEventAttendees'
import { GetGroupEvents } from '../../controllers/Group/GetGroupEvent'
import { InsertEventAttendee } from '../../controllers/Group/InsertEventAttendee'
import { UpdateEventAttendeeStatus } from '../../controllers/Group/UpdateEventAttendeeStatus'
import { getCookie } from '../../controllers/SetUserContextAndCookie'
import Snackbar from '../feedback/Snackbar'
import { ProfileProps } from '../Profile/Profile'
import styles from './Group.module.css'

interface GroupEventsProps {
  group_id: string
}

interface EventProps {
  event_id: string
  group_id: string
  creator_id: string
  creator_name: string
  title: string
  description: string
  event_start_time: string
  respondants: EventAttendeeProps[]
  total_going: number
  total_not_going: number
  creation_date: string
}

interface EventAttendeeProps {
  event_id: string
  attendee_id: string
  attending_status: number
  creation_date: string
}

interface CreateEventFormData {
  group_id: string
  creator_id: string
  title: string
  description: string
  event_start_time: string
}

const GroupEvents: React.FC<GroupEventsProps> = ({
  group_id
}) => {
  const userContext = useContext(UserContext)
  const [groupEvents, setGroupEvents] = useState<EventProps[] | null>(null)
  const [formData, setFormData] = useState<CreateEventFormData>({
    group_id: group_id,
    creator_id: userContext.user ? userContext.user.userId : "",
    title: "",
    description: "",
    event_start_time: ""
  });
  const [currentEvent, setCurrentEvent] = useState<EventProps | null>(null)
  const [updateTrigger, setUpdateTrigger] = useState<number>(0)
  const [error, setError] = useState<string | null>(null);
  const [snackbarOpen, setSnackbarOpen] = useState<boolean>(false);
  const [snackbarType, setSnackbarType] = useState<
    "success" | "error" | "warning"
  >("error");


  useEffect(() => {
    const fetchData = async () => {
      try {
        const events: EventProps[] = await GetGroupEvents(group_id)

        const modifiedEventsPromises = events.map(async (event: EventProps) => {
          const user: ProfileProps = await getUserByUserID(event.creator_id)
          event.creator_name = user.display_name

          event.event_start_time = `${event.event_start_time.split("T")[0]} ${event.event_start_time.split("T")[1].split(":00Z")[0]}`

          const eventAttendees: EventAttendeeProps[] = await GetEventAttendees(event.event_id)
          event.respondants = eventAttendees

          event.total_going = eventAttendees ? eventAttendees.filter((attendee) => attendee.attending_status === 1).length : 0
          event.total_not_going = eventAttendees ? eventAttendees.filter((attendee) => attendee.attending_status === 0).length : 0
          return event
        })

        const modifiedEvents = await Promise.all(modifiedEventsPromises)

        setGroupEvents(modifiedEvents);

      } catch (error) {
        if (error instanceof Error) {
          setError(error.message);
        } else {
          setError("An unexpected error occurred.");
        }
      }
    }

    fetchData()
  }, [updateTrigger])

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;

    setFormData((prevState) => ({
      ...prevState,
      [name]: value,
    }));

  }

  const handleSubmit = async (e: { preventDefault: () => void }) => {
    e.preventDefault();
    const data = { data: formData };
    const options = {
      method: "POST",
      headers: {
        Authorization: "Bearer " + getCookie("sessionToken"),
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    };
    try {
      const response = await handleAPIRequest("/event", options);
      if (response && response.status === "success") {
        setSnackbarType("success");
        setSnackbarOpen(true);
        setUpdateTrigger((prev: number) => {
          return prev + 1
        })
      }

    } catch (error) {
      if (error instanceof Error) {
        setError("Could not create event");
        setSnackbarType("error");
        setSnackbarOpen(true);
      } else {
        setError("An unexpected error occurred");
        setSnackbarType("error");
        setSnackbarOpen(true);
      }
    }
  }

  const handleAccept = (event: EventProps) => {
    if (userContext.user && userAlreadyResponded(event.respondants, userContext.user?.userId)) {
      UpdateEventAttendeeStatus(event.event_id, userContext.user.userId, true)
    } else if (userContext.user) {
      InsertEventAttendee(event.event_id, userContext.user.userId, true)
    }
    setUpdateTrigger((prev) => prev + 1)
  }

  const handleDecline = (event: EventProps) => {
    if (userContext.user && userAlreadyResponded(event.respondants, userContext.user?.userId)) {
      UpdateEventAttendeeStatus(event.event_id, userContext.user.userId, false)
    } else if (userContext.user) {
      InsertEventAttendee(event.event_id, userContext.user.userId, false)
    }
    setUpdateTrigger((prev) => prev + 1)
  }

  return (
    <div className={styles.groupeventscontainer}>
      <div className={styles.createevent}>
        <form onSubmit={handleSubmit}>
          <div>
            <input
              required
              type="text"
              id="title"
              name="title"
              onChange={handleChange}
              placeholder="Event Title"
              value={formData.title}
              maxLength={20}
            />
          </div>
          <div>
            <input
              required
              type="text"
              id="description"
              name="description"
              onChange={handleChange}
              placeholder="Description"
              value={formData.description}
              maxLength={40}
            />
          </div>
          <div>
            <label htmlFor='event_start_time'>Start Time</label>
            <input
              required
              type="datetime-local"
              id="event_start_time"
              name="event_start_time"
              onChange={handleChange}
              value={formData.event_start_time}
            />
          </div>
          <div>
            <button type='submit'>Create Event</button>
          </div>

        </form>
      </div>
      <div className={styles.listofevents}>
        Events
        {groupEvents?.map((event) => (
          <div key={event.event_id} className={styles.event}>
            <div>{event.title}</div>
            <div>{event.description}</div>
            <div>Starts {event.event_start_time}</div>
            <div>Created by {event.creator_name}</div>
            <div>
              <span
                style={{ cursor: 'pointer', color: 'lightgreen' }}
                onClick={() => {
                  handleAccept(event)
                }}>
                <TiTick />{event.total_going}
              </span>
              <span style={{ cursor: 'pointer', color: 'darkred' }} onClick={() => {
                handleDecline(event)
              }}>
                <TiTimes />{event.total_not_going}
              </span>
            </div>
          </div>
        ))}
      </div>
      <Snackbar
        open={snackbarOpen}
        onClose={() => {
          setSnackbarOpen(false);
          setError(null);
        }}
        message={error ? error : "Event Created!"}
        type={snackbarType}
      />
    </div>
  )
}

function userAlreadyResponded(eventAttendees: EventAttendeeProps[], user_id: string): boolean {
  if (!eventAttendees) {
    return false
  }
  return eventAttendees.some((attendee) => attendee.attendee_id === user_id);
}

export default GroupEvents