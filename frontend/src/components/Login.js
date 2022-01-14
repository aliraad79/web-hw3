import {
  Form,
  Button,
  FloatingLabel,
  Container,
  Row,
  Col,
} from "react-bootstrap";
import { useState } from "react";
import { Navigate, useNavigate } from "react-router-dom";
import consts from "../consts";
import MyNavbar from "./Navbar";
import MyModal from "./MyModal";

const Login = ({ setAuthToken, getAuthToken }) => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [showModalError, setShowModalError] = useState(false);

  const navigate = useNavigate();

  const login_user = async (event) => {
    event.preventDefault();
    await fetch(`${consts.BASE_SERVER_URL}/login`, {
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
        setAuthToken(`${response.token}`);
        navigate("/notes", { replace: true });
      });
  };

  const token = getAuthToken();
  return token ? (
    <Navigate to={{ pathname: "/notes" }} />
  ) : (
    <>
      <MyNavbar getAuthToken={getAuthToken}/>
      <br />
      <Container>
        <Row>
          <Col></Col>
          <Col xs={6}>
            <h2>Login</h2>
          </Col>
          <Col></Col>
        </Row>
        <Row>
          <Col></Col>
          <Col xs={6}>
            <div
              style={{
                boxShadow: "5px 5px 2px #9E9E9E",
                border: "6px solid",
                borderColor: "#616161 #9bc5c3 #9bc5c3 #616161",
                fontWeight: "550",
              }}
            >
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
            </div>
          </Col>
          <Col></Col>
        </Row>

        {showModalError && (
          <MyModal text="User not found" title="Login Error" />
        )}
      </Container>
    </>
  );
};

export default Login;
