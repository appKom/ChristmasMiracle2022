import React, { FC } from "react";
import PageContainer from "../components/PageContainer";
import RegisterForm from "../components/Register";

const Register: FC = () => {
  const onSubmit = (data: any) => {
    console.log(data);
  };

  return (
    <PageContainer title="Registrer">
      <RegisterForm onSubmit={onSubmit} />
    </PageContainer>
  );
};

export default Register;
