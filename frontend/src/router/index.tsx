import React, { FC } from "react";
import { Routes, Route } from "react-router-dom";
import Home from "../pages/Home";
import Profile from "../pages/Profile";
import RequireAuth from "./RequireAuth";
import Tasks from "../pages/Tasks";
import Leaderboard from "../pages/Leaderboard";
import Register from "../pages/Register";
import Login from "../pages/Login";

const Router: FC = () => {
  return (
    <Routes>
      <Route
        path="/profile"
        element={
          <RequireAuth>
            <Profile />
          </RequireAuth>
        }
      />
      <Route path="/tasks" element={<Tasks />} />
      <Route path="/leaderboard" element={<Leaderboard />} />
      <Route path="register" element={<Register />} />
      <Route path="login" element={<Login />} />
      <Route path="/" element={<Home />} />
    </Routes>
  );
};

export default Router;
