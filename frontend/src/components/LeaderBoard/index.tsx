import React, { FC } from "react";
import { TableContainer, Thead, Table, Tr, Th, Tbody } from "@chakra-ui/react";
import { ScoreBoardUser } from "../../types/api";
import LeaderBoardUser from "./LeaderBoardUser";

type LeaderBoardProps = {
  users: ScoreBoardUser[];
};

const LeaderBoard: FC<LeaderBoardProps> = ({ users }: LeaderBoardProps) => {
  return (
    <TableContainer w="75%" ml="auto" mr="auto">
      <Table variant="simple">
        <Thead color="red">
          <Tr>
            <Th>Plassering</Th>
            <Th>Brukernavn</Th>
            <Th>Poeng</Th>
          </Tr>
        </Thead>
        <Tbody>
          {users.map((user, index) => (
            <LeaderBoardUser placement={index + 1} user={user} />
          ))}
        </Tbody>
      </Table>
    </TableContainer>
  );
};
export default LeaderBoard;
