import React, { FC } from "react";
import { Box } from "@chakra-ui/react";
import { TaskType } from "../../types/api";
import TaskComponent from "./TaskComponent";

type TaskListProps = {
  tasks: TaskType[];
};

const TaskList: FC<TaskListProps> = ({ tasks }: TaskListProps) => {
  return (
    <Box>
      {tasks.map((task) => (
        <TaskComponent key={task.ID} task={task} />
      ))}
    </Box>
  );
};

export default TaskList;
