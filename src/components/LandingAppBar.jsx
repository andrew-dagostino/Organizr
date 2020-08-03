import React from 'react';

import { Button, Container, Navbar, NavbarBrand, NavbarToggle, NavbarCollapse, NavLink, Nav } from 'react-bootstrap';

const jwtDecode = require('jwt-decode');

export default function LandingAppBar(props) {
    const submitLogout = (e) => {
        props.setJWT(null);
        localStorage.removeItem('session');
        location.replace('/');
    };

    return (
        <Navbar sticky='top' collapseOnSelect expand='md' bg='primary' variant='dark' className='mb-5'>
            <Container>
                <Navbar.Brand href='/'>
                    {props.jwt ? 'Welcome, ' + jwtDecode(props.jwt).username : 'Landing Page'}
                </Navbar.Brand>
                <Navbar.Toggle />
                <Navbar.Collapse>
                    <Nav className='mr-auto'></Nav>
                    <Nav>
                        {props.jwt ? (
                            <Nav.Link href='javascript:void(0)' onClick={submitLogout}>
                                Logout
                            </Nav.Link>
                        ) : (
                            <Nav.Link href='/login'>Login</Nav.Link>
                        )}
                    </Nav>
                </Navbar.Collapse>
            </Container>
        </Navbar>
    );
}
