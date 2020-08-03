import React from 'react';

import { Button, Container, Navbar, NavbarBrand, NavbarToggle, NavbarCollapse, NavLink, Nav } from 'react-bootstrap';

export default function LoginAppBar(props) {
    return (
        <Navbar sticky='top' collapseOnSelect expand='md' bg='primary' variant='dark' className='mb-5'>
            <Container>
                <Navbar.Brand href='/'>Login / Register</Navbar.Brand>
                <Navbar.Toggle />
                <Navbar.Collapse>
                    <Nav className='mr-auto'></Nav>
                </Navbar.Collapse>
            </Container>
        </Navbar>
    );
}
