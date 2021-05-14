import React from 'react';
import { Card, Form, Grid } from 'semantic-ui-react';

export default class AddBoard extends React.Component {
    constructor(props) {
        super(props);

        this.state = {};
    }

    render() {
        return (
            <Grid doubling stackable container>
                <Grid.Row centered>
                    <Grid.Column width="6">
                        <Card fluid>
                            <Card.Content>
                                <Card.Header style={{ marginBottom: '1rem' }}>
                                    Add Board
                                </Card.Header>
                                <Form>
                                    <Form.Group widths="equal">
                                        <Form.Input
                                            label="Board Name"
                                            type="text"
                                            required
                                            fluid
                                        />
                                    </Form.Group>
                                    <Form.Group widths="equal">
                                        <Form.Input
                                            label="Board Image (Optional)"
                                            type="file"
                                            fluid
                                        />
                                    </Form.Group>
                                    <Form.Button
                                        type="submit"
                                        size="medium"
                                        primary
                                    >
                                        Create
                                    </Form.Button>
                                </Form>
                            </Card.Content>
                        </Card>
                    </Grid.Column>
                </Grid.Row>
            </Grid>
        );
    }
}
