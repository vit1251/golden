
import { Color } from "./color.ts";
import { useEffect } from "./index.ts";
import { useKeyboard } from "./Keyboard.ts";
import type { Screen } from "./Screen.ts";
import { store } from "./store/index.ts";


export const Oops = (screen: Screen) => {

    // Шаг 1. Занулить работу с горячими клавишами
    useEffect(() => {
        useKeyboard({});
    }, []);

    // Шаг 2. Заголовок
    screen.setForegroudColor(Color.Yellow);
    screen.writeText(4, 2, 'Ошибка приложeния');
    screen.writeText(4, 3, '=================');

    // Шаг 3. Текст сообщения
    screen.setForegroudColor(Color.Gray);
    screen.writeText(4, 5, `Произошла внутренняя неустранимая ошибка.`);
    screen.writeText(4, 7, `Перезагрузите приложение.`);
    
};