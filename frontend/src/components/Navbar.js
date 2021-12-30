import { Container, Navbar, Nav } from "react-bootstrap";

const MyNavbar = () => {
  const token = localStorage.token;
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
                <Nav.Link href="/">Sign In</Nav.Link>
              )}
            </Nav>
          </Navbar.Collapse>
        </Container>
      </Navbar>
      ;
    </>
  );
};

export default MyNavbar;
