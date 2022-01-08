import { Navigate, useNavigate } from "react-router-dom";
import { useState, useEffect } from "react";
import { FaPlus } from "react-icons/fa";
import { Button } from "react-bootstrap";
import Note from "./Note";
import BASE_SERVER_URL from "../consts";
import MyNavbar from "./Navbar";

const Notes = (props) => {
  const [notes, setNotes] = useState([]);
  let navigate = useNavigate();
  const token = localStorage.token;

  useEffect(() => {
    const getNotes = async () => {
      await fetch(`${BASE_SERVER_URL}/notes/`, {
        headers: {
          Authorization: token,
        },
      })
        .then((response) => {
          if (response.status === 401) navigate("/");
          if (response.status === 200) return response.json();
        })
        .then((response) => {
          setNotes(response !== null ? response : []);
        });
    };
    getNotes();
  }, []);

  const onDelete = (id) => {
    fetch(`${BASE_SERVER_URL}/notes/${id}`, {
      method: "DELETE",
      headers: {
        Authorization: token,
      },
    });
    setNotes(notes.filter((note) => note.ID !== id));
  };

  const updateOrAddNote = async (note) => {
    if (note.ID === 0) {
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
            ...notes.filter((n) => n.ID !== note.ID),
            { ...note, ID: response.ID },
          ])
        );
    } else {
      await fetch(`${BASE_SERVER_URL}/notes/${note.ID}`, {
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
        key={note.ID}
        note={note}
        onDelete={onDelete}
        onUpdateOrAdd={updateOrAddNote}
      />
    );
  });

  console.log(notes);

  return !token || token === "" ? (
    <Navigate to={{ pathname: "/" }} />
  ) : (
    <>
      <MyNavbar />
      <center>
        <br />
        {notesItems}
        <Button
          onClick={(e) =>
            setNotes([...notes.filter((note) => note.ID !== 0), { ID: 0 }])
          }
        >
          <FaPlus />
        </Button>
      </center>
    </>
  );
};

export default Notes;
