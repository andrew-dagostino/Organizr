import React from 'react';
import ReactDOM from 'react-dom';
import Container from '@material-ui/core/Container';
import CssBaseline from '@material-ui/core/CssBaseline';

import LandingAppBar from '../components/LandingAppBar.jsx';

const jwtDecode = require('jwt-decode');

export default function App() {
    const storedJWT = localStorage.getItem("session");
    const [jwt, setJWT] = React.useState(storedJWT || null);

    // Remove expired session
    if(jwt && jwtDecode(jwt).exp < parseInt(Date.now()/1000)) {
        setJWT(null);
        localStorage.removeItem("session");
        location.replace("/");
    }

    return (
        <React.Fragment>
            <CssBaseline />
            <LandingAppBar jwt={jwt} setJWT={setJWT} />
        </React.Fragment>
    );
}

ReactDOM.render(<App />, document.querySelector('#app'));