import React, { FC } from "react";
import { Container, Box, Heading } from "@chakra-ui/react";
import { useUser } from "../state/auth";
import { Navigate } from "react-router-dom";

const Profile: FC = () => {
  const user = useUser();

  if (!user) {
    return <Container>Loading...</Container>;
  }

  return (
    <Container>
      <Box>
        <Heading as="h1">Profile</Heading>
        <p>{user.profile.name}</p>
      </Box>
    </Container>
  );
};

export default Profile;
