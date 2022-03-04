import React, { FC } from "react";
import {
  Heading,
  Text,
  Divider,
  OrderedList,
  ListItem,
} from "@chakra-ui/react";
import PageContainer from "../components/PageContainer";

const Home: FC = () => {
  return (
    <PageContainer title="Christmas Miracle 2022">
      <Text w="50%" ml="auto" mr="auto">
        Velkommen til Online sin julekalender 2022. Julekalenderen ineholder
        mange forskjellige oppgaver for alle ferdigheter. Julekalenderen er
        utviklet av Appkom, og ved oppdagelse av feil eller mangler kan du
        kontakte Martin Skatvedt på slack, eller opprette et issue i github
        repoet.
      </Text>

      <Divider width="40%" ml="auto" mr="auto" mt={5} mb={5} />
      <Heading as="h2" size="lg" mb={4} mt={4}>
        Regler
      </Heading>
      <OrderedList spacing={2} listStylePos="inside">
        <ListItem>
          Ikke bruk automatiske verktøy mot noen av nettsidene som er satt opp
        </ListItem>
        <ListItem>
          Ondsinnet oppførsel mot julekalenderen eller andre deltagere er
          strengt forbudt
        </ListItem>
      </OrderedList>
      <Text w="50%" ml="auto" mr="auto" mb={4} mt={4}>
        {
          "Brudd på disse reglene kan føre til utestengelse fra julekalenderen. Vær snille mot hverandre og julekalenderen <3"
        }
      </Text>
    </PageContainer>
  );
};

export default Home;
