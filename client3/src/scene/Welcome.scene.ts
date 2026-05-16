
import { useEffect } from "../index.ts";
import { useKeyboard } from "../Keyboard.ts";
import { socketSend } from "../middleware/socketMiddleware.ts";
import type { Screen } from "../Screen.ts";
import { changeScene } from "../store/appSlice.ts";
import { store } from "../store/index.ts";
import { writeCenter } from "../util.ts";


export const Welcome = (screen: Screen) => {
    
    // Подключаем горячие клавиши
    useEffect(() => {
        useKeyboard({
            Escape: () => {
                // Шаг 0. Дисгностическое сообщение
                console.log(`Нажат Escape на экране welcome. Запускаем обработчик.`);
                // Шаг 1. Запрашиваем список эхоконференций
                store.dispatch(socketSend({
                    msg: {
                        type: 'ECHO_INDEX',
                    },
                }));
                // Шаг 2. Переключаемся на страницу списка эхоконференций
                store.dispatch(changeScene('echo/index'));
            },
        });
    }, []);
    
    // Шаг 1. Рисуем рамку
    screen.setForegroudColor('#C0C0C0');
    screen.drawRect(0,0, 80 - 1, 25 - 1);

    // Шаг 2. Рисуем логотип
    screen.setForegroudColor('#d69a19');
    screen.writeText(4, 1, '   █████████           ████      █████                         ');
    screen.writeText(4, 2, '  ███▒▒▒▒▒███         ▒▒███     ▒▒███                          ');
    screen.writeText(4, 3, ' ███     ▒▒▒   ██████  ▒███   ███████   ██████  ████████       ');
    screen.writeText(4, 4, '▒███          ███▒▒███ ▒███  ███▒▒███  ███▒▒███▒▒███▒▒███      ');
    screen.writeText(4, 5, '▒███    █████▒███ ▒███ ▒███ ▒███ ▒███ ▒███████  ▒███ ▒███      ');
    screen.writeText(4, 6, '▒▒███  ▒▒███ ▒███ ▒███ ▒███ ▒███ ▒███ ▒███▒▒▒   ▒███ ▒███      ');
    screen.writeText(4, 7, ' ▒▒█████████ ▒▒██████  █████▒▒████████▒▒██████  ████ █████     ');
    screen.writeText(4, 8, '  ▒▒▒▒▒▒▒▒▒   ▒▒▒▒▒▒  ▒▒▒▒▒  ▒▒▒▒▒▒▒▒  ▒▒▒▒▒▒  ▒▒▒▒ ▒▒▒▒▒      ');
    
    screen.setForegroudColor('#f27d10');
    screen.writeText(14, 10, '    ███████████            ███              █████     ');
    screen.writeText(14, 11, '   ▒▒███▒▒▒▒▒███          ▒▒▒              ▒▒███      ');
    screen.writeText(14, 12, '    ▒███    ▒███  ██████  ████  ████████   ███████    ');
    screen.writeText(14, 13, '    ▒██████████  ███▒▒███▒▒███ ▒▒███▒▒███ ▒▒▒███▒     ');
    screen.writeText(14, 14, '    ▒███▒▒▒▒▒▒  ▒███ ▒███ ▒███  ▒███ ▒███   ▒███      ');
    screen.writeText(14, 15, '    ▒███        ▒███ ▒███ ▒███  ▒███ ▒███   ▒███ ███  ');
    screen.writeText(14, 16, '    █████       ▒▒██████  █████ ████ █████  ▒▒█████   ');
    screen.writeText(14, 17, '   ▒▒▒▒▒         ▒▒▒▒▒▒  ▒▒▒▒▒ ▒▒▒▒ ▒▒▒▒▒    ▒▒▒▒▒    ');
                                                                                                         
    // Шаг 2. Пишем название программы и версию
    
    screen.setForegroudColor('#c98f25');
    writeCenter(screen, 19, `v1.2.19`);

    // Шаг 3. Нажать Escape для продолжения
    screen.setForegroudColor('#3ebd0c');
    writeCenter(screen, 20, 'Нажмите Escape для продолжения...');

    // Шаг 4. Имена разработчиков
    screen.setForegroudColor('#ebb208');
    const contributors: string = 'Sergey Anohin, Andrey Mundirov, Jaroslav Bespalov';
    writeCenter(screen, 22, contributors);
    const contributors2: string = 'Richard Menedetter, Tommi Koivula, Rudi Timmermans';
    writeCenter(screen, 23, contributors2);

};
