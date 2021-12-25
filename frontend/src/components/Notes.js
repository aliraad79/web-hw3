import { Navigate } from "react-router-dom";
import { useState } from "react";
import { FaPlus } from "react-icons/fa";
import { Button } from "react-bootstrap";
import Note from "./Note";

const Notes = (props) => {
  const [notes, setNotes] = useState([]);

  const token = localStorage.token;
  return !token || token === "" ? (
    <Navigate to={{ pathname: "/" }} />
  ) : (
    <center>
      <br />
      <Note subject="TH1" text="text1" />
      <Note subject="TH2" text="textw" />
      <Note readOnly={false} />
      <Button>
        <FaPlus />
      </Button>
    </center>
  );
};

export default Notes;
