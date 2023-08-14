import React, { useEffect, useState } from "react";
import { Route, Routes } from "react-router-dom";
import Container from "./Containers/Container";
import CreatePost from "./Post/CreatePost";
import PostFeed from "./Post/PostFeed";

export default function Dashboard() {




  return (
    <Container>

      <Routes>
        <Route path="/" element={<PostFeed />} />
        <Route path="/createpost" element={<CreatePost />} />
      </Routes>

    </Container>
  );
}
