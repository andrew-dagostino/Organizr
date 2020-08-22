import React from 'react';

import Card from 'react-bootstrap/Card';

export default function Post(props) {
    const styles = {
        maxHeight: '15rem',
        objectFit: 'contain',
        background: '#000',
    };

    return (
        <Card>
            {props.img ? <Card.Img variant='top' src={props.img} style={styles} /> : ''}
            <Card.Body>
                <Card.Title className='text-center'>{props.title}</Card.Title>
                <Card.Text>{props.text}</Card.Text>
            </Card.Body>
        </Card>
    );
}
