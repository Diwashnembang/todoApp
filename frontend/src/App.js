import { Route, Routes } from "react-router-dom";
import "./App.css";
import Home from "./component/Home";
import ImageUploder from "./component/ImageUploder";

function App() {
  return (
    <Routes>
      <Route path="/" element={<Home></Home>} />
      <Route path="/addfile" element={<ImageUploder></ImageUploder>} />
    </Routes>
  );
}

export default App;
