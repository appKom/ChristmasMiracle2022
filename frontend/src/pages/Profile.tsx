import React, { FC } from "react";
import { Container, Box, Heading, Text } from "@chakra-ui/react";
import { useRecoilState } from "recoil";
import { tokenState } from "../state/auth";

const Profile: FC = () => {
  const [token] = useRecoilState(tokenState);
  return (
    <Container>
      <Box>
        <Heading as="h1">Profile</Heading>
        <Text>
          Navn: {token ? token.Access : "Ingen token"}
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
