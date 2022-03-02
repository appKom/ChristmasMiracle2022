import React, { FC } from "react";
import { Box } from "@chakra-ui/react";
import { useUser } from "../state/auth";

const Home: FC = () => {
  const user = useUser();
  //console.log(user)
  return (
    <Box>
      <h1>Home</h1>
    </Box>
  );
};

export default Home;
