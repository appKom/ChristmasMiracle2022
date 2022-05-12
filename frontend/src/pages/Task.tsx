import React, { FC, useState, useEffect } from "react";
import { Box, Heading, Text } from "@chakra-ui/react";
import { useParams } from "react-router-dom";
import { fetchTask } from "../api";
import { TaskType } from "../types/api";
import PageContainer from "../components/PageContainer";

const Task: FC = () => {
  const [task, setTask] = useState<TaskType | null>(null);
  const { id } = useParams();

  useEffect(() => {
    fetchTask(Number(id)).then((task) => setTask(task));
  }, [id]);

  if (!task) {
    return null;
  }

  return (
    <PageContainer title={task.title}>
      <Text>{task.content}</Text>
    </PageContainer>
  );
};
export default Task;
