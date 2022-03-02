import React, { FC } from "react";
import NavLink from "./NavLink";
import NavButton from "./NavButton";
import { Spacer, Flex } from "@chakra-ui/react";
import { useUser, logOutUser } from "../../state/auth";
import userManager from "../../lib/oidc";

enum NavPages {
  LEADERBOARD = "/leaderboard",
  TASKS = "/tasks",
  PROFILE = "/profile",
  HOME = "/",
}

const NavLinks: FC = () => {
  const user = useUser();

  const signInWithRedirect = () => {
    userManager.signinRedirect();
  };

  const onLogout = () => {
    logOutUser();
    window.location.reload();
  };

  return (
    <Flex w="30%" h="50px" marginLeft="auto" padding="2px">
      <NavLink to={NavPages.HOME}>Hjem</NavLink>
      <Spacer />
      <NavLink to={NavPages.LEADERBOARD}>Leaderboard</NavLink>
      <Spacer />
      <NavLink to={NavPages.TASKS}>Oppgaver</NavLink>
      <Spacer />
      {user ? (
        <>
          <NavLink to={NavPages.PROFILE}>Profil</NavLink>
          <Spacer />
          <NavButton func={onLogout}>Logg ut</NavButton>
        </>
      ) : (
        <NavButton func={signInWithRedirect}>Logg inn</NavButton>
      )}
    </Flex>
  );
};

export default NavLinks;
