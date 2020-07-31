import React from 'react';
import ReactDOM from 'react-dom';

import AccountCircle from '@material-ui/icons/AccountCircle';
import Lock from '@material-ui/icons/Lock';

import Button from '@material-ui/core/Button';
import Container from '@material-ui/core/Container';
import CssBaseline from '@material-ui/core/CssBaseline';
import Grid from '@material-ui/core/Grid';
import InputAdornment from '@material-ui/core/InputAdornment';
import InputLabelProps from '@material-ui/core/InputLabel';
import TextField from '@material-ui/core/TextField';

import LoginAppBar from '../../components/LoginAppBar.jsx';

import axios from 'axios';

export default function App() {
    return (
        <React.Fragment>
            <CssBaseline />
            <LoginAppBar />
            <Container fixed>
                <form id="loginForm" onSubmit={submitLogin}>
                    <Grid container spacing={3} direction="column">
                        <Grid item align="center">
                            <TextField id="username" name="username" label="Username" variant="outlined" type="text" autoComplete="current-username" required
                                InputProps={{
                                    startAdornment: (
                                        <InputAdornment position="start">
                                            <AccountCircle />
                                        </InputAdornment>
                                    ),
                                }}
                            />
                        </Grid>
                        <Grid item align="center">
                            <TextField id="password" name="password" label="Password" variant="outlined" type="password" autoComplete="current-password" required
                                InputProps={{
                                    startAdornment: (
                                        <InputAdornment position="start">
                                            <Lock />
                                        </InputAdornment>
                                    ),
                                }}
                            />
                        </Grid>
                        <Grid item align="center">
                            <Button variant="contained" color="primary" type="submit">
                                Submit
                            </Button>
                        </Grid>
                    </Grid>
                </form>
            </Container>
        </React.Fragment>
    );
}

function submitLogin(e) {
    e.preventDefault();

    let formdata = new FormData(document.getElementById("loginForm"));
    axios.post("/login", formdata).then(
        (res) => {

        },
        (err) => {

        }
    );
}

ReactDOM.render(<App />, document.querySelector('#app'));