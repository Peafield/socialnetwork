import React from "react";
import { Router, Routes, Route } from "react-router-dom";
import { AuthContext } from "./context/AuthContext";
import { useAuth } from "./hooks/useAuth";
import SignIn from "./components/SignIn";
import SignUp from "./components/Auth/SignUp";
import "./App.css";

function App() {
  const { user, setUser } = useAuth();

  return (
    <AuthContext.Provider value={{ user, setUser }}>
      <Router>
        <Routes>
          <Route path="/" element={<SignIn />} />
          <Route path="/signin" element={<SignIn />} />
          <Route path="/signup" element={<SignUp />} />
          {/* <Route
            path="/dashboard/*"
            element={<ProtectedRoute element={<Dashboard />} />}
          /> */}
        </Routes>
      </Router>
    </AuthContext.Provider>
  );
}

export default App;
