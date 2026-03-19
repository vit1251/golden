
import { createSlice } from '@reduxjs/toolkit';

export const mailerSlice = createSlice({
    name: 'mailer',
    initialState: {
        value: 0,
    },
    reducers: {
        increment: (state) => { state.value += 1; },
        decrement: (state) => { state.value -= 1; },
    },
});

export const { increment, decrement } = mailerSlice.actions;

export default mailerSlice.reducer;
