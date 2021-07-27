function rootReducer(state = {}, action) {
    switch (action.type) {
        case 'init':
            return action.payload.reduce((acc, cur) => {
                acc[cur.gid] = cur;
                return acc;
            }, {});

        case 'board/add':
            return {
                ...state,
                [action.payload.gid]: action.payload,
            };
        case 'board/update':
            return {
                ...state,
                [action.payload.gid]: action.payload,
            };
        case 'board/remove':
            return {
                ...state,
                [action.payload.gid]: null,
            };

        default:
            return state;
    }
}

export default rootReducer;
