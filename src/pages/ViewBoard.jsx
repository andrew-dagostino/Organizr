import React from 'react';
import { Grid } from 'semantic-ui-react';
import Column from '../components/Column';

export default class ViewBoard extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            columns: 4,
        };
    }

    render() {
        const { columns } = this.state;

        return (
            <Grid columns={columns} container stackable doubling>
                <Grid.Row style={{ height: '100%' }}>
                    <Column />
                </Grid.Row>
            </Grid>
        );
    }
}
