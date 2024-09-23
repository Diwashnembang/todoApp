import {
  Box,
  FormControl,
  Input,
  FormHelperText,
  FormLabel,
  Button,
} from "@chakra-ui/react";
import React, { useState } from "react";

const TaskInput = () => {
  const [task, setTask] = useState("");

  const handleSubmit = (e) => {
    e.preventDefault();
    const addTask = async () => {
      try {
        const respons = await fetch("http://localhost:8000/", {
          method: "POST",
          headers: { "Content-type": "application/json; charset=UTF-8" },
          body: JSON.stringify({
            task: task,
          }),
        });
        const isOK = respons;
        console.log(isOK);
      } catch (e) {
        console.log("couldn't fetch data ", e);
      }
    };
    addTask();
  };
  const handleChange = (e) => {
    setTask(e.target.value);
  };
  return (
    <Box className="input_form">
      <form onSubmit={handleSubmit}>
        <FormControl>
          <FormLabel>Add Task </FormLabel>
          <Input type="text" onChange={handleChange} />
          <FormHelperText>Add Your Task</FormHelperText>
        </FormControl>
        <Button colorScheme="blue" mt={6} type="submit">
          Submit
        </Button>
      </form>
    </Box>
  );
};

export default TaskInput;
