import React from 'react';
import PropTypes from 'prop-types';
import { Card, Form, Icon, Input } from 'semantic-ui-react';
import { Draggable } from 'react-beautiful-dnd';

class Task extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            id: props.id,
            title: props.title,
            description: props.description,
        };
    }

    render() {
        const { id, title, description } = this.state;
        const { index } = this.props;

        return (
            <div>
                <Draggable draggableId={id} index={index}>
                    {(provided) => (
                        <div
                            ref={provided.innerRef}
                            // eslint-disable-next-line react/jsx-props-no-spreading
                            {...provided.draggableProps}
                            // eslint-disable-next-line react/jsx-props-no-spreading
                            {...provided.dragHandleProps}
                        >
                            <Card>
                                <Card.Content>
                                    <Card.Header>
                                        <Icon
                                            name="drag handle"
                                            style={{
                                                width: '100%',
                                                cursor: 'pointer',
                                            }}
                                        />
                                        <Input
                                            placeholder="Task Name..."
                                            fluid
                                            transparent
                                            size="mini"
                                            value={title}
                                            onChange={(e, data) =>
                                                this.setState({
                                                    title: data.value,
                                                })
                                            }
                                        />
                                    </Card.Header>
                                    <hr />
                                    <Form>
                                        <Form.TextArea
                                            style={{ resize: 'none' }}
                                            size="tiny"
                                            rows="6"
                                            placeholder="Description..."
                                            value={description}
                                            onChange={(e, data) => {
                                                this.setState({
                                                    description: data.value,
                                                });
                                            }}
                                        />
                                    </Form>
                                </Card.Content>
                            </Card>
                        </div>
                    )}
                </Draggable>
                <br />
            </div>
        );
    }
}

Task.propTypes = {
    id: PropTypes.string.isRequired,
    title: PropTypes.string.isRequired,
    description: PropTypes.string.isRequired,
    index: PropTypes.number.isRequired,
};

export default Task;
