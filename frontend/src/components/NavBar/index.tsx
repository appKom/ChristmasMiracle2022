import React, { FC } from "react";
import { Box } from "@chakra-ui/react";

import NavLinks from "./NavLinks";

const NavBar: FC = () => {
  return (
    <Box w="100%" bg="black" h="50px">
      <NavLinks />
    </Box>
  );
};

export default NavBar;
