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
import {BASE_SERVER_URL} from "../consts";
import MyNavbar from "./Navbar";
import MyModal from "./MyModal";

const SignUp = ({ getAuthToken }) => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [adminSecret, setadminSecret] = useState("");
  const [showErrorModal, setShowErrorModal] = useState(false);
  const [adminMode, setAdminMode] = useState(false);

  const navigate = useNavigate();

  const signup_user = async (event) => {
    event.preventDefault();
    await fetch(`${BASE_SERVER_URL}/signup`, {
      method: "POST",
      body: JSON.stringify({
        username,
        password,
        secret_phrase: adminSecret,
      }),
    })
      .then((response) => {
        if (response.status === 200) {
          return response.json();
        } else {
          setShowErrorModal(true);
          return;
        }
      })
      .then((response) => {
        navigate("/", { replace: true });
      });
  };

  const token = getAuthToken();
  return token ? (
    <Navigate to={{ pathname: "/notes" }} />
  ) : (
    <>
      <MyNavbar getAuthToken={getAuthToken} />
      <br />
      <Container>
        <Row>
          <Col></Col>
          <Col xs={6}>
            <h2>Sign Up</h2>
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
                borderColor: "#9bc5c3 #616161 #616161 #9bc5c3",
                fontWeight: "550",
              }}
            >
              <Form onSubmit={signup_user}>
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
                {adminMode && (
                  <FloatingLabel label="Admins Secret">
                    <Form.Control
                      type="password"
                      onChange={(e) => {
                        setadminSecret(e.target.value);
                      }}
                    />
                  </FloatingLabel>
                )}
                <Row>
                  <Col></Col>
                  <Col>
                    <Button variant="info" type="submit">
                      Sign UP
                    </Button>
                  </Col>
                  <Col>
                    <Button
                      variant="light"
                      onClick={(e) => {
                        setAdminMode(!adminMode);
                      }}
                    >
                      Admin signup
                    </Button>
                  </Col>
                </Row>
              </Form>
            </div>
          </Col>
          <Col></Col>
        </Row>

        {showErrorModal && (
          <MyModal text="Username or password is incorrect" title="Error" />
        )}
      </Container>
    </>
  );
};

export default SignUp;
