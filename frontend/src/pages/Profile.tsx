import React, { FC } from "react";
import { Container, Box, Heading, Text } from "@chakra-ui/react";
import { useUser } from "../state/auth";

const Profile: FC = () => {
  const user = useUser();
  console.log(user);
  return (
    <Container>
      <Box>
        <Heading as="h1">Profile</Heading>
        <Text>
          Navn: {"name"}
          <br />
          Kallenavn: {"name"}
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
