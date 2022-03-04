import React, { FC } from "react";
import { Text } from "@chakra-ui/react";
import PageContainer from "../components/PageContainer";

const Leaderboard: FC = () => {
  return (
    <PageContainer title="Ledertavle">
      <Text>Her finner man de deltagerne med høyest poengsum!</Text>
    </PageContainer>
  );
};

export default Leaderboard;
