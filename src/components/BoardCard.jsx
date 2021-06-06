import React from 'react';
import PropTypes from 'prop-types';

import { Card, Dropdown, Grid, Icon, Image } from 'semantic-ui-react';
import { Link } from 'react-router-dom';

function OptionsMenu(props) {
    const { deleteBoard } = props;

    return (
        <Dropdown compact icon="vertical ellipsis" className="icon">
            <Dropdown.Menu>
                <Dropdown.Header content="Actions" style={{ margin: 0 }} />
                <Dropdown.Divider />
                <Dropdown.Item icon="trash" onClick={deleteBoard}>
                    Delete
                </Dropdown.Item>
            </Dropdown.Menu>
        </Dropdown>
    );
}

OptionsMenu.propTypes = {
    deleteBoard: PropTypes.func.isRequired,
};

/**
 * Card overview of board
 */
function BoardCard({ gid, title, memberCount, deleteBoard }) {
    return (
        <Card style={{ marginBottom: '2rem' }} color="black">
            <Image src="/img/blue-square.png" style={{ maxHeight: '15rem' }} />
            <Card.Content>
                <Card.Header>
                    <Link
                        to={`/board/${gid}`}
                        style={{
                            color: 'black',
                            display: 'block',
                            width: '100%',
                        }}
                    >
                        {title || '...'}
                    </Link>
                </Card.Header>
            </Card.Content>
            <Card.Content extra>
                <Grid columns="2">
                    <Grid.Row>
                        <Grid.Column width="13">
                            <Icon name="small user friends" />
                            <span
                                style={{
                                    verticalAlign: 'top',
                                    marginLeft: '0.5rem',
                                }}
                            >
                                {memberCount} Member
                                {memberCount === 1 ? '' : 's'}
                            </span>
                        </Grid.Column>
                        <Grid.Column width="3" textAlign="center">
                            <OptionsMenu
                                deleteBoard={() => deleteBoard({ gid })}
                            />
                        </Grid.Column>
                    </Grid.Row>
                </Grid>
            </Card.Content>
        </Card>
    );
}

BoardCard.propTypes = {
    gid: PropTypes.string,
    title: PropTypes.string,
    memberCount: PropTypes.number,
    deleteBoard: PropTypes.func.isRequired,
};

BoardCard.defaultProps = {
    gid: '',
    title: '...',
    memberCount: 0,
};

export default BoardCard;
