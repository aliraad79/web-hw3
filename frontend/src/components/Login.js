import {
  Form,
  Button,
  Modal,
  FloatingLabel,
  Container,
  Row,
  Col,
} from "react-bootstrap";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import BASE_SERVER_URL from "../consts";

const Login = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [showModalError, setShowModalError] = useState(false);

  const handleClose = () => setShowModalError(false);
  const navigate = useNavigate();

  const login_user = async (event) => {
    event.preventDefault();
    await fetch(`${BASE_SERVER_URL}/login`, {
      method: "POST",
      body: JSON.stringify({
        username: username,
        password: password,
      }),
    })
      .then((response) => {
        if (response.status !== 401) {
          return response.json();
        } else {
          setShowModalError(true);
        }
      })
      .then((response) => {
        localStorage.setItem("token", `${response.token}`);
        navigate("/notes", { replace: true });
      });
  };

  return (
    <Container>
      <Row>
        <Col></Col>
        <Col>
          <Form onSubmit={login_user}>
            <FloatingLabel label="Username">
              <Form.Control
                type="username"
                placeholder="username"
                onChange={(e) => {
                  setUsername(e.target.value);
                }}
              />
            </FloatingLabel>

            <FloatingLabel label="Password">
              <Form.Control
                type="password"
                placeholder="Password"
                onChange={(e) => {
                  setPassword(e.target.value);
                }}
              />
            </FloatingLabel>
            <center>
              <Button variant="primary" type="submit">
                Login
              </Button>
            </center>
          </Form>
        </Col>
        <Col></Col>
      </Row>
      <Modal show={showModalError} onHide={handleClose}>
        <Modal.Header closeButton>
          <Modal.Title>Login Error</Modal.Title>
        </Modal.Header>
        <Modal.Body>Username or password is incorrect</Modal.Body>
        <Modal.Footer>
          <Button variant="primary" onClick={handleClose}>
            OK
          </Button>
        </Modal.Footer>
      </Modal>
    </Container>
  );
};

export default Login;
