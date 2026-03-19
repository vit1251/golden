
import { createSlice } from '@reduxjs/toolkit';

export const identitySlice = createSlice({
    name: 'identity',
    initialState: {
        value: 0,
    },
    reducers: {
        increment: (state) => { state.value += 1; },
        decrement: (state) => { state.value -= 1; },
    },
});

export const { increment, decrement } = identitySlice.actions;

export default identitySlice.reducer;
