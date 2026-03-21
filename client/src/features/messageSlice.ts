
import { createSlice } from '@reduxjs/toolkit';

export const messageSlice = createSlice({
    name: 'message',
    initialState: {
        records: [],
        activeIndex: 0,
    },
    reducers: {
        firstMessage: (state) => {
            state.activeIndex = 0;
        },
        prevMessage: (state) => {
            state.activeIndex = state.activeIndex - 1 >= 0 ? state.activeIndex - 1 : state.activeIndex;
        },
        nextMessage: (state) => {
            state.activeIndex = state.activeIndex + 1 < state.records.length ? state.activeIndex + 1 : state.activeIndex;
        },
        lastMessage: (state) => {
            state.activeIndex = state.records.length - 1;
        }
    },
    extraReducers: (builder) => {
        builder.addCase('ECHO_MSG_INDEX', (state, action) => {
            const { headers = [] } = action;
            state.records = headers ?? [];
        });
    },
});

export const { firstMessage, prevMessage, nextMessage, lastMessage } = messageSlice.actions;

export default messageSlice.reducer;
