import React, { FC } from "react";
import { Text, Link } from "@chakra-ui/react";
import { ExternalLinkIcon } from "@chakra-ui/icons";
import { Link as RouterLink } from "react-router-dom";
type FooterLinkProps = {
  name: string;
  href: string;
  isExternal: boolean;
};

const FooterLink: FC<FooterLinkProps> = ({
  name,
  href,
  isExternal,
}: FooterLinkProps) => {
  return isExternal ? (
    <Link href={href} target={isExternal ? "_blank" : ""}>
      <Text fontSize="1xl" display="block" textDecoration="none">
        {name}
        {isExternal && <ExternalLinkIcon mx="5px" />}
      </Text>
    </Link>
  ) : (
    <Link as={RouterLink} to={href} target={isExternal ? "_blank" : ""}>
      <Text fontSize="1xl" display="block" textDecoration="none">
        {name}
        {isExternal && <ExternalLinkIcon mx="5px" />}
      </Text>
    </Link>
  );
};

export default FooterLink;
