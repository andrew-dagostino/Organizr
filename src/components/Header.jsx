import React, { Component } from 'react';
import PropTypes from 'prop-types';
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
        const { title, handleChange, showTextfield } = this.props;

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
                    {showTextfield ? (
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
                            onChange={handleChange}
                        >
                            <input
                                style={{ textAlign: 'center' }}
                                value={title}
                            />
                        </Input>
                    ) : (
                        ''
                    )}
                    <Menu.Item position="right" onClick={this.handleLogOut}>
                        Log out
                    </Menu.Item>
                </Menu>
            </Sticky>
        );
    }
}

Header.propTypes = {
    title: PropTypes.string,
    handleChange: PropTypes.func.isRequired,
    showTextfield: PropTypes.bool,
};

Header.defaultProps = {
    title: '',
    showTextfield: false,
};
