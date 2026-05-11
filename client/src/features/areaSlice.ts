
import { createAction, createSlice } from '@reduxjs/toolkit';
import { type Area } from '../models/Area.model.ts';

export type EchoIndexPayload = {
    areas: Array<Area>,
};

export const echoIndex = createAction<EchoIndexPayload>('ECHO_INDEX');

export const areaSlice = createSlice({
    name: 'area',
    initialState: {
        records: [] as Array<Area>,
        activeIndex: 0,
        currentPage: 1,
        pageSize: 10,
    },
    reducers: {
        /* Переход между разделами */
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
        /* Страницы */
        setPage: (state, action) => {
            const page = action.payload;
            const maxPage = Math.ceil(state.records.length / state.pageSize);
            state.currentPage = page > 0 ? (page <= maxPage ? page : maxPage) : 1;
        },
        setPageSize: (state, action) => {
            state.pageSize = action.payload;
            state.currentPage = 1; // Сброс на первую страницу
        },
    },
    extraReducers: (builder) => {
        builder.addCase(echoIndex, (state, action) => {
            const { areas } = action.payload;
            // 1. Обновляем записи
            state.records = areas ?? [];
            // 2. Корректируем
            state.activeIndex = state.activeIndex >= state.records.length ? state.records.length - 1 : state.activeIndex;
        });
    },
});

export const { nextArea, prevArea } = areaSlice.actions;
export const { setPage, setPageSize } = areaSlice.actions;

export default areaSlice.reducer;
