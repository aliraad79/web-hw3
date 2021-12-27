import { Navigate, useNavigate } from "react-router-dom";
import { useState, useEffect } from "react";
import { FaPlus } from "react-icons/fa";
import { Button } from "react-bootstrap";
import Note from "./Note";
import BASE_SERVER_URL from "../consts";

const Notes = (props) => {
  const [notes, setNotes] = useState([]);
  let navigate = useNavigate();
  const token = localStorage.token;

  useEffect(() => {
    const getNotes = async () => {
      const res = await fetch(`${BASE_SERVER_URL}/notes/`, {
        headers: {
          Authorization: token,
        },
      });
      if (res.status === 401) navigate("/");
      else {
        const data = await res.json();
        setNotes(data !== "null" ? data : []);
      }
    };
    getNotes();
  }, []);

  const onDelete = (id) => {
    console.log(id);
    fetch(`${BASE_SERVER_URL}/notes/${id}`, {
      method: "DELETE",
      headers: {
        Authorization: token,
      },
    });
    setNotes(notes.filter((note) => note.id !== id));
  };

  const updateOrAddNote = async (note) => {
    if (note.id === 0) {
      fetch(`${BASE_SERVER_URL}/notes/`, {
        method: "POST",
        headers: {
          Authorization: token,
        },
        body: JSON.stringify(note),
      })
        .then((response) => response.json())
        .then((response) =>
          setNotes([
            ...notes.filter((n) => n.id !== note.id),
            { ...note, id: response.ID },
          ])
        );
    } else {
      await fetch(`${BASE_SERVER_URL}/notes/${note.id}`, {
        method: "PUT",
        headers: {
          Authorization: token,
        },
        body: JSON.stringify(note),
      });
    }
  };

  const notesItems = notes.map((note) => {
    return (
      <Note
        key={note.id}
        note={note}
        onDelete={onDelete}
        onUpdateOrAdd={updateOrAddNote}
      />
    );
  });

  return !token || token === "" ? (
    <Navigate to={{ pathname: "/" }} />
  ) : (
    <center>
      <br />
      {notesItems}
      <Button
        onClick={(e) =>
          setNotes([...notes.filter((note) => note.id !== 0), { id: 0 }])
        }
      >
        <FaPlus />
      </Button>
    </center>
  );
};

export default Notes;
