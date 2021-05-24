import React, { Component } from 'react';

import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';

import ViewBoard from './pages/ViewBoard';
import ViewBoards from './pages/ViewBoards';
import SplashPage from './pages/SplashPage';

export default class App extends Component {
    constructor(props) {
        super(props);

        this.routes = [
            {
                id: 0,
                path: '/',
                exact: true,
                children: () => <SplashPage />,
            },
            {
                id: 1,
                path: '/board',
                exact: true,
                children: () => <ViewBoards />,
            },
            {
                id: 2,
                path: '/board/:id',
                children: ({ id }) => <ViewBoard id={id} />,
            },
        ];
    }

    render() {
        return (
            <Router>
                <Switch>
                    {this.routes.map((route) => (
                        <Route
                            key={route.id}
                            path={route.path}
                            exact={route.exact}
                        >
                            {route.children}
                        </Route>
                    ))}
                </Switch>
            </Router>
        );
    }
}
