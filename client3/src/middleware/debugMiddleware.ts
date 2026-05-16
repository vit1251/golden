
import { type Middleware } from "redux";

export const debugMiddleware: Middleware = (store) => (next) => (action) => {

    // 1. Сначала пропускаем экшен дальше, чтобы состояние обновилось
    const result = next(action);

    // 2. Отзеркалим отладкой событие
    console.log(action);
    
    return result;
};
