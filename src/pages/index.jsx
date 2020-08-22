import React from 'react';
import ReactDOM from 'react-dom';

import { Button, Col, Container, Row } from 'react-bootstrap';

import LandingAppBar from '../components/LandingAppBar.jsx';
import PostTimeline from '../components/PostTimeline.jsx';
import PostModal from '../components/PostModal.jsx';

const jwtDecode = require('jwt-decode');

export default function App() {
    const storedJWT = localStorage.getItem('session');
    const [jwt, setJWT] = React.useState(storedJWT || null);

    const [show, showModal] = React.useState(false);

    // Remove expired session
    if (jwt && jwtDecode(jwt).exp < parseInt(Date.now() / 1000)) {
        setJWT(null);
        localStorage.removeItem('session');
        location.replace('/');
    }

    return (
        <React.Fragment>
            <LandingAppBar jwt={jwt} setJWT={setJWT} />
            <Container>
                <Row>
                    <Col>
                        <Button variant='secondary' onClick={() => showModal(true)}>
                            Add Post
                        </Button>
                    </Col>
                </Row>
                <PostTimeline jwt={jwt} />
            </Container>
            <PostModal jwt={jwt} show={show} showModal={showModal} />
        </React.Fragment>
    );
}

ReactDOM.render(<App />, document.querySelector('#app'));
