import { Navigate } from "react-router-dom";
import { useState, useEffect } from "react";
import { FaPlus } from "react-icons/fa";
import { Button } from "react-bootstrap";
import Note from "./Note";

const Notes = (props) => {
  const [notes, setNotes] = useState([]);

  useEffect(() => {
    const getNotes = async () => {
      const notesFromServer = await fetchAllNotes();
      setNotes(notesFromServer);
    };
    getNotes();
  }, []);

  const fetchAllNotes = async () => {
    const res = await fetch(`http://localhost:8080/notes/`);
    const data = await res.json();
    return data;
  };

  const onDelete = (id) => {
    console.log(id);
    fetch(`http://localhost:8080/notes/${id}`, { method: "DELETE" });
    setNotes(notes.filter((note) => note.id !== id));
  };

  const updateOrAddNote = async (note) => {
    let new_id = 0;
    if (note.id === 0) {
      fetch(`http://localhost:8080/notes/new`, {
        method: "POST",
        body: JSON.stringify(note),
      })
        .then((response) => response.json())
        .then((response) => (new_id = response.ID));
    } else {
      const res = await fetch(`http://localhost:8080/notes/${note.id}`, {
        method: "PUT",
        body: JSON.stringify(note),
      });

      const data = await res.json();
      console.log(data);
    }
    setNotes([
      ...notes.filter((n) => n.id !== note.id),
      { ...note, id: new_id },
    ]);
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

  const token = localStorage.token;
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
