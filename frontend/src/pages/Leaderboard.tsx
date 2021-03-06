import React, { FC, useEffect, useState } from "react";
import { Text } from "@chakra-ui/react";
import PageContainer from "../components/PageContainer";
import { fetchScoreBoard } from "../api";
import { useRecoilState } from "recoil";
import { tokenState } from "../state/auth";
import { ScoreBoardUser } from "../types/api";
import LeaderBoard from "../components/LeaderBoard";

const Leaderboard: FC = () => {
  const [token] = useRecoilState(tokenState);
  const [scoreBoard, setScoreBoard] = useState<ScoreBoardUser[]>([]);

  useEffect(() => {
    fetchScoreBoard(token).then(setScoreBoard);
  }, [token]);

  return (
    <PageContainer title="Ledertavle">
      <Text>Her finner man de deltagerne med høyest poengsum!</Text>
      <LeaderBoard users={scoreBoard} />
    </PageContainer>
  );
};

export default Leaderboard;
