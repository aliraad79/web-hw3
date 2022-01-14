import { Container, Navbar, Nav } from "react-bootstrap";

const MyNavbar = ({ getAuthToken }) => {
  const token = getAuthToken();
  return (
    <>
      <Navbar bg="dark" variant="dark">
        <Container>
          <Navbar.Brand href="#">Notes</Navbar.Brand>
          <Navbar.Collapse id="basic-navbar-nav">
            <Nav className="me-auto">
              {token ? (
                <>
                  <Nav.Link href="/notes">Home</Nav.Link>
                  <Nav.Link href="/signout">Sign Out</Nav.Link>
                </>
              ) : (
                <>
                  <Nav.Link href="/signup">Sign Up</Nav.Link>
                  <Nav.Link href="/">Login</Nav.Link>
                </>
              )}
            </Nav>
          </Navbar.Collapse>
        </Container>
      </Navbar>
    </>
  );
};

export default MyNavbar;
