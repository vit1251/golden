
import { createSlice } from '@reduxjs/toolkit';

export const viewSlice = createSlice({
    name: 'view',
    initialState: {
        echo: '',
        from: '',
        to: '',
        subject: '',
        date: '',
        headers: [],
        content: '',
    },
    reducers: {
//        removeMessage: (state) => { state.value += 1; },
    },
    extraReducers: (builder) => {
        builder.addCase('ECHO_MSG_VIEW', (state, action) => {
            const { area = {}, body, echo = {} } = action;
            state.echo = area.name ?? '';
            state.content = body;
            state.from = echo.from ?? '';
            state.to = echo.to ?? '';
            state.subject = echo.subject ?? '';
            state.date = echo.date ?? '';
        });
    },
});

//export const { removeMessage } = viewSlice.actions;

export default viewSlice.reducer;
