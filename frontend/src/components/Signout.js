import { Navigate } from "react-router-dom";

const Signout = ({ setAuthToken }) => {
  setAuthToken("");
  return <Navigate to={{ pathname: "/" }} />;
};

export default Signout;
