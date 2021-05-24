import React from 'react';
import { Grid } from 'semantic-ui-react';
import LoginRegisterCard from '../components/LoginRegisterCard';

export default function SplashPage() {
    return (
        <Grid columns="3" doubling style={{ height: '80vh' }}>
            <Grid.Row centered>
                <Grid.Column verticalAlign="middle">
                    <LoginRegisterCard />
                </Grid.Column>
            </Grid.Row>
        </Grid>
    );
}
