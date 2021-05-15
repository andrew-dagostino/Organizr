import React from 'react';
import { Card, Grid, Icon } from 'semantic-ui-react';
import { v4 as uuidv4 } from 'uuid';
import PropTypes from 'prop-types';
import { DragDropContext } from 'react-beautiful-dnd';
import Column from '../components/Column';

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
            columns: [],
        };
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
            this.updateColumn({
                id: sColumn.id,
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

            this.updateColumn({
                id: sColumn.id,
                title: sColumn.title,
                tasks: newResult[sId],
            });

            this.updateColumn({
                id: dColumn.id,
                title: dColumn.title,
                tasks: newResult[dId],
            });
        }
    };

    addColumn = () => {
        const { columns } = this.state;
        this.setState({ columns: columns.concat([{ id: uuidv4() }]) });
    };

    updateColumn = (column) => {
        const { columns } = this.state;
        const index = columns.findIndex((col) => col.id === column.id);
        columns[index] = column;
        this.setState({ columns });
    };

    getColumn = (id) => {
        const { columns } = this.state;
        return columns.filter((column) => column.id === id)[0];
    };

    render() {
        const { columns } = this.state;
        return (
            <DragDropContext onDragEnd={this.onDragEnd}>
                <Grid columns="4" container doubling stackable={false}>
                    <Grid.Row style={{ height: '100%' }}>
                        {columns.map((column) => (
                            <Column
                                key={column.id}
                                id={column.id}
                                title={column.title}
                                tasks={column.tasks}
                                updateColumn={this.updateColumn}
                                getColumn={this.getColumn}
                            />
                        ))}
                        <AddColumnWidget onClick={() => this.addColumn()} />
                    </Grid.Row>
                </Grid>
            </DragDropContext>
        );
    }
}
