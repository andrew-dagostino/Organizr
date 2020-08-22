import React from 'react';

import axios from 'axios';

import { Row, Col } from 'react-bootstrap';

import Post from 'src/components/Post.jsx';

export default class PostTimeline extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            jwt: props.jwt,
            posts: [],
        };

        this.getPosts = this.getPosts.bind(this);
        this.renderPost = this.renderPost.bind(this);
    }

    renderPost(title, text, img) {
        return (
            <Col xs={12} lg={6} xl={4}>
                <Post title={title} text={text} img={img} />
            </Col>
        );
    }

    getPosts() {
        axios.get('/api/post', { headers: { Authorization: `Bearer ${this.state.jwt}` } }).then(
            (res) => this.setState({ posts: res.data }),
            (err) => toastr.error(err.response.data && err.response.data.error ? err.response.data.error : err.message),
        );
    }

    componentDidMount() {
        this.getPosts();
    }

    render() {
        return (
            <Row className='justify-content-center'>
                {this.state.posts.map((post) => this.renderPost(post.title, post.text, post.img))}
            </Row>
        );
    }
}
