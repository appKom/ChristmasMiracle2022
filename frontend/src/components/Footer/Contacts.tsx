import React, { FC } from "react";
import { Box, VStack, Text, Link } from "@chakra-ui/react";
import { EmailIcon } from "@chakra-ui/icons";
const Contacts: FC = () => {
  return (
    <VStack>
      <Text fontWeight={700}>KONTAKT</Text>
      <Box>
        <Text>
          <Link href="mailto: appkom@online.ntnu.no">
            <EmailIcon margin="4px" />
            appkom@online.ntnu.no
          </Link>
        </Text>
      </Box>
    </VStack>
  );
};

export default Contacts;
