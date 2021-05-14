import React from 'react';
import { Grid } from 'semantic-ui-react';

export default class ViewBoard extends React.Component {
    constructor(props) {
        super(props);

        this.state = {};
    }

    render() {
        return (
            <Grid container stackable doubling>
                <Grid.Row>
                    <Grid.Column>
                        <div>test</div>
                    </Grid.Column>
                </Grid.Row>
            </Grid>
        );
    }
}
