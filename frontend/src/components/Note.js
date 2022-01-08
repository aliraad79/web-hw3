import { Button, Form, Container, Row, Col } from "react-bootstrap";
import { useState } from "react";
import { FaTrash, FaPencilAlt, FaCheckCircle } from "react-icons/fa";

const Note = ({ note, onDelete, onUpdateOrAdd }) => {
  let empty_note = note.ID === 0;
  const [title, setTitle] = useState(note.title);
  const [body, setBody] = useState(note.body);
  const [editMode, setEditMode] = useState(!empty_note);
  const [accept, setAccept] = useState(empty_note);

  return (
    <Container className="" style={{
      backgroundImage: "linear-gradient(to left, #616161, #9bc5c3)",
    }}>
      <br />
      <Form>
        <Row>
          <Col xs={10}>
            <Form.Control
              type="text"
              placeholder="Subject"
              readOnly={editMode}
              value={title}
              onChange={(e) => {
                setTitle(e.target.value);
              }}
            />
          </Col>
          <Col>
            <Button onClick={(e) => onDelete(note.ID)}>
              <FaTrash />
            </Button>
            <Button
              onClick={(e) => {
                setEditMode(!editMode);
                setAccept(editMode);
              }}
            >
              <FaPencilAlt />
            </Button>
            {accept && (
              <Button
                onClick={(e) => {
                  setEditMode(!editMode);
                  setAccept(editMode);
                  onUpdateOrAdd({ ID: note.ID, Title: title, Body: body });
                }}
              >
                <FaCheckCircle />
              </Button>
            )}
          </Col>
        </Row>
        <br />
        <Row>
          <Col>
            <Form.Control
              as="textarea"
              placeholder="Text"
              rows={3}
              readOnly={editMode}
              defaultValue={body}
              onChange={(e) => {
                setBody(e.target.value);
              }}
            />
          </Col>
        </Row>
        <br />
      </Form>
    </Container>
  );
};

Note.defaultProps = {
  readOnly: true,
};

export default Note;
