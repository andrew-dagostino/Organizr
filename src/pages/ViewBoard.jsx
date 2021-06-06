import React from 'react';
import { Card, Grid, Icon, Loader } from 'semantic-ui-react';
import PropTypes from 'prop-types';
import { DragDropContext } from 'react-beautiful-dnd';

import Column from '../components/Column';
import Header from '../components/Header';

import {
    retrieveBoard,
    updateBoard,
    retrieveColumns,
    updateColumn,
    createColumn,
    retrieveTasks,
    updateTask,
} from '../util/board_functions';

/**
 * Card widget linking to the new board page
 */
function AddColumnWidget(props) {
    const { onClick } = props;

    return (
        <Grid.Column>
            <Card
                onClick={onClick}
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
                        Add Column
                    </Card.Header>
                </Card.Content>
            </Card>
        </Grid.Column>
    );
}

AddColumnWidget.propTypes = {
    onClick: PropTypes.func.isRequired,
};

const reorder = (list, startIndex, endIndex) => {
    const result = Array.from(list);
    const [removed] = result.splice(startIndex, 1);
    result.splice(endIndex, 0, removed);

    return result;
};

const move = (
    source,
    destination = [],
    droppableSource,
    droppableDestination
) => {
    const sourceClone = Array.from(source);
    const destClone = Array.from(destination);
    const [removed] = sourceClone.splice(droppableSource.index, 1);

    destClone.splice(droppableDestination.index, 0, removed);

    const result = {};
    result[droppableSource.droppableId] = sourceClone;
    result[droppableDestination.droppableId] = destClone;

    return result;
};

export default class ViewBoard extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            title: '',
            gid: '',
            columns: [],
            loaded: false,
            titleTimer: undefined,
            columnTimers: {},
        };
    }

    componentDidMount() {
        const pathVars = window.location.pathname.split('/');
        const boardGid = pathVars[pathVars.length - 1];

        this.setState({ gid: boardGid });

        retrieveBoard(boardGid).then(({ data }) =>
            this.setState({ title: data.title })
        );
        retrieveColumns(boardGid).then(({ data }) => {
            const columns = data;
            columns.forEach(async (column, index, arr) => {
                const response = await retrieveTasks(column.gid);
                columns[index].tasks = [...response.data];

                if (index === arr.length - 1) {
                    this.setState({ columns, loaded: true });
                }
            });
            if (!columns.length) {
                this.setState({ loaded: true });
            }
        });
    }

    onDragEnd = (result) => {
        const { source, destination } = result;

        // dropped outside the list
        if (!destination) {
            return;
        }
        const sId = source.droppableId;
        const dId = destination.droppableId;

        const sColumn = this.getColumn(sId);
        const dColumn = this.getColumn(dId);

        if (sId === dId) {
            this.updateColumnUI({
                gid: sColumn.gid,
                title: sColumn.title,
                tasks: reorder(
                    sColumn.tasks,
                    result.source.index,
                    result.destination.index
                ),
            });
        } else {
            const newResult = move(
                sColumn.tasks,
                dColumn.tasks,
                source,
                destination
            );

            const taskGid = result.draggableId;
            const task = sColumn.tasks.filter((t) => t.gid === taskGid)[0];
            task.task_column_id = dColumn.id;

            const formdata = new FormData();
            formdata.append('title', task.title || '');
            formdata.append('description', task.description || '');

            updateTask(dColumn.gid, taskGid, formdata).then(() => {
                this.updateColumnUI({
                    gid: sColumn.gid,
                    title: sColumn.title,
                    tasks: newResult[sId],
                });
                this.updateColumnUI({
                    gid: dColumn.gid,
                    title: dColumn.title,
                    tasks: newResult[dId],
                });
            });
        }
    };

    addColumn = () => {
        const { gid } = this.state;

        const formdata = new FormData();
        formdata.append('title', '');

        createColumn(gid, formdata).then(({ data }) => {
            const { columns } = this.state;
            this.setState({ columns: columns.concat([data]) });
        });
    };

    updateColumnUI = (column) => {
        const { columns } = this.state;
        const index = columns.findIndex((col) => col.gid === column.gid);
        columns[index] = column;
        this.setState({ columns });
    };

    handleColumnChange = (column) => {
        const { gid, columns, columnTimers } = this.state;

        const index = columns.findIndex((col) => col.gid === column.gid);
        const oldColumn = columns[index];

        if (column.title !== oldColumn.title) {
            clearTimeout(columnTimers[column.gid]);

            const formdata = new FormData();
            formdata.append('title', column.title);

            columnTimers[column.gid] = setTimeout(
                () =>
                    updateColumn(gid, column.gid, formdata).then(({ data }) => {
                        this.updateColumnUI(data);
                    }),
                500
            );

            this.setState({ columnTimers });
        }

        this.updateColumnUI(column);
    };

    getColumn = (gid) => {
        const { columns } = this.state;
        return columns.filter((column) => column.gid === gid)[0];
    };

    handleBoardNameChange = (e, { value }) => {
        const { gid, titleTimer } = this.state;

        clearTimeout(titleTimer);
        if (value) {
            const formdata = new FormData();
            formdata.append('title', value);

            this.setState({
                title: value,
                titleTimer: setTimeout(
                    () =>
                        updateBoard(gid, formdata).then(({ data }) => {
                            this.setState({ title: data.title });
                        }),
                    500
                ),
            });
        }
    };

    render() {
        const { title, columns, loaded } = this.state;
        return (
            <>
                <Header
                    title={title}
                    handleChange={this.handleBoardNameChange}
                    showTextfield
                />
                {loaded ? (
                    <DragDropContext onDragEnd={this.onDragEnd}>
                        <Grid columns="4" container doubling stackable>
                            <Grid.Row style={{ height: '100%' }}>
                                {columns.map((column) => (
                                    <Column
                                        key={column.gid}
                                        gid={column.gid}
                                        title={column.title}
                                        tasks={column.tasks}
                                        updateColumn={this.handleColumnChange}
                                        getColumn={this.getColumn}
                                    />
                                ))}
                                <AddColumnWidget
                                    onClick={() => this.addColumn()}
                                />
                            </Grid.Row>
                        </Grid>
                    </DragDropContext>
                ) : (
                    <Loader active>Loading</Loader>
                )}
            </>
        );
    }
}
