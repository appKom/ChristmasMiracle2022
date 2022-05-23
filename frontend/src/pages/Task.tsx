import React, { FC, useState, useEffect, useRef } from "react";
import { Box, Button, Text } from "@chakra-ui/react";
import { useParams } from "react-router-dom";
import { fetchTask } from "../api";
import { TaskType } from "../types/api";
import PageContainer from "../components/PageContainer";

import SubmitForm from "../components/Task/SubmitForm";
const Task: FC = () => {
  const [task, setTask] = useState<TaskType | null>(null);
  const { id } = useParams();
  const modalRef = useRef<HTMLButtonElement | null>(null);
  useEffect(() => {
    fetchTask(Number(id)).then((task) => setTask(task));
  }, [id]);

  if (!task) {
    return null;
  }

  return (
    <PageContainer title={task.title}>
      <Text>{task.content}</Text>
      <Button ref={modalRef} colorScheme="teal" mr={3}>
        Send inn Flagg
      </Button>
      <SubmitForm id={task.ID} modalRef={modalRef} />
    </PageContainer>
  );
};
export default Task;
