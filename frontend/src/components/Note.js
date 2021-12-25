import { Button, Form, Container, Row, Col } from "react-bootstrap";
import { useState } from "react";
import { FaTrash, FaPencilAlt } from "react-icons/fa";

const Note = ({ subject, text, readOnly }) => {
  const [sub, setSubject] = useState(subject);
  const [_text, setText] = useState(text);
  const [editMode, setEditMode] = useState(readOnly);

  return (
    <Container className="bg-info">
      <br />
      <Form>
        <Row>
          <Col xs={10}>
            <Form.Control
              type="text"
              placeholder="Subject"
              readOnly={editMode}
              value={sub}
              onChange={(e) => {
                setSubject(e.target.value);
              }}
            />
          </Col>
          <Col>
            <Button>
              <FaTrash />
            </Button>
            <Button onClick={(e) => setEditMode(!editMode)}>
              <FaPencilAlt />
            </Button>
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
              defaultValue={_text}
              onChange={(e) => {
                setText(e.target.value);
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
