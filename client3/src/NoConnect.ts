
import { Color } from "./color.ts";
import { useEffect } from "./index.ts";
import { useKeyboard } from "./Keyboard.ts";
import type { Screen } from "./Screen.ts";

export const NoConnect = (screen: Screen) => {

    // Шаг 1. Занулить работу с горячими клавишами
    useEffect(() => {
        useKeyboard({});
    }, []);

    // Шаг 2. Заголовок
    screen.setForegroudColor(Color.Yellow);
    screen.writeText(4, 2, 'Ошибка связи');
    screen.writeText(4, 3, '============');

    // Шаг 3. Текст сообщения
    screen.setForegroudColor(Color.Gray);
    screen.writeText(4, 5, `Произошла внутренняя неустранимая ошибка соединения с приложением.`);
    screen.writeText(4, 7, `Перезагрузите сервер приложения.`);

}