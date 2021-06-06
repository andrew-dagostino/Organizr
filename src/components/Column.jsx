import React from 'react';
import PropTypes, { objectOf } from 'prop-types';

import { Button, Card, Dropdown, Grid, Icon, Input } from 'semantic-ui-react';
import { Droppable } from 'react-beautiful-dnd';

import Task from './Task';

import { updateTask, createTask, deleteTask } from '../util/board_functions';

function OptionsMenu(props) {
    const { deleteColumn } = props;

    return (
        <Dropdown compact icon="vertical ellipsis" className="icon">
            <Dropdown.Menu>
                <Dropdown.Header content="Actions" />
                <Dropdown.Divider />
                <Dropdown.Item icon="trash" onClick={deleteColumn}>
                    Delete
                </Dropdown.Item>
            </Dropdown.Menu>
        </Dropdown>
    );
}

OptionsMenu.propTypes = {
    deleteColumn: PropTypes.func.isRequired,
};

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

        createTask(gid, formdata).then((data) => {
            updateColumn({
                gid,
                title,
                tasks: tasks.concat([data]),
            });
        });
    };

    updateTaskUI = (task) => {
        const { gid, title, tasks, updateColumn } = this.props;
        const index = tasks.findIndex((t) => t.gid === task.gid);
        tasks[index] = task;
        updateColumn({ gid, title, tasks });
    };

    removeTask = (tGid) => {
        const { gid, title, tasks, updateColumn } = this.props;
        tasks.splice(
            tasks.findIndex((t) => t.gid === tGid),
            1
        );

        updateColumn({
            gid,
            title,
            tasks,
        });
    };

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
                () => updateTask(gid, task.gid, formdata),
                500
            );

            this.setState({ taskTimers });
        }

        this.updateTaskUI(task); // Updates UI
    };

    handleTaskRemove = (task) => {
        const { gid } = this.props;

        deleteTask(gid, task.gid).then(() => {
            this.removeTask(task.gid);
        });
    };

    handleTitle = (value) => {
        const { gid, tasks, updateColumn } = this.props;
        updateColumn({ gid, title: value, tasks });
    };

    render() {
        const { gid, title, tasks, deleteColumn } = this.props;

        return (
            <Grid.Column style={{ marginBottom: '2rem' }}>
                <Card>
                    <Card.Content
                        style={{
                            paddingBottom: tasks.length ? '0rem' : '7.5rem',
                        }}
                    >
                        <Card.Header>
                            <Grid columns="2">
                                <Grid.Row>
                                    <Grid.Column width="13">
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
                                    </Grid.Column>
                                    <Grid.Column width="3" textAlign="center">
                                        <OptionsMenu
                                            deleteColumn={() =>
                                                deleteColumn({ gid })
                                            }
                                        />
                                    </Grid.Column>
                                </Grid.Row>
                            </Grid>
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
                            <Icon name="small plus" />
                            <span>Add Task</span>
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
                                            deleteTask={this.handleTaskRemove}
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
    deleteColumn: PropTypes.func.isRequired,
};

Column.defaultProps = {
    title: '',
    tasks: [],
};

export default Column;
