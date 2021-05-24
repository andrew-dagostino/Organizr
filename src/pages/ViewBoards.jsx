import axios from 'axios';
import React from 'react';
import { Card, Grid, Icon } from 'semantic-ui-react';
import BoardCard from '../components/BoardCard';
import Header from '../components/Header';

import config from '../config.json';

const JWT = window.localStorage.getItem('jwt');

/**
 * Card widget linking to the new board page
 */
function AddBoardWidget() {
    function handleAddBoard() {
        const formdata = new FormData();
        formdata.append('title', '');

        axios
            .post(`${config.API_URL}/board`, formdata, {
                headers: {
                    'Content-Type': 'multipart/form-data',
                    Authorization: `Bearer ${JWT}`,
                },
            })
            .then((response) => {
                window.location.replace(`/board/${response.data.gid}`);
            });
    }

    return (
        <Grid.Column>
            <Card
                onClick={handleAddBoard}
                style={{
                    marginBottom: '2rem',
                    backgroundColor: '#efefef',
                    color: '#afafaf',
                }}
            >
                <div
                    style={{
                        height: '15rem',
                        width: '100%',
                        display: 'flex',
                    }}
                >
                    <Icon
                        name="add"
                        style={{
                            fontSize: '10rem',
                            margin: 'auto auto',
                        }}
                    />
                </div>
                <Card.Content>
                    <Card.Header
                        textAlign="center"
                        style={{
                            color: '#afafaf',
                        }}
                    >
                        Add Board
                    </Card.Header>
                </Card.Content>
            </Card>
        </Grid.Column>
    );
}

/**
 * Grid view of existing boards
 */
export default class ViewBoards extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            boards: [],
        };
    }

    render() {
        const { boards } = this.state;

        return (
            <>
                <Header />
                <Grid columns="4" container doubling stackable>
                    <Grid.Row>
                        {boards.map((board) => (
                            <Grid.Column key={board.id}>
                                <BoardCard
                                    id={board.id}
                                    name={board.name}
                                    memberCount={board.memberCount}
                                />
                            </Grid.Column>
                        ))}
                        <AddBoardWidget />
                    </Grid.Row>
                </Grid>
            </>
        );
    }
}
