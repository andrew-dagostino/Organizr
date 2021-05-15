import React from 'react';
import { Button, Card, Grid, Icon, Input } from 'semantic-ui-react';
import { v4 as uuidv4 } from 'uuid';
import { DragDropContext, Droppable } from 'react-beautiful-dnd';
import Task from './Task';

const reorder = (list, startIndex, endIndex) => {
    const result = Array.from(list);
    const [removed] = result.splice(startIndex, 1);
    result.splice(endIndex, 0, removed);

    return result;
};

export default class Column extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            tasks: [],
        };
    }

    onDragEnd = (result) => {
        // dropped outside the list
        if (!result.destination) {
            return;
        }

        const { tasks } = this.state;
        this.setState({
            tasks: reorder(
                tasks,
                result.source.index,
                result.destination.index
            ),
        });
    };

    addTask = (task) => {
        const { tasks } = this.state;

        if (!task) {
            tasks.push({
                id: uuidv4(),
                title: '',
                description: '',
            });
        } else {
            tasks.push(task);
        }

        this.setState({ tasks });
    };

    render() {
        const { tasks } = this.state;

        return (
            <Grid.Column>
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
                        <DragDropContext onDragEnd={this.onDragEnd}>
                            <Droppable droppableId="droppable">
                                {(provided) => (
                                    <div
                                        // eslint-disable-next-line react/jsx-props-no-spreading
                                        {...provided.droppableProps}
                                        ref={provided.innerRef}
                                    >
                                        {tasks.map((task, index) => (
                                            <Task
                                                key={task.id}
                                                id={task.id}
                                                title={task.title}
                                                description={task.description}
                                                index={index}
                                            />
                                        ))}
                                        {provided.placeholder}
                                    </div>
                                )}
                            </Droppable>
                        </DragDropContext>
                    </Card.Content>
                </Card>
            </Grid.Column>
        );
    }
}
