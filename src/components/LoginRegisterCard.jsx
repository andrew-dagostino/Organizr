import React from 'react';
import { Button, Card, Form, Grid, Message, Tab } from 'semantic-ui-react';
import PropTypes from 'prop-types';
import axios from 'axios';

import config from '../config.json';

function LoginForm({ setError }) {
    const [username, setUsername] = React.useState('');
    const [password, setPassword] = React.useState('');
    const [remember, setRemember] = React.useState(false);

    const [loading, setLoading] = React.useState(false);

    function handleLogin() {
        setError('');
        setLoading(true);

        const formdata = new FormData();
        formdata.append('username', username);
        formdata.append('password', password);
        formdata.append('remember', remember);

        axios
            .post(`${config.UNAUTH_API_URL}/login`, formdata, {
                headers: { 'Content-Type': 'multipart/form-data' },
            })
            .then((response) => {
                window.localStorage.setItem('jwt', response.data.jwt);
                window.location.replace('/board');
            })
            .catch((error) => {
                setError(
                    error.response ? error.response.data.message : error.message
                );
            })
            .finally(() => setLoading(false));
    }

    return (
        <Form
            onSubmit={handleLogin}
            style={{
                textAlign: 'left',
            }}
        >
            <Form.Input
                type="text"
                label="Username / Email"
                required
                onChange={(e, { value }) => setUsername(value)}
            />
            <Form.Input
                type="password"
                label="Password"
                required
                onChange={(e, { value }) => setPassword(value)}
            />
            <Form.Field>
                <Form.Checkbox
                    label="Remember Me"
                    onChange={(e, { checked }) => setRemember(checked)}
                />
            </Form.Field>
            <Grid columns="3">
                <Grid.Row centered>
                    <Grid.Column width="6">
                        <Button
                            size="medium"
                            color="blue"
                            fluid
                            loading={loading}
                            style={{ margin: 'auto' }}
                        >
                            Login
                        </Button>
                    </Grid.Column>
                </Grid.Row>
            </Grid>
        </Form>
    );
}
LoginForm.propTypes = PropTypes.func.isRequired;

function RegisterForm({ setError }) {
    const [username, setUsername] = React.useState('');
    const [email, setEmail] = React.useState('');
    const [password, setPassword] = React.useState('');
    const [confirmPassword, setConfirmPassword] = React.useState('');

    const [loading, setLoading] = React.useState(false);

    function handleRegister() {
        if (password !== confirmPassword) {
            setError('Passwords do not match');
            return;
        }

        setError('');
        setLoading(true);

        const formdata = new FormData();
        formdata.append('username', username);
        formdata.append('email', email);
        formdata.append('password', password);

        axios
            .post(`${config.UNAUTH_API_URL}/register`, formdata, {
                headers: { 'Content-Type': 'multipart/form-data' },
            })
            .then((response) => {
                window.localStorage.setItem('jwt', response.data.jwt);
                window.location.replace('/board');
            })
            .catch((error) => {
                setError(
                    error.response ? error.response.data.message : error.message
                );
            })
            .finally(() => setLoading(false));
    }

    return (
        <Form
            onSubmit={handleRegister}
            style={{
                textAlign: 'left',
            }}
        >
            <Form.Input
                type="text"
                label="Username"
                required
                onChange={(e, { value }) => setUsername(value)}
            />
            <Form.Input
                type="email"
                label="Email"
                required
                onChange={(e, { value }) => setEmail(value)}
            />
            <Form.Input
                type="password"
                label="Password"
                required
                onChange={(e, { value }) => setPassword(value)}
            />
            <Form.Input
                type="password"
                label="Confirm Password"
                required
                onChange={(e, { value }) => setConfirmPassword(value)}
            />
            <Grid columns="3">
                <Grid.Row centered>
                    <Grid.Column width="6">
                        <Button
                            size="medium"
                            color="blue"
                            fluid
                            loading={loading}
                            style={{ margin: 'auto' }}
                        >
                            Register
                        </Button>
                    </Grid.Column>
                </Grid.Row>
            </Grid>
        </Form>
    );
}
RegisterForm.propTypes = PropTypes.func.isRequired;

export default function LoginRegisterCard() {
    const [error, setError] = React.useState('');

    const panes = [
        {
            menuItem: 'Login',
            render: function renderLogin() {
                return (
                    <Tab.Pane style={{ border: 'none' }}>
                        <LoginForm setError={setError} />
                    </Tab.Pane>
                );
            },
        },
        {
            menuItem: 'Register',
            render: function renderRegister() {
                return (
                    <Tab.Pane style={{ border: 'none' }}>
                        <RegisterForm setError={setError} />
                    </Tab.Pane>
                );
            },
        },
    ];

    return (
        <Card fluid raised style={{ margin: 'auto' }}>
            <Card.Content>
                <Message negative hidden={error === ''}>
                    <Message.Header content={error} />
                </Message>
                <Tab
                    menu={{
                        secondary: true,
                        pointing: true,
                        widths: '2',
                    }}
                    panes={panes}
                    defaultActiveIndex="1"
                />
            </Card.Content>
        </Card>
    );
}
