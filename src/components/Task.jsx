import React from 'react';
import PropTypes from 'prop-types';
import { Card, Form, Icon, Input } from 'semantic-ui-react';
import { Draggable } from 'react-beautiful-dnd';

class Task extends React.Component {
    handleTitle = (val) => {
        const { gid, description, updateTask } = this.props;
        updateTask({ gid, title: val, description });
    };

    handleDescription = (val) => {
        const { gid, title, updateTask } = this.props;
        updateTask({ gid, title, description: val });
    };

    render() {
        const { gid, title, description, index } = this.props;

        return (
            <div>
                <Draggable draggableId={gid} index={index}>
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
                                                this.handleTitle(data.value)
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
                                            onChange={(e, data) =>
                                                this.handleDescription(
                                                    data.value
                                                )
                                            }
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
    gid: PropTypes.string.isRequired,
    title: PropTypes.string,
    description: PropTypes.string,
    index: PropTypes.number.isRequired,
    updateTask: PropTypes.func.isRequired,
};

Task.defaultProps = {
    title: '',
    description: '',
};

export default Task;
