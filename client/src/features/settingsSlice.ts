
import { createSlice, PayloadAction } from '@reduxjs/toolkit';

export const settingsSlice = createSlice({
    name: 'settings',
    initialState: {
        code: 'en-US',
    },
    reducers: {
        setCode: (state, action: PayloadAction<string>) => {
            state.code = action.payload;
        },
    },
});

export const { setCode } = settingsSlice.actions;

export default settingsSlice.reducer;
