import React, { FC, useEffect, useState } from "react";
import { Container, Box, Heading, Text } from "@chakra-ui/react";
import { useRecoilState } from "recoil";
import { tokenState } from "../state/auth";
import { fetchProfile } from "../api";
import { User } from "../types/api";

const Profile: FC = () => {
  const [token] = useRecoilState(tokenState);
  const [user, setUser] = useState<User | null>(null);

  useEffect(() => {
    fetchProfile(token).then(setUser);
  }, [token]);

  if (!user) {
    return <Box> Loading...</Box>;
  }

  return (
    <Container>
      <Box>
        <Heading as="h1">Profile</Heading>
        <Text>
          Navn: {user.Username}
          <br />
          Epost: {user.Email}
          <br />
          Poeng: {user.Points}
          <br />
          Admin: {user.Admin ? "Ja" : "Nei"}
        </Text>
      </Box>
    </Container>
  );
};

export default Profile;
