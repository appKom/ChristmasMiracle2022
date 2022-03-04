import React, { FC } from "react";
import { Link, Button } from "@chakra-ui/react";
import { Link as RouterLink } from "react-router-dom";

type NavLinkProps = {
  to: string;
  children: React.ReactNode;
  asButton?: boolean;
};

const NavLink: FC<NavLinkProps> = ({
  to,
  children,
  asButton,
}: NavLinkProps) => {
  return asButton ? (
    <Button as={RouterLink} to={to} variant="outline" color="white" mr={4}>
      {children}
    </Button>
  ) : (
    <Link as={RouterLink} to={to} color="white" size="lg">
      {children}
    </Link>
  );
};

export default NavLink;
