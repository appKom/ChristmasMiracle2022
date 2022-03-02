import React from "react";
import { Navigate } from "react-router-dom";
import { useUser } from "../state/auth";

const RequireAuth = ({ children }: { children: JSX.Element }) => {
  const user = useUser();

  console.log(user);
  if (!user) {
    console.log("user is not logged in");
    return <Navigate to="/" replace />;
  }
  return children;
};

export default RequireAuth;
