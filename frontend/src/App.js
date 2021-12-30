import "./App.css";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";

import Login from "./components/Login";
import Notes from "./components/Notes";
import Signout from "./components/Signout"
import 'bootstrap/dist/css/bootstrap.min.css';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/notes" element={<Notes />} />
        <Route path="/signout" element={<Signout />}/>
      </Routes>
    </Router>
  );
}

export default App;
