import { Navigate } from "react-router-dom";

const Signout = () => {
  localStorage.token = "";
  return <Navigate to={{ pathname: "/notes" }} />;
};

export default Signout;
