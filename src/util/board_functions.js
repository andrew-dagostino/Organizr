import axios from 'axios';
import config from '../config.json';

const JWT = window.localStorage.getItem('jwt');

function retrieveBoard(gid) {
    return axios.get(`${config.API_URL}/board/${gid}`, {
        headers: {
            Authorization: `Bearer ${JWT}`,
        },
    });
}

function retrieveBoards() {
    return axios.get(`${config.API_URL}/board`, {
        headers: {
            Authorization: `Bearer ${JWT}`,
        },
    });
}

function updateBoard(gid, formdata) {
    return axios.put(`${config.API_URL}/board/${gid}`, formdata, {
        headers: {
            'Content-Type': 'multipart/form-data',
            Authorization: `Bearer ${JWT}`,
        },
    });
}

function createBoard(formdata) {
    axios.post(`${config.API_URL}/board`, formdata, {
        headers: {
            'Content-Type': 'multipart/form-data',
            Authorization: `Bearer ${JWT}`,
        },
    });
}

function retrieveColumns(gid) {
    return axios.get(`${config.API_URL}/column/${gid}`, {
        headers: {
            Authorization: `Bearer ${JWT}`,
        },
    });
}

function updateColumn(bGid, cGid, formdata) {
    return axios.put(`${config.API_URL}/column/${bGid}/${cGid}`, formdata, {
        headers: {
            'Content-Type': 'multipart/form-data',
            Authorization: `Bearer ${JWT}`,
        },
    });
}

function createColumn(cGid, formdata) {
    return axios.post(`${config.API_URL}/column/${cGid}`, formdata, {
        headers: {
            'Content-Type': 'multipart/form-data',
            Authorization: `Bearer ${JWT}`,
        },
    });
}

function deleteColumn(bGid, cGid) {
    return axios.delete(`${config.API_URL}/column/${bGid}/${cGid}`, {
        headers: {
            'Content-Type': 'multipart/form-data',
            Authorization: `Bearer ${JWT}`,
        },
    });
}

function retrieveTasks(cGid) {
    return axios.get(`${config.API_URL}/task/${cGid}`, {
        headers: {
            Authorization: `Bearer ${JWT}`,
        },
    });
}

function updateTask(cGid, tGid, formdata) {
    return axios.put(`${config.API_URL}/task/${cGid}/${tGid}`, formdata, {
        headers: {
            'Content-Type': 'multipart/form-data',
            Authorization: `Bearer ${JWT}`,
        },
    });
}

function createTask(cGid, formdata) {
    return axios.post(`${config.API_URL}/task/${cGid}`, formdata, {
        headers: {
            'Content-Type': 'multipart/form-data',
            Authorization: `Bearer ${JWT}`,
        },
    });
}

function deleteTask(cGid, tGid) {
    return axios.delete(`${config.API_URL}/task/${cGid}/${tGid}`, {
        headers: {
            'Content-Type': 'multipart/form-data',
            Authorization: `Bearer ${JWT}`,
        },
    });
}

export {
    retrieveBoard,
    retrieveBoards,
    updateBoard,
    createBoard,
    retrieveColumns,
    updateColumn,
    createColumn,
    deleteColumn,
    retrieveTasks,
    updateTask,
    createTask,
    deleteTask,
};
