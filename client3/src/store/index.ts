
import { configureStore } from '@reduxjs/toolkit';

import appReducer from './appSlice.ts';

import { socketMiddleware } from '../middleware/socketMiddleware.ts';
import { soundMiddleware } from '../middleware/soundMiddleware.ts';
import { debugMiddleware } from '../middleware/debugMiddleware.ts';

export const store = configureStore({
    middleware: (getDefaultMiddleware) => getDefaultMiddleware()
        .concat(socketMiddleware("ws://127.0.0.1:8082/api/v1"))
        .concat(soundMiddleware)
        .concat(debugMiddleware),
    reducer: {
        app: appReducer,
//        mailer: mailerReducer,
//        tosser: tosserReducer,
    }
});

// Экспортируем типы глобального состояния и диспетчера
export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
