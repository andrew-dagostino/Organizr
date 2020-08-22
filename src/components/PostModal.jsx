import React from 'react';

import axios from 'axios';

import Modal from 'react-bootstrap/Modal';
import Form from 'react-bootstrap/Form';
import { Button } from 'react-bootstrap';

export default class PostModal extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            jwt: props.jwt,
            show: false,
        };
    }

    hideModal = () => this.props.showModal(false);

    submitPost = (e) => {
        e.preventDefault();

        let formdata = new FormData(document.getElementById('postForm'));
        axios.post('/api/post', formdata, { headers: { Authorization: `Bearer ${this.state.jwt}` } }).then(
            (res) => {
                toastr.success('Post Created Successfully');
                this.hideModal();
            },
            (err) => toastr.error(err.response.data && err.response.data.error ? err.response.data.error : err.message),
        );
    };

    componentDidUpdate(oldProps) {
        if (oldProps.show !== this.props.show) {
            this.setState({ show: this.props.show });
        }
    }

    render() {
        return (
            <Modal show={this.props.show}>
                <Modal.Header closeButton>
                    <Modal.Title>New Post</Modal.Title>
                </Modal.Header>

                <Modal.Body>
                    <Form id='postForm'>
                        <Form.Group>
                            <Form.Label>Title</Form.Label>
                            <Form.Control type='text' required={true} name='title' />
                            <Form.Text className='text-muted'>Enter the title for your post</Form.Text>
                        </Form.Group>
                        <Form.Group>
                            <Form.Label>Text</Form.Label>
                            <Form.Control as='textarea' rows='5' name='text' />
                            <Form.Text className='text-muted'>Enter the text content for your post</Form.Text>
                        </Form.Group>
                    </Form>
                </Modal.Body>

                <Modal.Footer>
                    <Button type='button' variant='secondary' onClick={this.hideModal}>
                        Cancel
                    </Button>
                    <Button type='submit' variant='primary' onClick={this.submitPost}>
                        Save
                    </Button>
                </Modal.Footer>
            </Modal>
        );
    }
}
