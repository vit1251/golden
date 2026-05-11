
import { configureStore } from '@reduxjs/toolkit';

import { socketConnect, socketMiddleware } from '../middleware/socketMiddleware.ts';
import { soundMiddleware } from '../middleware/soundMiddleware.ts';

import identityReducer from '../features/identitySlice.ts';
import areaReducer from '../features/areaSlice.ts';
import mailerReducer from '../features/mailerSlice.ts';
import messageReducer from '../features/messageSlice.ts';
import viewReducer from '../features/viewSlice.ts';
import composeReducer from '../features/composeSlice.ts';
import settingsSlice from '../features/settingsSlice.ts';
import { debugMiddleware } from '../middleware/debugMiddleware.ts';

const rootReducer = {
    identity: identityReducer,
    areas: areaReducer,
    mailer: mailerReducer,
    messages: messageReducer,
    view: viewReducer,
    compose: composeReducer,
    settings: settingsSlice,
};

export const store = configureStore({
    middleware: (getDefaultMiddleware) => getDefaultMiddleware()
        .concat(socketMiddleware("ws://127.0.0.1:8081/api/v1"))
        .concat(soundMiddleware)
        .concat(debugMiddleware),
    reducer: rootReducer,
});

// Сразу открываем соединение с сервером
store.dispatch(socketConnect());

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>

// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch