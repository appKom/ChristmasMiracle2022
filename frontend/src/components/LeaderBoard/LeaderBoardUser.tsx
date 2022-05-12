import React, { FC } from "react";
import { ScoreBoardUser } from "../../types/api";
import { Tr, Td } from "@chakra-ui/react";

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
      <Td>{user.username}</Td>
      <Td>{user.points}</Td>
    </Tr>
  );
};
export default LeaderBoardUser;
