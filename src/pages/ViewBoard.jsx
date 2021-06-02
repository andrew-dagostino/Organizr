import React from 'react';
import { Card, Grid, Icon } from 'semantic-ui-react';
import PropTypes from 'prop-types';
import { DragDropContext } from 'react-beautiful-dnd';
import axios from 'axios';

import Column from '../components/Column';
import Header from '../components/Header';

import config from '../config.json';

const JWT = window.localStorage.getItem('jwt');

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
            titleTimer: undefined,
            columnTimers: {},
        };
    }

    componentDidMount() {
        const pathVars = window.location.pathname.split('/');
        const boardGid = pathVars[pathVars.length - 1];

        this.setState({ gid: boardGid });

        this.retrieveBoard(boardGid).then(({ data }) =>
            this.setState({ title: data.title })
        );
        this.retrieveColumns(boardGid).then(({ data }) => {
            const columns = data;
            columns.forEach(async (column, index) => {
                const response = await this.retrieveTasks(column.gid);
                columns[index].tasks = [...response.data];
                this.setState({ columns });
            });
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

            this.updateTask(dColumn.gid, taskGid, formdata).then(() => {
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
        this.createColumn(gid, formdata);
    };

    createColumn = (gid, formdata) => {
        axios
            .post(`${config.API_URL}/column/${gid}`, formdata, {
                headers: {
                    'Content-Type': 'multipart/form-data',
                    Authorization: `Bearer ${JWT}`,
                },
            })
            .then(({ data }) => {
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
                () => this.updateColumn(gid, column.gid, formdata),
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

    retrieveBoard = (gid) =>
        axios.get(`${config.API_URL}/board/${gid}`, {
            headers: {
                Authorization: `Bearer ${JWT}`,
            },
        });

    updateBoard = (gid, formdata) =>
        axios
            .put(`${config.API_URL}/board/${gid}`, formdata, {
                headers: {
                    'Content-Type': 'multipart/form-data',
                    Authorization: `Bearer ${JWT}`,
                },
            })
            .then(({ data }) => {
                this.setState({ title: data.title });
            });

    retrieveColumns = (gid) =>
        axios.get(`${config.API_URL}/column/${gid}`, {
            headers: {
                Authorization: `Bearer ${JWT}`,
            },
        });

    updateColumn = (bGid, cGid, formdata) =>
        axios
            .put(`${config.API_URL}/column/${bGid}/${cGid}`, formdata, {
                headers: {
                    'Content-Type': 'multipart/form-data',
                    Authorization: `Bearer ${JWT}`,
                },
            })
            .then(({ data }) => {
                this.updateColumnUI(data);
            });

    retrieveTasks = (cGid) =>
        axios.get(`${config.API_URL}/task/${cGid}`, {
            headers: {
                Authorization: `Bearer ${JWT}`,
            },
        });

    updateTask = (cGid, tGid, formdata) =>
        axios.put(`${config.API_URL}/task/${cGid}/${tGid}`, formdata, {
            headers: {
                'Content-Type': 'multipart/form-data',
                Authorization: `Bearer ${JWT}`,
            },
        });

    handleBoardNameChange = (e, { value }) => {
        const { gid, titleTimer } = this.state;

        clearTimeout(titleTimer);
        if (value) {
            const formdata = new FormData();
            formdata.append('title', value);

            this.setState({
                title: value,
                titleTimer: setTimeout(
                    () => this.updateBoard(gid, formdata),
                    500
                ),
            });
        }
    };

    render() {
        const { title, columns } = this.state;
        return (
            <>
                <Header
                    title={title}
                    handleChange={this.handleBoardNameChange}
                    showTextfield
                />
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
                            <AddColumnWidget onClick={() => this.addColumn()} />
                        </Grid.Row>
                    </Grid>
                </DragDropContext>
            </>
        );
    }
}
