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
    return axios.post(`${config.API_URL}/board`, formdata, {
        headers: {
            'Content-Type': 'multipart/form-data',
            Authorization: `Bearer ${JWT}`,
        },
    });
}

function deleteBoard(bGid) {
    return axios.delete(`${config.API_URL}/board/${bGid}`, {
        headers: {
            'Content-Type': 'multipart/form-data',
            Authorization: `Bearer ${JWT}`,
        },
    });
}

function retrieveColumns(bGid) {
    return axios.get(`${config.API_URL}/column?board_gid=${bGid}`, {
        headers: {
            Authorization: `Bearer ${JWT}`,
        },
    });
}

function updateColumn(cGid, formdata) {
    return axios.put(`${config.API_URL}/column/${cGid}`, formdata, {
        headers: {
            'Content-Type': 'multipart/form-data',
            Authorization: `Bearer ${JWT}`,
        },
    });
}

function createColumn(formdata) {
    return axios.post(`${config.API_URL}/column`, formdata, {
        headers: {
            'Content-Type': 'multipart/form-data',
            Authorization: `Bearer ${JWT}`,
        },
    });
}

function deleteColumn(cGid) {
    return axios.delete(`${config.API_URL}/column/${cGid}`, {
        headers: {
            'Content-Type': 'multipart/form-data',
            Authorization: `Bearer ${JWT}`,
        },
    });
}

function retrieveTasks(cGid) {
    return axios.get(`${config.API_URL}/task?column_gid=${cGid}`, {
        headers: {
            Authorization: `Bearer ${JWT}`,
        },
    });
}

function updateTask(tGid, formdata) {
    return axios.put(`${config.API_URL}/task/${tGid}`, formdata, {
        headers: {
            'Content-Type': 'multipart/form-data',
            Authorization: `Bearer ${JWT}`,
        },
    });
}

function createTask(formdata) {
    return axios.post(`${config.API_URL}/task`, formdata, {
        headers: {
            'Content-Type': 'multipart/form-data',
            Authorization: `Bearer ${JWT}`,
        },
    });
}

function deleteTask(tGid) {
    return axios.delete(`${config.API_URL}/task/${tGid}`, {
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
    deleteBoard,
    retrieveColumns,
    updateColumn,
    createColumn,
    deleteColumn,
    retrieveTasks,
    updateTask,
    createTask,
    deleteTask,
};
