
import { createSlice } from '@reduxjs/toolkit';

export const composeSlice = createSlice({
    name: 'compose',
    initialState: {
        headers: [],
        content: '',
    },
    reducers: {
//        removeMessage: (state) => { state.value += 1; },
    },
    extraReducers: (builder) => {
//        builder.addCase('ECHO_MSG_INDEX', (state, action) => {
//            const { headers = [] } = action;
//            state.records = headers ?? [];
//        });
    },
});

//export const { removeMessage } = viewSlice.actions;

export default composeSlice.reducer;