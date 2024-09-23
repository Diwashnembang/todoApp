import React, { useState } from "react";

import {
  Button,
  FormControl,
  FormLabel,
  HStack,
  Input,
} from "@chakra-ui/react";

const ImageUploder = () => {
  const [file, setFile] = useState("");
  const handleChange = (e) => {
    setFile(e.target.files[0]);
    const formData = new FormData();
    formData.append("file", file);

    // You would send the formData to your server here using fetch or axios
    console.log("File submitted:", formData);
  };

  const handlePost = (e) => {
    // e.preventDefault()
    // const response = fetch("http://localhost:8000/addfile",{
    //     method:"POST",
    //     body:JSON.stringify{
    //     }
    // })
  };

  return (
    <HStack>
      <form onChange={handlePost}></form>
      <FormControl>
        <FormLabel>upload file</FormLabel>
        <Input type="file" onChange={handleChange}></Input>
        <Button colorScheme="blue" mt={6} type="submit">
          Submit
        </Button>
      </FormControl>
    </HStack>
  );
};

export default ImageUploder;
