import React, { FC } from "react";
import { Box, Stack, Text } from "@chakra-ui/react";
import Contacts from "./Contacts";
import Links from "./Links";
import Nav from "./Nav";
const Footer: FC = () => {
  return (
    <Box w="100%" bg="#830e10" height="10vh">
      <Stack
        direction={["column", "row", "row", "row"]}
        spacing="10%"
        height="70%"
        justify="center"
        align="center"
      >
        <Contacts />
        <Nav />
        <Links />
      </Stack>
    </Box>
  );
};

export default Footer;
