
import { createStore } from "redux";

const initialState = {
    summary: {},
    areas: [],
};

const reducer = (state = initialState, action) => {
    if (action.type === 'ECHO_INDEX') {
        const { areas = [] } = action;
        return {
            ...state,
            areas,
        };
    } else if (action.type === 'SUMMARY') {
        const {
            NetMessageCount = 0,
            EchoMessageCount = 0,
            FileCount = 0,
        } = action;
        return {
            ...state,
            summary: {
                NetMessageCount,
                EchoMessageCount,
                FileCount,
            },
        };
    } else {
        return state;
    }
};

export const store = createStore(reducer);
