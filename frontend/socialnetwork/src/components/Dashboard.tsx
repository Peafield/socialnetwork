import React, { useEffect, useState } from "react";
import { Route, Routes } from "react-router-dom";
import Container from "./Containers/Container";
import { NavBar } from "./NavBar";
import CreatePost from "./Post/CreatePost";
import PostFeed from "./Post/PostFeed";
import Profile from "./Profile/Profile";
import styles from "./Dashboard.module.css"
import Group from "./Group/Group";
import SideModal from "./Containers/SideModal";
import FriendsMessagingList from "./Chat/FriendsMessagingList";
import { useWebSocket } from "../Socket";
import { getCookie } from "../controllers/SetUserContextAndCookie";
import { LuMessagesSquare } from 'react-icons/lu'
import EditProfile from "./Profile/EditProfile";

export default function Dashboard() {
  const { message, sendMessage } = useWebSocket("ws://localhost:8080/ws", {
    headers: {
      Authorization: `Bearer ${getCookie("sessionToken")}`
    }
  });

  const [isModalOpen, setIsModalOpen] = useState(false);

  return (
    <Container>
      <div className={styles.dashboardcontainer}>
        <div className={styles.navbarcontainer}>
          <NavBar />
        </div>
        <div className={styles.mainelementcontainer}>
          <Routes>
            <Route path="/" element={<PostFeed />} />
            <Route path="/createpost" element={<CreatePost />} />
            <Route path="/user/:username" element={<Profile />} />
            <Route path="/user/edit/:username" element={<EditProfile />} />
            <Route path="/group/:groupname" element={<Group />} />
          </Routes>
        </div>
        <div className={styles.sidebarbuttoncontainer}>
          <button onClick={() => { setIsModalOpen(true) }} style={{padding: "5px"}}>
            <LuMessagesSquare />
          </button>
        </div>
        <SideModal open={isModalOpen} onClose={() => setIsModalOpen(false)}>
          <FriendsMessagingList message={message} sendMessage={sendMessage} />
        </SideModal>
      </div>

    </Container>
  );
}
