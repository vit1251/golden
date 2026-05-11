
import { createAction, createSlice } from '@reduxjs/toolkit';
import { type Message } from '../models/Message.model.ts';

export const echoMsgIndex = createAction<{ headers: Array<Message> }>('ECHO_MSG_INDEX');

export const messageSlice = createSlice({
    name: 'message',
    initialState: {
        records: [] as Array<Message>,
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
        builder.addCase(echoMsgIndex, (state, action) => {
            const { headers } = action.payload;
            // 1. Обновим записи
            state.records = headers ?? [];
            // 2. Корректируем позицию курсора (ставим принудительно на последний элемент)
            state.activeIndex = state.activeIndex >= state.records.length ? state.records.length - 1 : state.activeIndex;
        });
    },
});

export const { firstMessage, prevMessage, nextMessage, lastMessage } = messageSlice.actions;

export default messageSlice.reducer;
