import { Box, HStack } from "@chakra-ui/react";
import React from "react";
import TaskInput from "./TaskInput";

const Home = () => {
  return (
    <HStack>
      <Box className="tasks"></Box>
      <Box className="addTask">
        <TaskInput></TaskInput>
      </Box>
    </HStack>
  );
};

export default Home;
