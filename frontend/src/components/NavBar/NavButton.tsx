import React, { FC } from "react";
import { Button } from "@chakra-ui/react";

type NavButtonProps = {
  children: React.ReactNode;
  func: () => void;
};

const NavButton: FC<NavButtonProps> = ({ children, func }: NavButtonProps) => {
  return (
    <Button onClick={func} variant="outline" color="white" mr={4}>
      {children}
    </Button>
  );
};

export default NavButton;
