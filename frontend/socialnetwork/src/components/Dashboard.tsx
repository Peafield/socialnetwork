import React, { useEffect, useState } from "react";
import { Route, Routes } from "react-router-dom";
import Container from "./Containers/Container";
import CreatePost from "./Post/CreatePost";
import PostFeed from "./Post/PostFeed";
import Profile from "./Profile/Profile";

export default function Dashboard() {
  return (
    <Container>
      <Routes>
        <Route path="/" element={<PostFeed />} />
        <Route path="/createpost" element={<CreatePost />} />
        <Route path="/user/:username" element={<Profile />} />
      </Routes>
    </Container>
  );
}
