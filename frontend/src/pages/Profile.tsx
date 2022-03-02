import React, { FC } from "react";
import { Container, Box, Heading, Text } from "@chakra-ui/react";
import { useUser } from "../state/auth";

const Profile: FC = () => {
  const user = useUser();

  if (!user) {
    return <Container>Loading...</Container>;
  }
  console.log(user);
  return (
    <Container>
      <Box>
        <Heading as="h1">Profile</Heading>
        <Text>
          Navn: {user.profile.name}
          <br />
          Kallenavn: {user.profile.nickname}
          <br />
          Poeng: {0}
          <br />
          Flagg: {0}
        </Text>
      </Box>
    </Container>
  );
};

export default Profile;
