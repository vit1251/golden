
import { configureStore } from '@reduxjs/toolkit';
import { socketMiddleware } from '../middleware/socketMiddleware';

import identityReducer from '../features/identitySlice';
import areaReducer from '../features/areaSlice';
import mailerReducer from '../features/mailerSlice';
import messageReducer from '../features/messageSlice';
import viewReducer from '../features/viewSlice';
import composeReducer from '../features/composeSlice';

const rootReducer = {
    identity: identityReducer,
    areas: areaReducer,
    mailer: mailerReducer,
    messages: messageReducer,
    view: viewReducer,
    compose: composeReducer,
};

export const store = configureStore({
    middleware: (getDefaultMiddleware) => getDefaultMiddleware().concat(socketMiddleware("ws://127.0.0.1:8081/api/v1")),
    reducer: rootReducer,
});

// Сразу открываем соединение с сервером
store.dispatch({ type: 'SOCKET_CONNECT' });
