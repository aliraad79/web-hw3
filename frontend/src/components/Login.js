import { Form, Button } from "react-bootstrap";
import { useState } from "react";

const Login = () => {
  const [token, setToken] = useState("");

  const login = async (cred) => {
    const res = await fetch("http://localhost:8080/login", {
      method: "POST",
      headers: { "Content-type": "application/json" },
      body: JSON.stringify({
        username: cred.target[0].value,
        password: cred.target[1].value,
      }),
    });
    const data = await res.json();
    alert(data);
    // setToken(data);
  };

  return (
    <center>
      <Form onSubmit={login}>
        <Form.Group className="mb-3" controlId="formBasicUsername">
          <Form.Label>Username</Form.Label>
          <Form.Control type="username" placeholder="username" />
        </Form.Group>

        <Form.Group className="mb-3" controlId="formBasicPassword">
          <Form.Label>Password</Form.Label>
          <Form.Control type="password" placeholder="Password" />
        </Form.Group>
        <Button variant="primary" type="submit">
          Login
        </Button>
      </Form>
    </center>
  );
};

export default Login;
