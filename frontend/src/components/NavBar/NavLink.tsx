import React, { FC } from "react";
import { Link } from "@chakra-ui/react";
import { Link as RouterLink } from "react-router-dom";

type NavLinkProps = {
  to: string;
  children: React.ReactNode;
};

const NavLink: FC<NavLinkProps> = ({ to, children }: NavLinkProps) => {
  return (
    <Link as={RouterLink} to={to} color="white">
      {children}
    </Link>
  );
};

export default NavLink;
