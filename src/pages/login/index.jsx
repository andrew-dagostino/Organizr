import React from 'react';
import ReactDOM from 'react-dom';

import Person from '@material-ui/icons/Person';
import PersonAdd from '@material-ui/icons/PersonAdd';

import { makeStyles } from '@material-ui/core/styles';
import Box from '@material-ui/core/Box';
import AppBar from '@material-ui/core/AppBar';
import Button from '@material-ui/core/Button';
import Container from '@material-ui/core/Container';
import CssBaseline from '@material-ui/core/CssBaseline';
import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import Tab from '@material-ui/core/Tab';
import Tabs from '@material-ui/core/Tabs';
import TextField from '@material-ui/core/TextField';
import Typography from '@material-ui/core/Typography';

import LoginAppBar from '../../components/LoginAppBar.jsx';

import axios from 'axios';

function TabPanel(props) {
    const { children, value, index, ...other } = props;

    return (
        <div
            role="tabpanel"
            hidden={value !== index}
            id={`nav-tabpanel-${index}`}
            aria-labelledby={`nav-tab-${index}`}
            {...other}
        >
            {value === index && (
                <Box p={3}>
                    <Typography component={"span"}>{children}</Typography>
                </Box>
            )}
        </div>
    );
}

const useStyles = makeStyles((theme) => ({
    root: {
        flexGrow: 1,
        backgroundColor: theme.palette.background.paper,
    },
    largeIcon: {
        width: 64,
        height: 64
    },
}));

function a11yProps(index) {
    return {
        id: `nav-tab-${index}`,
        'aria-controls': `nav-tabpanel-${index}`,
    };
}

function NavTabs() {
    const classes = useStyles();
    const [value, setValue] = React.useState(0);

    const handleChange = (event, newValue) => {
        setValue(newValue);
    };

    return (
        <div className={classes.root}>
            <AppBar position="static">
                <Tabs variant="fullWidth" value={value} onChange={handleChange}>
                    <Tab label="Login" {...a11yProps(0)} />
                    <Tab label="Register" {...a11yProps(1)} />
                </Tabs>
            </AppBar>
            <TabPanel value={value} index={0}>
                <Container fixed maxWidth="xs">
                    <Grid container spacing={4} justify="center">
                        <Grid item>
                            <Person color="primary" className={classes.largeIcon} />
                        </Grid>
                    </Grid>
                    <Typography variant="h5" align="left" color="primary" gutterBottom noWrap>
                        Sign In
                    </Typography>
                    <form id="loginForm" onSubmit={submitLogin}>
                        <TextField
                            id="loginUsername"
                            name="username"
                            label="Username"
                            variant="outlined"
                            type="text"
                            autoComplete="current-username"
                            required
                            fullWidth
                            margin="normal"
                        />
                        <TextField
                            id="loginPassword"
                            name="password"
                            label="Password"
                            variant="outlined"
                            type="password"
                            autoComplete="current-password"
                            required
                            fullWidth
                            margin="normal"
                        />
                        <Grid container spacing={4} justify="flex-end">
                            <Grid item>
                                <Button variant="contained" fullWidth color="primary" type="submit">
                                    Submit
                                </Button>
                            </Grid>
                        </Grid>
                    </form>
                </Container>
            </TabPanel>
            <TabPanel value={value} index={1}>
                <Container fixed maxWidth="xs">
                    <Grid container spacing={4} justify="center">
                        <Grid item>
                            <PersonAdd color="primary" className={classes.largeIcon} />
                        </Grid>
                    </Grid>
                    <Typography variant="h5" align="left" color="primary" gutterBottom noWrap>
                        Create Account
                    </Typography>
                    <form id="registerForm" onSubmit={submitRegister}>
                        <TextField
                            id="registerUsername"
                            name="username"
                            label="Username"
                            variant="outlined"
                            type="text"
                            autoComplete="off"
                            required
                            fullWidth
                            margin="normal"
                        />
                        <TextField
                            id="registerPassword"
                            name="password"
                            label="Password"
                            variant="outlined"
                            type="password"
                            autoComplete="off"
                            required
                            fullWidth
                            margin="normal"
                        />
                        <TextField
                            id="registerPasswordConfirm"
                            label="Confirm Password"
                            variant="outlined"
                            type="password"
                            autoComplete="off"
                            required
                            fullWidth
                            margin="normal"
                        />
                        <Grid container spacing={4} justify="flex-end">
                            <Grid item>
                                <Button variant="contained" fullWidth color="primary" type="submit">
                                    Submit
                                </Button>
                            </Grid>
                        </Grid>
                    </form>
                </Container>
            </TabPanel>
        </div>
    );
}

ReactDOM.render(
    <React.Fragment>
        <CssBaseline />
        <LoginAppBar />
        <Container fixed maxWidth="xs">
            <Paper elevation={2}>
                <NavTabs />
            </Paper>
        </Container>
    </React.Fragment>,
    document.querySelector('#app')
);

function submitLogin(e) {
    e.preventDefault();

    let formdata = new FormData(document.getElementById("loginForm"));
    axios.post("/api/login", formdata).then(
        (res) => {
            console.log(res.data);

            axios.get(`/api/restricted/user/9`, { headers: { Authorization: `Bearer ${res.data.token}` } }).then(res => console.log(res.data))
        },
        (err) => {
            console.log(err.response ? err.response.data.error : err);
        }
    );
}

function submitRegister(e) {
    e.preventDefault();

    let formdata = new FormData(document.getElementById("registerForm"));
    axios.post("/api/register", formdata).then(
        (res) => {
            console.log(res.data);
        },
        (err) => {
            console.log(err.response ? err.response.data.error : err);
        }
    );
}