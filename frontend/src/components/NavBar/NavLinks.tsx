import React, { FC } from "react";
import NavLink from "./NavLink";
import NavButton from "./NavButton";
import { Spacer, Flex } from "@chakra-ui/react";
import { useUser, logOutUser } from "../../state/auth";

enum NavPages {
  REGISTER = "/register",
  LOGIN = "/login",
  LEADERBOARD = "/leaderboard",
  TASKS = "/tasks",
  PROFILE = "/profile",
  HOME = "/",
}

const NavLinks: FC = () => {
  const user = useUser();

  const onLogout = () => {
    logOutUser();
    window.location.reload();
  };

  return (
    <Flex w="40%" marginLeft="auto" height="100%" p={3} align="center">
      <NavLink to={NavPages.HOME}>Hjem</NavLink>
      <Spacer />
      <NavLink to={NavPages.LEADERBOARD}>Ledertavle</NavLink>
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
        <>
          <NavLink asButton to={NavPages.LOGIN}>
            Logg inn
          </NavLink>
          <Spacer />
          <NavLink asButton to={NavPages.REGISTER}>
            Registrer
          </NavLink>
        </>
      )}
    </Flex>
  );
};

export default NavLinks;
