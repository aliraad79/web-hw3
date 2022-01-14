import "./App.css";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import { useCookies } from "react-cookie";

import Login from "./components/Login";
import Notes from "./components/Notes";
import Signout from "./components/Signout";
import SignUp from "./components/SignUp";
import { SESSION_AGE_IN_HOURS } from "./consts";
import "bootstrap/dist/css/bootstrap.min.css";

function App() {
  const [cookies, setCookie] = useCookies(["authToken"]);
  const current = new Date();
  const expireDate = new Date();
  expireDate.setHours(current.getHours() + SESSION_AGE_IN_HOURS);

  const setAuthToken = (token) => {
    setCookie("authToken", `${token}`, { expires: expireDate });
  };

  const getAuthToken = () => {
    return cookies.authToken;
  };

  return (
    <Router>
      <Routes>
        <Route
          path="/"
          element={
            <Login setAuthToken={setAuthToken} getAuthToken={getAuthToken} />
          }
        />
        <Route
          path="/notes"
          element={
            <Notes setAuthToken={setAuthToken} getAuthToken={getAuthToken} />
          }
        />
        <Route
          path="/signup"
          element={<SignUp getAuthToken={getAuthToken} />}
        />
        <Route
          path="/signout"
          element={<Signout setAuthToken={setAuthToken} />}
        />
      </Routes>
    </Router>
  );
}

export default App;
