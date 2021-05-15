import React from 'react';
import { Card, Grid, Icon } from 'semantic-ui-react';
import { v4 as uuidv4 } from 'uuid';
import PropTypes from 'prop-types';
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

export default class ViewBoard extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            columns: [],
        };
    }

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

    render() {
        const { columns } = this.state;

        return (
            <Grid columns="4" container doubling stackable={false}>
                <Grid.Row style={{ height: '100%' }}>
                    {columns.map((column) => (
                        <Column
                            key={column.id}
                            id={column.id}
                            title={column.title}
                            tasks={column.tasks}
                            updateColumn={this.updateColumn}
                        />
                    ))}
                    <AddColumnWidget onClick={() => this.addColumn()} />
                </Grid.Row>
            </Grid>
        );
    }
}
