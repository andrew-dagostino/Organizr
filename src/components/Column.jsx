import React from 'react';
import PropTypes, { objectOf } from 'prop-types';

import { Button, Card, Grid, Icon, Input } from 'semantic-ui-react';
import { Droppable } from 'react-beautiful-dnd';
import axios from 'axios';

import config from '../config.json';

import Task from './Task';

const JWT = window.localStorage.getItem('jwt');

class Column extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            taskTimers: {},
        };
    }

    addTask = () => {
        const { gid, title, tasks, updateColumn } = this.props;

        const formdata = new FormData();
        formdata.append('title', '');
        formdata.append('description', '');

        this.createTask(gid, formdata).then((data) => {
            updateColumn({
                gid,
                title,
                tasks: tasks.concat([data]),
            });
        });
    };

    createTask = (cGid, formdata) =>
        axios.post(`${config.API_URL}/task/${cGid}`, formdata, {
            headers: {
                'Content-Type': 'multipart/form-data',
                Authorization: `Bearer ${JWT}`,
            },
        });

    updateTaskUI = (task) => {
        const { gid, title, tasks, updateColumn } = this.props;
        const index = tasks.findIndex((t) => t.gid === task.gid);
        tasks[index] = task;
        updateColumn({ gid, title, tasks });
    };

    updateTask = (cGid, tGid, formdata) =>
        axios.put(`${config.API_URL}/task/${cGid}/${tGid}`, formdata, {
            headers: {
                'Content-Type': 'multipart/form-data',
                Authorization: `Bearer ${JWT}`,
            },
        });

    handleTaskChange = (task) => {
        const { taskTimers } = this.state;
        const { tasks, gid } = this.props;

        const index = tasks.findIndex((t) => t.gid === task.gid);
        const oldTask = tasks[index];

        if (
            task.title !== oldTask.title ||
            task.description !== oldTask.description
        ) {
            clearTimeout(taskTimers[task.gid]);

            const formdata = new FormData();
            formdata.append('title', task.title);
            formdata.append('description', task.description);

            taskTimers[task.gid] = setTimeout(
                () => this.updateTask(gid, task.gid, formdata),
                500
            );

            this.setState({ taskTimers });
        }

        this.updateTaskUI(task); // Updates UI
    };

    handleTitle = (value) => {
        const { gid, tasks, updateColumn } = this.props;
        updateColumn({ gid, title: value, tasks });
    };

    render() {
        const { gid, title, tasks } = this.props;

        return (
            <Grid.Column style={{ marginBottom: '2rem' }}>
                <Card>
                    <Card.Content
                        style={{
                            paddingBottom: tasks.length ? '0rem' : '7.5rem',
                        }}
                    >
                        <Card.Header>
                            <Input
                                placeholder="Column Name..."
                                fluid
                                size="small"
                                transparent
                                value={title}
                                onChange={(e, { value }) =>
                                    this.handleTitle(value)
                                }
                            />
                        </Card.Header>
                        <hr />
                        <Button
                            type="button"
                            compact
                            basic
                            fluid
                            color="grey"
                            onClick={() => this.addTask()}
                        >
                            <Icon name="add" />
                            <span style={{ verticalAlign: 'text-top' }}>
                                Add Task
                            </span>
                        </Button>
                        <br />
                        <Droppable droppableId={gid}>
                            {(provided) => (
                                <div
                                    // eslint-disable-next-line react/jsx-props-no-spreading
                                    {...provided.droppableProps}
                                    ref={provided.innerRef}
                                >
                                    {tasks.map((task, index) => (
                                        <Task
                                            key={task.gid}
                                            gid={task.gid}
                                            title={task.title}
                                            description={task.description}
                                            index={index}
                                            updateTask={this.handleTaskChange}
                                        />
                                    ))}
                                    {provided.placeholder}
                                </div>
                            )}
                        </Droppable>
                    </Card.Content>
                </Card>
            </Grid.Column>
        );
    }
}

Column.propTypes = {
    gid: PropTypes.string.isRequired,
    title: PropTypes.string,
    tasks: PropTypes.arrayOf(
        objectOf({
            gid: PropTypes.string.isRequired,
            title: PropTypes.string,
            description: PropTypes.string,
        })
    ),
    updateColumn: PropTypes.func.isRequired,
    getColumn: PropTypes.func.isRequired,
};

Column.defaultProps = {
    title: '',
    tasks: [],
};

export default Column;
