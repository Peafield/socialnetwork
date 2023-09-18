import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { handleAPIRequest } from '../../controllers/Api'
import { getCookie } from '../../controllers/SetUserContextAndCookie'
import styles from './Group.module.css'

export interface GroupProps {
    group_id: string
    title: string
    description: string
    creator_id: string
    creator_name: string
    creation_date: string
}

interface NewGroupFormData {
    title: string
    description: string
}

const CreateGroup = () => {
    const navigate = useNavigate()
    const [formData, setFormData] = useState<NewGroupFormData>({
        title: "",
        description: ""
    })
    const [snackbarOpen, setSnackbarOpen] = useState<boolean>(false);
    const [snackbarType, setSnackbarType] = useState<
        "success" | "error" | "warning"
    >("error");
    const [error, setError] = useState<string | null>(null)

    const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        const { name, value } = e.target;
        setFormData({
            ...formData,
            [name]: value,
        });
    };

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
            const response = await handleAPIRequest("/group", options);
            if (response && response.status === "success") {
                setSnackbarType("success");
                setSnackbarOpen(true);
                setTimeout(() => {
                    navigate("/dashboard/group/" + formData.title);
                }, 1000);
            }
        } catch (error) {
            if (error instanceof Error) {
                setError("Could not create group");
                setSnackbarType("error");
                setSnackbarOpen(true);
            } else {
                setError("An unexpected error occurred");
                setSnackbarType("error");
                setSnackbarOpen(true);
            }
        }
    };


    return (
        <div className={styles.creategroupcontainer}>
            <form onSubmit={handleSubmit}>
                <div className={styles.inputgroup}>
                    <input
                        className={styles.input}
                        type="text"
                        placeholder="Group Title"
                        value={formData.title}
                        name="title"
                        onChange={handleChange}
                    />
                </div>
                <div className={styles.inputgroup}>
                    <textarea
                        maxLength={100}
                        className={styles.input}
                        placeholder="Description"
                        value={formData.description}
                        name="description"
                        onChange={handleChange}
                    />
                </div>
                <div className={styles.inputgroup}>
                    <button
                        className={styles.button}
                        type="submit"
                        onClick={() => {
                            setSnackbarOpen(false);
                            setError(null);
                        }}
                    >
                        Create New Group
                    </button>
                </div>
            </form>
        </div>
    )
}

export default CreateGroup