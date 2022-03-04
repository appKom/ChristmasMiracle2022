import React, { FC } from "react";
import { Text } from "@chakra-ui/react";
import PageContainer from "../components/PageContainer";

const Tasks: FC = () => {
  return (
    <PageContainer title="Oppgaver">
      <Text>Her finner man oppgaver!</Text>
    </PageContainer>
  );
};

export default Tasks;
