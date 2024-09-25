import { Route, Routes } from "react-router-dom";
import "./App.css";
import Home from "./component/Home";
import ImageUploder from "./component/ImageUploder";
import Login from "./component/Login";

function App() {
  return (
    <Routes>
      <Route path="/" element={<Home></Home>} />
      <Route path="/addfile" element={<ImageUploder></ImageUploder>} />
      <Route path="/Login" element={<Login/>} />
    </Routes>
  );
}

export default App;
