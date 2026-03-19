
export const socketMiddleware = (url: string) => {
    return (store) => {
        let socket = null;
        return (next) => (action) => {
            // 0. Отладочное сообщение
            console.log('событие по шине = ', action);
            // 1. Инициализация сокета по специальному экшену
            if (action.type === 'SOCKET_CONNECT') {
                socket = new WebSocket(url);
                socket.onmessage = (event) => {
                    const data = JSON.parse(event.data);
                    store.dispatch(data);
                };
            }
            // 2. Возможность отправить сообщение на сервер через сокет
            if (action.type === 'SOCKET_SEND' && socket) {
                socket.send(JSON.stringify(action.payload));
            }
            return next(action);
        };
    };
};
