import React, { FC } from "react";
import { VStack, Text } from "@chakra-ui/react";
import FooterLink from "./FooterLink";

const Links: FC = () => {
  return (
    <VStack>
      <Text fontWeight={700}>LENKER</Text>

      <FooterLink
        name="Github"
        href="https://github.com/appKom/ChristmasMiracle2022"
        isExternal
      />
    </VStack>
  );
};

export default Links;
