import React from 'react';
import ReactDOM from 'react-dom';
import Container from '@material-ui/core/Container';
import CssBaseline from '@material-ui/core/CssBaseline';

import LandingAppBar from '../components/LandingAppBar.jsx';

export default function App() {
    return (
        <React.Fragment>
            <CssBaseline />
            <LandingAppBar />
            <Container fixed>
            </Container>
        </React.Fragment>
    );
}

ReactDOM.render(<App />, document.querySelector('#app'));