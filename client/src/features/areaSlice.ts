
import { createSlice } from '@reduxjs/toolkit';

export const areaSlice = createSlice({
    name: 'area',
    initialState: {
        records: [],
    },
    reducers: {
    },
    extraReducers: (builder) => {
        builder.addCase('ECHO_INDEX', (state, action) => {
            const { areas = [] } = action;
            state.records = areas ?? [];
        });
    },
});

//export const { } = areaSlice.actions;

export default areaSlice.reducer;
