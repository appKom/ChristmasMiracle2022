import React, { FC } from "react";
import { Container, Box, Heading, Text } from "@chakra-ui/react";

const Profile: FC = () => {
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
