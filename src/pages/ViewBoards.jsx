import React from 'react';
import { Card, Grid, Icon, Loader } from 'semantic-ui-react';
import BoardCard from '../components/BoardCard';
import Header from '../components/Header';

import {
    createBoard,
    retrieveBoards,
    deleteBoard,
} from '../util/board_functions';

/**
 * Card widget linking to the new board page
 */
function AddBoardWidget() {
    function handleAddBoard() {
        const formdata = new FormData();
        formdata.append('title', '');

        createBoard(formdata).then(({ data }) =>
            window.location.replace(`/board/${data.gid}`)
        );
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

    removeBoard = (bGid) => {
        const { boards } = this.state;
        boards.splice(
            boards.findIndex((b) => b.gid === bGid),
            1
        );
        this.setState({ boards });
    };

    handleBoardRemove = (board) => {
        deleteBoard(board.gid).then(() => {
            this.removeBoard(board.gid);
        });
    };

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
                                        gid={board.gid}
                                        title={board.title}
                                        memberCount={board.board_member_count}
                                        deleteBoard={this.handleBoardRemove}
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
