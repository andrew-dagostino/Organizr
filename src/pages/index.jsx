import React from 'react';
import ReactDOM from 'react-dom';
import Button from '@material-ui/core/Button';
import Container from '@material-ui/core/Container';
import CssBaseline from '@material-ui/core/CssBaseline';

import LandingAppBar from '../components/LandingAppBar.jsx';

export default function App() {
    return (
        <React.Fragment>
            <CssBaseline />
            <LandingAppBar />
            <Container fixed>
                <Button variant="contained" color="primary">
                    Hello World
                </Button>
            </Container>
        </React.Fragment>
    );
}

ReactDOM.render(<App />, document.querySelector('#app'));