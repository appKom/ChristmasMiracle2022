import React, { FC } from "react";
import PageContainer from "../components/PageContainer";
import LoginForm from "../components/Login";
import { useRecoilState } from "recoil";
import { tokenState } from "../state/auth";
import { Navigate } from "react-router-dom";

const Login: FC = () => {
  const [token] = useRecoilState(tokenState);
  if (token) {
    return <Navigate to="/" />;
  }

  return (
    <PageContainer title="Logg inn">
      <LoginForm />
    </PageContainer>
  );
};

export default Login;
