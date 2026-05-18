

export class RPC {
    readonly requests: Record<string, { timerId: number, callback: (err: Error | null, msg?: unknown) => void }> = {};
    protected requestTimeoutMs: number = 5_000; // Время ожидания ответа (5 сек.)

    constructor(readonly socket: WebSocket) {
        socket.onmessage = (event: MessageEvent<any>) => {
            const msg: any = JSON.parse(event.data);
            const { id } = msg;
            this.handleResponse(id, msg);
        };
    }

    handleResponse(requestId: string, msg: unknown) {
        const { callback, timerId } = this.requests[requestId] ?? {};
        // Шаг 1. Отменяем таймер таймаута
        clearTimeout(timerId);
        // Шаг 2. Обрабатываем задачу
        if (callback) {
            callback(null, msg);
        }
        // Шаг 3. Выкидываем запись
        delete this.requests[requestId];
    }

    fetch(payload: any, callback: (err: null | Error) => void) {
        const requestId: string = crypto.randomUUID();;
        const timerId: number = setTimeout(() => {
            // Шаг 1. Бросаем ошибку
            const { callback } = this.requests[requestId] ?? {};
            if (callback) {
                callback(new Error('Tiemout'));
            }
            // Шаг 2. Удаляем запись
            delete this.requests[requestId];
        }, this.requestTimeoutMs);
        this.requests[requestId] = {
            callback: callback,
            timerId: timerId,
        };
        const packet = {
            ...payload,
            id: requestId,
        }
        this.socket.send(JSON.stringify(packet));
    }

}
