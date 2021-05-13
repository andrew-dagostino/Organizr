import React, { Component } from 'react';
import { Menu } from 'semantic-ui-react';

export default class Header extends Component {
    constructor(props) {
        super(props);

        this.state = {
            activeItem: 'boards',
        };
    }

    handleItemClick = (e, { name }) => this.setState({ activeItem: name });

    render() {
        const { activeItem } = this.state;

        return (
            <Menu>
                <Menu.Item header>Organizr</Menu.Item>
                <Menu.Item
                    name="about"
                    active={activeItem === 'about'}
                    onClick={this.handleItemClick}
                />
                <Menu.Item
                    name="boards"
                    active={activeItem === 'boards'}
                    onClick={this.handleItemClick}
                />
                <Menu.Item
                    name="addBoard"
                    active={activeItem === 'addBoard'}
                    onClick={this.handleItemClick}
                />
            </Menu>
        );
    }
}
