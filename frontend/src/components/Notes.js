import { Navigate } from "react-router-dom";
import axios from "axios";

// import { useState } from "react";

const Notes = (props) => {
  var token = localStorage.token;
  return (!token || token === "") ? <Navigate to={{ pathname: '/' }} /> : <div>{token}</div>;
};

export default Notes;
