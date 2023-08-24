import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { UserContextProvider } from "./context/AuthContext";
import SignIn from "./components/Auth/SignIn";
import SignUp from "./components/Auth/SignUp";
import "./App.css";
import Dashboard from "./components/Dashboard";
import ProtectedRoute from "./util/ProtectedRoute";

function App() {
  return (
    <UserContextProvider>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<SignIn />} />
          <Route path="/signin" element={<SignIn />} />
          <Route path="/signup" element={<SignUp />} />
          <Route
            path="/dashboard/*"
            element={<ProtectedRoute element={<Dashboard />} />}
          />
        </Routes>
      </BrowserRouter>
    </UserContextProvider>
  );
}

export default App;
