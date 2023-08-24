import React, { useEffect, useState } from "react";
import { Route, Routes } from "react-router-dom";
import Container from "./Containers/Container";
import { NavBar } from "./NavBar";
import CreatePost from "./Post/CreatePost";
import PostFeed from "./Post/PostFeed";
import Profile from "./Profile/Profile";
import styles from "./Dashboard.module.css"
import Group from "./Group/Group";

export default function Dashboard() {
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
            <Route path="/group/:groupname" element={<Group />} />
          </Routes>
        </div>
      </div>

    </Container>
  );
}
