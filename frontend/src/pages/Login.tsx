import React, { FC } from "react";
import PageContainer from "../components/PageContainer";
import LoginForm from "../components/Login";

const Register: FC = () => {
  const onSubmit = (data: any) => {
    console.log(data);
  };

  return (
    <PageContainer title="Logg inn">
      <LoginForm onSubmit={onSubmit} />
    </PageContainer>
  );
};

export default Register;
