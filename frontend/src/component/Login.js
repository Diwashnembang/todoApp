import {
  Button,
  FormControl,
  FormLabel,
  HStack,
  Input,
} from "@chakra-ui/react";
import React, { useState } from "react";

const Login = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const handleLogin = async (e) => {
    e.preventDefault();
    let formData = new FormData();
    formData.append("username", username);
    formData.append("passoword", password);
    console.log(formData.get("username"));
    try {
      const respons = await fetch("http://localhost:8000/login", {
        method: "POST",
        body: formData,
      });
    } catch (e) {
      console.log("error in posting login details", e);
    }
  };
  return (
    <HStack>
      <form onSubmit={handleLogin}>
        <FormControl>
          <FormLabel>username</FormLabel>
          <Input onChange={(e) => setUsername(e.target.value)}></Input>
        </FormControl>
        <FormControl>
          <FormLabel>passwordk</FormLabel>
          <Input
            type="password"
            onChange={(e) => setPassword(e.target.value)}
          ></Input>
        </FormControl>
        <Button colorScheme="blue" mt={6} type="submit">
          Log in
        </Button>
      </form>
    </HStack>
  );
};

export default Login;
