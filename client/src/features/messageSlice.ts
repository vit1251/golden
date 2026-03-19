
import { createSlice } from '@reduxjs/toolkit';

export const messageSlice = createSlice({
    name: 'message',
    initialState: {
        records: [],
    },
    reducers: {
    },
    extraReducers: (builder) => {
        builder.addCase('ECHO_MSG_INDEX', (state, action) => {
            const { headers = [] } = action;
            state.records = headers ?? [];
        });
    },
});

//export const {  } = messageSlice.actions;

export default messageSlice.reducer;
