import React, { Component } from 'react';
import { Menu } from 'semantic-ui-react';

import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';

import ViewBoards from '../pages/ViewBoards';
import About from '../pages/About';
import AddBoard from '../pages/AddBoard';

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
                path: '/board/new',
                children: () => <AddBoard />,
            },
        ];
    }

    handleItemClick = (e, { name }) => this.setState({ activeItem: name });

    render() {
        const { activeItem } = this.state;

        return (
            <Router>
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
                    <Link to="/board/new">
                        <Menu.Item
                            name="addBoard"
                            active={activeItem === 'addBoard'}
                            onClick={this.handleItemClick}
                        />
                    </Link>
                </Menu>

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
