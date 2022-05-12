import React, { FC, useEffect, useState } from "react";
import {
  Box,
  TableContainer,
  Table,
  Thead,
  Tr,
  Td,
  Tbody,
} from "@chakra-ui/react";
import { useRecoilState } from "recoil";
import { tokenState } from "../state/auth";
import { fetchProfile } from "../api";
import { User } from "../types/api";
import PageContainer from "../components/PageContainer";

const Profile: FC = () => {
  const [token] = useRecoilState(tokenState);
  const [user, setUser] = useState<User | null>(null);

  useEffect(() => {
    fetchProfile(token).then(setUser);
  }, [token]);

  if (!user) {
    return <Box> Loading...</Box>;
  }

  return (
    <PageContainer title="Profil">
      <Box>
        <TableContainer w="75%" ml="auto" mr="auto">
          <Table variant="simple">
            <Thead>
              <Tr>
                <Td>Brukernavn</Td>
                <Td>Email</Td>
                <Td>Poeng</Td>
                {user.admin ? <Td>Admin</Td> : null}
              </Tr>
            </Thead>
            <Tbody>
              <Tr>
                <Td>{user.username}</Td>
                <Td>{user.email}</Td>
                <Td>{user.points}</Td>
                {user.admin ? <Td>Ja</Td> : null}
              </Tr>
            </Tbody>
          </Table>
        </TableContainer>
      </Box>
    </PageContainer>
  );
};

export default Profile;
