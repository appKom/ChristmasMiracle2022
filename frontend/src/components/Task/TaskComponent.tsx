import React, { FC } from "react";
import { Box, Flex, Link, Text } from "@chakra-ui/react";
import { TaskType } from "../../types/api";
import { Link as ReachLink } from "react-router-dom";
type TaskComponentProps = {
  task: TaskType;
};

const TaskComponent: FC<TaskComponentProps> = ({
  task,
}: TaskComponentProps) => {
  return (
    <Box
      border="1px solid red"
      w="65%"
      margin="5px auto 5px auto"
      _hover={{ bg: "red" }}
    >
      <Link
        as={ReachLink}
        to={`/task/${task.ID}`}
        _hover={{ textDecoration: "none" }}
      >
        <Flex>
          <Box textAlign="left" w="70%" p="1%">
            <Text fontSize="xl">{task.title}</Text>
          </Box>
          <Box textAlign="right" w="30%" p="1%">
            <Text fontSize="xl">{task.points} poeng</Text>
          </Box>
        </Flex>
      </Link>
    </Box>
  );
};
export default TaskComponent;
