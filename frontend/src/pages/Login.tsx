import React, { FC } from "react";
import PageContainer from "../components/PageContainer";
import LoginForm from "../components/Login";
import { loginUser } from "../api/auth";

const Register: FC = () => {
  const onSubmit = (data: any) => {
    const username = data.username;
    const password = data.password;
    loginUser({ username, password }).then((resp) => console.log(resp));
  };

  return (
    <PageContainer title="Logg inn">
      <LoginForm onSubmit={onSubmit} />
    </PageContainer>
  );
};

export default Register;
