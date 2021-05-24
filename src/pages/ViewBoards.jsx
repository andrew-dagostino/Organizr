import React from 'react';
import { Link } from 'react-router-dom';
import { Card, Grid, Icon } from 'semantic-ui-react';
import { v4 as uuidv4 } from 'uuid';
import BoardCard from '../components/BoardCard';
import Header from '../components/Header';

/**
 * Card widget linking to the new board page
 */
function AddBoardWidget() {
    return (
        <Grid.Column>
            <Link to={`/board/${uuidv4()}`}>
                <Card
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
            </Link>
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
