import React, { Component } from 'react';
import { Input, Menu, Sticky } from 'semantic-ui-react';

import { Link } from 'react-router-dom';

export default class Header extends Component {
    constructor(props) {
        super(props);

        this.state = {
            activeItem: '',
        };
    }

    handleItemClick = (e, { name }) => {
        this.setState({ activeItem: name });
    };

    handleLogOut = () => {
        window.localStorage.removeItem('jwt');
        window.location.replace('/');
    };

    render() {
        const { activeItem } = this.state;

        return (
            <Sticky style={{ marginBottom: '2rem' }}>
                <Menu>
                    <Menu.Item header>Organizr</Menu.Item>
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
                            left: '44%',
                            height: '3rem',
                        }}
                    >
                        <input style={{ textAlign: 'center' }} />
                    </Input>
                    <Menu.Item position="right" onClick={this.handleLogOut}>
                        Log out
                    </Menu.Item>
                </Menu>
            </Sticky>
        );
    }
}
