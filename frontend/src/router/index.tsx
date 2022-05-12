import React, { FC } from "react";
import { Routes, Route } from "react-router-dom";
import Home from "../pages/Home";
import Profile from "../pages/Profile";
import Tasks from "../pages/Tasks";
import Leaderboard from "../pages/Leaderboard";
import Register from "../pages/Register";
import Login from "../pages/Login";
import Task from "../pages/Task";

const Router: FC = () => {
  return (
    <Routes>
      <Route path="/profile" element={<Profile />} />
      <Route path="/tasks" element={<Tasks />} />
      <Route path="/task/:id" element={<Task />} />
      <Route path="/leaderboard" element={<Leaderboard />} />
      <Route path="/register" element={<Register />} />
      <Route path="/login" element={<Login />} />
      <Route path="/" element={<Home />} />
    </Routes>
  );
};

export default Router;
