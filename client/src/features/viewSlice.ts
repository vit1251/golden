
import { createAction, createSlice } from '@reduxjs/toolkit';

export type EchoMsgViewPayload = {
    area: {
        name: string,              // Имя эхоконференции
    },
    body: string,                  // Тело сообщения
    echo: {
        from: string,
        to: string,
        subject: string,
        date: string,
    },
};

export const echoMsgView = createAction<EchoMsgViewPayload>('ECHO_MSG_VIEW');

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
        builder.addCase(echoMsgView, (state, action) => {
            const { area, body, echo } = action.payload;
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
