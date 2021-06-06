import React from 'react';
import { Card, Grid, Icon, Loader } from 'semantic-ui-react';
import BoardCard from '../components/BoardCard';
import Header from '../components/Header';

import { createBoard, retrieveBoards } from '../util/board_functions';

/**
 * Card widget linking to the new board page
 */
function AddBoardWidget() {
    function handleAddBoard() {
        const formdata = new FormData();
        formdata.append('title', '');

        createBoard(formdata).then(({ data }) => {
            window.location.replace(`/board/${data.gid}`);
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
                        name="plus"
                        style={{
                            fontSize: '5rem',
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
            loaded: false,
        };
    }

    componentDidMount() {
        retrieveBoards().then(({ data }) => {
            this.setState({ boards: data, loaded: true });
        });
    }

    render() {
        const { boards, loaded } = this.state;

        return (
            <>
                <Header />
                {loaded ? (
                    <Grid columns="4" container doubling stackable>
                        <Grid.Row>
                            {boards.map((board) => (
                                <Grid.Column key={board.gid}>
                                    <BoardCard
                                        id={board.gid}
                                        title={board.title}
                                        memberCount={board.board_member_count}
                                    />
                                </Grid.Column>
                            ))}
                            <AddBoardWidget />
                        </Grid.Row>
                    </Grid>
                ) : (
                    <Loader active>Loading</Loader>
                )}
            </>
        );
    }
}
