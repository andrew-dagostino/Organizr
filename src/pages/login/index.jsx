import React from 'react';
import ReactDOM from 'react-dom';

import axios from 'axios';

import {
    Button,
    Container,
    Card,
    CardHeader,
    CardBody,
    CardText,
    CardTitle,
    Col,
    Form,
    FormControl,
    FormGroup,
    FormLabel,
    FormText,
    Nav,
    NavItem,
    NavLink,
    Row,
    Tab,
    TabContainer,
    TabContent,
    Tabs,
} from 'react-bootstrap';

import LoginAppBar from '../../components/LoginAppBar.jsx';

function LoginRegisterCard() {
    const submitLogin = (e) => {
        e.preventDefault();

        let formdata = new FormData(document.getElementById('loginForm'));
        axios.post('/api/login', formdata).then(
            (res) => {
                localStorage.setItem('session', res.data.token);
                !document.referrer || document.location == document.referrer
                    ? location.replace('/')
                    : location.replace(document.referrer);
            },
            (err) => {
                let msg =
                    err.response && err.response.data && err.response.data.error
                        ? err.response.data.error
                        : err.message;
                toastr.error(msg);
            },
        );
    };

    const submitRegister = (e) => {
        e.preventDefault();

        let formdata = new FormData(document.getElementById('registerForm'));
        axios.post('/api/register', formdata).then(
            (res) => {
                toastr.success('Registration Successful!');
                document.getElementById('loginTabHeader').click();
            },
            (err) => {
                let msg =
                    err.response && err.response.data && err.response.data.error
                        ? err.response.data.error
                        : err.message;
                toastr.error(msg);
            },
        );
    };

    return (
        <Row className='justify-content-center'>
            <Col xs={12} sm={8} md={6}>
                <Card>
                    <Tab.Container defaultActiveKey='login'>
                        <Card.Header as='h5'>
                            <Nav variant='pills' className='text-center'>
                                <Nav.Item className='w-50'>
                                    <Nav.Link id='loginTabHeader' eventKey='login'>
                                        Login
                                    </Nav.Link>
                                </Nav.Item>
                                <Nav.Item className='w-50'>
                                    <Nav.Link id='registerTabHeader' eventKey='register'>
                                        Register
                                    </Nav.Link>
                                </Nav.Item>
                            </Nav>
                        </Card.Header>
                        <Card.Body bg='light'>
                            <Tab.Content>
                                <Tab.Pane eventKey='login'>
                                    <Form id='loginForm' onSubmit={submitLogin}>
                                        <Form.Group controlId='loginUsername'>
                                            <Form.Label>Username</Form.Label>
                                            <Form.Control type='text' name='username' />
                                        </Form.Group>

                                        <Form.Group controlId='loginPassword'>
                                            <Form.Label>Password</Form.Label>
                                            <Form.Control type='password' name='password' />
                                        </Form.Group>
                                        <Button variant='primary' type='submit'>
                                            Submit
                                        </Button>
                                    </Form>
                                </Tab.Pane>
                                <Tab.Pane eventKey='register'>
                                    <Form id='registerForm' onSubmit={submitRegister}>
                                        <Form.Group controlId='registerUsername'>
                                            <Form.Label>Username</Form.Label>
                                            <Form.Control type='text' name='username' />
                                        </Form.Group>

                                        <Form.Group controlId='registerPassword'>
                                            <Form.Label>Password</Form.Label>
                                            <Form.Control type='password' name='password' />
                                        </Form.Group>

                                        <Form.Group controlId='registerPasswordConfirm'>
                                            <Form.Label>Confirm Password</Form.Label>
                                            <Form.Control type='password' />
                                        </Form.Group>
                                        <Button variant='primary' type='submit'>
                                            Submit
                                        </Button>
                                    </Form>
                                </Tab.Pane>
                            </Tab.Content>
                        </Card.Body>
                    </Tab.Container>
                </Card>
            </Col>
        </Row>
    );
}

ReactDOM.render(
    <React.Fragment>
        <LoginAppBar />
        <Container>
            <LoginRegisterCard />
        </Container>
    </React.Fragment>,
    document.querySelector('#app'),
);
