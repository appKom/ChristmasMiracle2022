import React, { FC } from "react";
import {
  Box,
  Heading,
  Text,
  Divider,
  OrderedList,
  ListItem,
} from "@chakra-ui/react";

const Home: FC = () => {
  return (
    <Box w="80%" marginLeft="auto" marginRight="auto" textAlign="center">
      <Heading as="h1">Christmas Miracle 2022</Heading>
      <Text>This is a placeholder for the Christmas Miracle 2022 website.</Text>

      <Divider width="40%" ml="auto" mr="auto" mt={5} mb={5} />
      <Heading as="h2" size="lg">
        Regler
      </Heading>
      <OrderedList spacing={2} listStylePos="inside">
        <ListItem>Regel 1</ListItem>
        <ListItem>Regel 2</ListItem>
      </OrderedList>
    </Box>
  );
};

export default Home;
