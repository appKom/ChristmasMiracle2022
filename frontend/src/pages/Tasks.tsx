import React, { FC, useEffect } from "react";
import { Text } from "@chakra-ui/react";
import PageContainer from "../components/PageContainer";
import { fetchTasks } from "../api";
import { TaskType } from "../types/api";
import TaskList from "../components/Task";

const Tasks: FC = () => {
  const [tasks, setTasks] = React.useState<TaskType[]>([]);

  useEffect(() => {
    fetchTasks().then((tasks) => setTasks(tasks));
  }, []);

  return (
    <PageContainer title="Oppgaver">
      <Text>Her finner man oppgaver!</Text>
      <TaskList tasks={tasks} />
    </PageContainer>
  );
};

export default Tasks;
