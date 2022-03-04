import React, { FC } from "react";
import { VStack, Text } from "@chakra-ui/react";
import FooterLink from "./FooterLink";

const Nav: FC = () => {
  return (
    <VStack>
      <Text fontWeight={700}>SIDEMENY</Text>
      <FooterLink name="Ledertavle" href="/leaderboard" isExternal={false} />
      <FooterLink name="Oppgaver" href="/tasks" isExternal={false} />
    </VStack>
  );
};

export default Nav;
