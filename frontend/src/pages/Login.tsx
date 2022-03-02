import React from "react";
import userManager from "../lib/oidc";
import { useUser } from "../state/auth";
import { Heading, Container, Button } from "@chakra-ui/react";
import { Navigate } from "react-router-dom";

const Login: React.FC = () => {
  const user = useUser();
  const signInWithRedirect = () => userManager.signinRedirect();

  if (user) {
    return <Navigate to="/" />;
  }

  return (
    <Container>
      <Heading>Login</Heading>
      <Button onClick={signInWithRedirect}>Sign in with redirect</Button>
    </Container>
  );
};

export default Login;
