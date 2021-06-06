import React from 'react';
import PropTypes from 'prop-types';
import { Card, Dropdown, Form, Grid, Icon, Input } from 'semantic-ui-react';
import { Draggable } from 'react-beautiful-dnd';

function OptionsMenu() {
    return (
        <Dropdown icon="vertical ellipsis" className="icon">
            <Dropdown.Menu>
                <Dropdown.Header content="Filter by tag" />
                <Dropdown.Divider />
                <Dropdown.Item>Important</Dropdown.Item>
                <Dropdown.Item>Announcement</Dropdown.Item>
                <Dropdown.Item>Discussion</Dropdown.Item>
            </Dropdown.Menu>
        </Dropdown>
    );
}

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
                                            name="grip lines"
                                            style={{
                                                width: '100%',
                                                cursor: 'pointer',
                                            }}
                                        />
                                        <Grid columns="2">
                                            <Grid.Row>
                                                <Grid.Column width="13">
                                                    <Input
                                                        placeholder="Task Name..."
                                                        fluid
                                                        transparent
                                                        size="mini"
                                                        value={title}
                                                        onChange={(e, data) =>
                                                            this.handleTitle(
                                                                data.value
                                                            )
                                                        }
                                                    />
                                                </Grid.Column>
                                                <Grid.Column
                                                    width="3"
                                                    textAlign="center"
                                                >
                                                    <OptionsMenu />
                                                </Grid.Column>
                                            </Grid.Row>
                                        </Grid>
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
