import { Form, Button, Modal } from "react-bootstrap";
import { useState } from "react";
import {  useNavigate } from "react-router-dom";
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
          console.log("User not found");
        }
      })
      .then((response) => {
        localStorage.setItem("token", `${response.token}`);
        navigate("/notes", { replace: true });
      });
  };

  return (
    <center>
      <Form onSubmit={login_user}>
        <Form.Group className="mb-3" controlId="formBasicUsername">
          <Form.Label>Username</Form.Label>
          <Form.Control
            type="username"
            placeholder="username"
            onChange={(e) => {
              setUsername(e.target.value);
            }}
          />
        </Form.Group>

        <Form.Group className="mb-3" controlId="formBasicPassword">
          <Form.Label>Password</Form.Label>
          <Form.Control
            type="password"
            placeholder="Password"
            onChange={(e) => {
              setPassword(e.target.value);
            }}
          />
        </Form.Group>
        <Button variant="primary" type="submit">
          Login
        </Button>
      </Form>
      
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
    </center>
  );
};

export default Login;
