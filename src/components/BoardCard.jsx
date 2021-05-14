import React from 'react';
import PropTypes from 'prop-types';

import { Card, Icon, Image } from 'semantic-ui-react';
import { Link } from 'react-router-dom';

/**
 * Card overview of board
 */
function BoardCard({ id, name, memberCount }) {
    return (
        <Link to={`/board/${id}`}>
            <Card style={{ marginBottom: '2rem' }} color="black">
                <Image
                    src="/img/blue-square.png"
                    style={{ maxHeight: '15rem' }}
                />
                <Card.Content>
                    <Card.Header>{name}</Card.Header>
                </Card.Content>
                <Card.Content extra>
                    <Icon name="person" />
                    <span style={{ verticalAlign: 'top' }}>
                        {memberCount} Member(s)
                    </span>
                </Card.Content>
            </Card>
        </Link>
    );
}

BoardCard.propTypes = {
    id: PropTypes.string,
    name: PropTypes.string,
    memberCount: PropTypes.number,
};

BoardCard.defaultProps = {
    id: '',
    name: '...',
    memberCount: 0,
};

export default BoardCard;
