import React from 'react';
import { Button, Card, Grid, Icon, Input } from 'semantic-ui-react';
import { v4 as uuidv4 } from 'uuid';
import { Droppable } from 'react-beautiful-dnd';
import PropTypes, { objectOf } from 'prop-types';
import Task from './Task';

class Column extends React.Component {
    addTask = () => {
        const { gid, title, tasks, updateColumn } = this.props;
        updateColumn({
            gid,
            title,
            tasks: tasks.concat([
                {
                    gid: uuidv4(),
                },
            ]),
        });
    };

    updateTask = (task) => {
        const { gid, title, tasks, updateColumn } = this.props;
        const index = tasks.findIndex((t) => t.gid === task.gid);
        tasks[index] = task;
        updateColumn({ gid, title, tasks });
    };

    handleTitle = (value) => {
        const { gid, tasks, updateColumn } = this.props;
        updateColumn({ gid, title: value, tasks }, value);
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
                                            id={task.gid}
                                            title={task.title}
                                            description={task.description}
                                            index={index}
                                            updateTask={this.updateTask}
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
