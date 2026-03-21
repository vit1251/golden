
import { createSlice } from '@reduxjs/toolkit';

export const areaSlice = createSlice({
    name: 'area',
    initialState: {
        records: [],
        activeIndex: 0,
    },
    reducers: {
        prevArea: (state) => {
            console.log(`prev`);
            state.activeIndex = state.activeIndex - 1 >= 0 ? state.activeIndex - 1 : state.activeIndex;
            console.log(state.activeIndex);
        },
        nextArea: (state) => {
            console.log(`next`);
            state.activeIndex = state.activeIndex + 1 < state.records.length ? state.activeIndex + 1 : state.activeIndex;
            console.log(state.records.length);
            console.log(state.activeIndex);
        },
    },
    extraReducers: (builder) => {
        builder.addCase('ECHO_INDEX', (state, action) => {
            const { areas = [] } = action;
            state.records = areas ?? [];
        });
    },
});

export const { nextArea, prevArea } = areaSlice.actions;

export default areaSlice.reducer;
