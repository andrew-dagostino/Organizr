import React, { Component } from 'react';
import { Input, Menu, Sticky } from 'semantic-ui-react';

import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';

import ViewBoard from '../pages/ViewBoard';
import ViewBoards from '../pages/ViewBoards';
import About from '../pages/About';

export default class Header extends Component {
    constructor(props) {
        super(props);

        this.state = {
            activeItem: 'boards',
        };

        this.routes = [
            {
                id: 0,
                path: ['/', '/board'],
                exact: true,
                children: () => <ViewBoards />,
            },
            {
                id: 1,
                path: '/about',
                children: () => <About />,
            },
            {
                id: 2,
                path: '/board/:id',
                children: ({ id }) => <ViewBoard id={id} />,
            },
        ];
    }

    handleItemClick = (e, { name }) => this.setState({ activeItem: name });

    render() {
        const { activeItem } = this.state;

        return (
            <Router>
                <Sticky style={{ marginBottom: '2rem' }}>
                    <Menu>
                        <Menu.Item header>Organizr</Menu.Item>
                        <Link to="/about">
                            <Menu.Item
                                name="about"
                                active={activeItem === 'about'}
                                onClick={this.handleItemClick}
                            />
                        </Link>
                        <Link to="/board">
                            <Menu.Item
                                name="boards"
                                active={activeItem === 'boards'}
                                onClick={this.handleItemClick}
                            />
                        </Link>
                        <Input
                            type="text"
                            transparent
                            size="huge"
                            placeholder="Board Name..."
                            style={{
                                position: 'absolute',
                                left: '40%',
                                height: '3rem',
                            }}
                        >
                            <input style={{ textAlign: 'center' }} />
                        </Input>
                    </Menu>
                </Sticky>

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
