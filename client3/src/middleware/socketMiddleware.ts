
import { type Middleware } from '@reduxjs/toolkit';
import { createAction } from '@reduxjs/toolkit';

export const socketConnect = createAction<void>('SOCKET_CONNECT');
export const socketSend = createAction<{ msg: any }>('SOCKET_SEND');
export const socketRecv = createAction<{ msg: any }>('SOCKET_RECV');

export const socketMiddleware: (url: string) => Middleware = (url: string) => {
    return (store) => {
        let socket: WebSocket | null = null;
        let socketReady: boolean = false;
        let queue: Array<any> = [];
        return (next) => (action) => {
            
            // 1. Установлено соединение
            if (socketConnect.match(action)) {
                console.log(`Устанавливаем соединение с "${url}"...`);
                socket = new WebSocket(url);
                socket.onopen = (event: Event) => {
                    console.log(`Соединение с "${url}" установлено.`);
                    // Шаг 1. Устанвливаем соединение
                    socketReady = true;
                    // Шаг 2. Отправляет отлорженные запросы
                    if (queue && socket) {
                        console.log(`Отправка очереди запросов`);
                        // Шаг 1. Перебираем все запросы
                        for (const msg of queue) {
                            console.log(`TX[Q]: msg = `, msg);
                            socket.send(JSON.stringify(msg));
                        }
                        // Шаг 2. Отчистка очереди
                        queue = [];
                    }
                    
                };
                socket.onclose = (event: Event) => {
                    console.log(`Соединение прервано`);
                    socketReady = false;
                    socket = null;
                };
                socket.onmessage = (event: MessageEvent<any>) => {
                    const msg: unknown = JSON.parse(event.data);
                    console.log(`RX[D]: msg = `, msg);
                    store.dispatch(socketRecv({
                        msg: msg,
                    }));
                };
            }

            // 2. Отправка запроса на сервер
            if (socketSend.match(action)) {
                // Шаг 1. Достаем статус соеднинения
                const { app } = store.getState();
                console.log(app);
                // Шаг 2. Соединение установлено
                if (socketReady && socket) {
                    const { msg } = action.payload;
                    console.log(`TX[D]: msg = `, msg);
                    socket.send(JSON.stringify(msg));
                } else {
                    console.log(`Соединение не установлено. Кладем в очередь запросов.`);
                    const { msg } = action.payload;
                    queue.push(msg);
                }
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
