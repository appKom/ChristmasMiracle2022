import React, { FC } from "react";
import { ScoreBoardUser } from "../../types/api";
import { Box, Tr, Td } from "@chakra-ui/react";

type LeaderBoardUserProps = {
  user: ScoreBoardUser;
  placement: number;
};

const LeaderBoardUser: FC<LeaderBoardUserProps> = ({
  placement,
  user,
}: LeaderBoardUserProps) => {
  return (
    <Tr>
      <Td>{placement}</Td>
      <Td>{user.Username}</Td>
      <Td>{user.Points}</Td>
    </Tr>
  );
};
export default LeaderBoardUser;
