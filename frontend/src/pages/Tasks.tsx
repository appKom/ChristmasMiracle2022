import React, { FC, useEffect } from "react";
import { Text } from "@chakra-ui/react";
import PageContainer from "../components/PageContainer";
import { fetchTasks } from "../api";
import { TaskType } from "../types/api";

const Tasks: FC = () => {
  const [tasks, setTasks] = React.useState<TaskType[]>([]);

  useEffect(() => {
    fetchTasks().then((tasks) => setTasks(tasks));
  }, []);

  return (
    <PageContainer title="Oppgaver">
      <Text>Her finner man oppgaver!</Text>
      {tasks.map((task) => (
        <div key={task.Title}> Title: {task.Title}</div>
      ))}
    </PageContainer>
  );
};

export default Tasks;
