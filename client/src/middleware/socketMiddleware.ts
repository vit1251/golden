
import { type Middleware } from '@reduxjs/toolkit';
import { createAction } from '@reduxjs/toolkit';

export const socketConnect = createAction<unknown>('SOCKET_CONNECT');
export const socketSend = createAction<{ msg: any }>('SOCKET_SEND');
export const socketRecv = createAction<{ msg: any }>('SOCKET_RECV');

export const socketMiddleware: (url: string) => Middleware = (url: string) => {
    return (store) => {
        let socket: WebSocket | null = null;
        return (next) => (action) => {
            
            // 1. Установлено соединение
            if (socketConnect.match(action)) {
                socket = new WebSocket(url);
                socket.onmessage = (event: MessageEvent<any>) => {
                    const msg: unknown = JSON.parse(event.data);
                    store.dispatch(socketRecv({
                        msg: msg,
                    }));
                };
            }

            // 2. Отправка запроса на сервер
            if (socket && socketSend.match(action)) {
                const { msg } = action.payload;
                socket.send(JSON.stringify(msg));
            }

            // 3. Получение ответа от сервера
            if (socketRecv.match(action)) {
                const { msg } =  action.payload;
                const { type: actionType, ...data } = msg as Record<string, any>;
                store.dispatch({
                    type: actionType,
                    payload: data,
                });
            }

            return next(action);
        };
    };
};
