
import { Color } from "../color.ts";
import { useEffect } from "../index.ts";
import { useKeyboard } from "../Keyboard.ts";
import { socketSend } from "../middleware/socketMiddleware.ts";
import { soundEvent } from "../middleware/soundMiddleware.ts";
import { Oops } from "../Oops.ts";
import type { Screen } from "../Screen.ts";
import { changeScene, messageNext, messagePrev, messageScrollDown, messageScrollUp, type Area, type Message } from "../store/appSlice.ts";
import { store } from "../store/index.ts";
import { fillEnd, searchArea, searchMessage, searchMessagePosition, useScope } from "../util.ts";

class Paragraph {
    protected constructor(readonly line: string, readonly prefix: string, readonly level: number) {
    }
    static parse(line: string) {
        const match = line.match(/^(\s*[A-Za-zА-Яа-я0-9]{1,3}>{1,3}|\s+[>]{1,3}>)/);
        if (match) {
            const prefix = match[0];
            // Считаем количество знаков ">", чтобы понять уровень вложенности
            const level: number = (prefix.match(/>/g) ?? []).length;
            return new Paragraph(line, prefix, level);
        }
        return new Paragraph(line, '', 0);
    }
}

/**
 * Цвета для выделения цитирования
 *
 */
function matchColor(level: number): string {
    if (level === 1) return Color.Cyan;
    if (level === 2) return Color.Magenta;
    if (level === 3) return Color.Green;
    return Color.Gray;
}

/**
 * Разбиваем на строки
 * 
 */
function parseLines(content: string): string[] {
    const text: string = content.replace(/\r\n|\r|\n/g, '\n');
    return text.split('\n');
}

/**
 * Проверяем строка относиться к специальным строкам
 * 
 */
function isOriginLine(str: string): boolean {
    return str.startsWith('---') || str.startsWith(' * Origin:');
}

/**
 * Переводим в число символов 
 * 
 */
function makeHuman(size: number): string {
    // Выводим в мегабайтах
    if (size > 1048576) {
        const megaByte: number = Math.round(size / 1048576);
        return `${megaByte}M`
    }
    // Выводим в килобайтах
    if (size > 1024) {
        const kiloByte: number = Math.round(size / 1024);
        return `${kiloByte}k`
    }
    // Выводим в байтах
    const byte: number = size;
    return `${byte}`
}

function makeFlags(msg: Message): string[] {
    const msgFlags: string[] = [];
    //msgFlags.push('Rcv');
    if (msg.view_count === 0) {
        msgFlags.push('New');
    }
    return msgFlags;
}

export const EchoMessageView = (screen: Screen) => {

    const { areas, areaIndex } = store.getState().app;
    const { messages, messageIndex } = store.getState().app;
    const { content, contentIndex } = store.getState().app;

    const area: Area | null = searchArea(areas, areaIndex);
    const message: Message | null = searchMessage(messages, messageIndex);
    if ((area === null) || (message === null)) {
        return Oops(screen);
    }

    // Шаг 1. Регистрируем обработчики клавиатуры
    useEffect(() => {
        useKeyboard({
            Escape: () => {
                // Шаг 1. Обновляем список сообещний
                store.dispatch(socketSend({
                    msg: {
                        type: 'ECHO_MSG_INDEX',
                        echoTag: areaIndex,
                    },
                }));
                // Шаг 2. Возвращаемся обратно
                store.dispatch(changeScene('echo/message/index'));
            },
            ArrowUp: () => store.dispatch(messageScrollUp()),
            ArrowDown: () => store.dispatch(messageScrollDown()),
            ArrowLeft: () => {
                // Шаг 1. Интерфейсные звуки
                useScope(() => {
                    const { areas, areaIndex } = store.getState().app;
                    const { messages, messageIndex } = store.getState().app;
                    const messagePostion: number = searchMessagePosition(messages, messageIndex) ?? Number.NaN;
                    if (messagePostion === 0) store.dispatch(soundEvent('SND_THEEND'));
                });
                // Шаг 2. Переходим на предыдущее сообщение
                store.dispatch(messagePrev());
                // Шаг 3. Загружаем сообщение
                useScope(() => {
                    const { areas, areaIndex } = store.getState().app;
                    const { messages, messageIndex } = store.getState().app;
                    store.dispatch(socketSend({
                        msg: {
                            type: 'ECHO_MSG_VIEW',
                            echoTag: areaIndex,
                            msgId: messageIndex,
                        },
                    }));
                });
            },
            ArrowRight: () => {
                // Шаг 1. Интерфейсные звуки
                useScope(() => {
                    const { areas, areaIndex } = store.getState().app;
                    const { messages, messageIndex } = store.getState().app;
                    const messagePostion: number = searchMessagePosition(messages, messageIndex) ?? Number.NaN;
                    if (messagePostion + 1 === messages.length) store.dispatch(soundEvent('SND_THEEND'));
                });
                // Шаг 2. Переходим на следующее сообщение
                store.dispatch(messageNext());
                // Шаг 3. Загружаем сообщение
                useScope(() => {
                    const { areas, areaIndex } = store.getState().app;
                    const { messages, messageIndex } = store.getState().app;
                    store.dispatch(socketSend({
                        msg: {
                            type: 'ECHO_MSG_VIEW',
                            echoTag: areaIndex,
                            msgId: messageIndex,
                        },
                    }));
                });
            }
        });
    }, []);

    // шаг 2. Ренлдерим шапку
    const msgIndex: number = searchMessagePosition(messages, messageIndex) ?? Number.NaN;
    const msgCount: number = messages.length;
    screen.setForegroudColor(Color.Blue); screen.drawLine(0);
    if (area.summary) {
        screen.setForegroudColor(Color.Yellow); screen.writeText(1, 0, ` ${area.summary} `);
    }
    const areaName: string = ` ${area.name} `;
    const posX: number = 80 - areaName.length - 1;
    screen.setForegroudColor(Color.Yellow); screen.writeText(posX, 0, areaName);

    // Первый столбец
    screen.setForegroudColor('#C0C0C0');
    screen.writeText(0, 1, ` №    : ${msgIndex + 1} of ${msgCount}`); // Номер сообзения в базе
    screen.writeText(0, 2, ` От   : ${fillEnd(message.from, 20)}`); // Отправитьель
    screen.writeText(0, 3, ` К    : ${fillEnd(message.to, 20)}`); // Получатель
    screen.writeText(0, 4, ` Тема : ${fillEnd(message.subject, 40)}`); // Тема сообщения
    // Второй столбец
    screen.setForegroudColor('#C0C0C0');
    const msgFlags: string[] = makeFlags(message);
    screen.writeText(40, 1, msgFlags.join(' ')); // Флаги
    screen.writeText(40, 2, ``); // Адрес отправителя
    screen.writeText(40, 3, ``); // Адрес получателя
    // Третий столбец
    screen.setForegroudColor('#C0C0C0');
    screen.writeText(58, 2, `${message.date}`); // Дата отправки
    screen.writeText(58, 3, `${message.date}`); // Дата получения
    //
    screen.setForegroudColor(Color.Blue); screen.drawLine(5);
    const msgSize: string = makeHuman(content.length);
    screen.setForegroudColor(Color.Yellow); screen.writeText(1, 5, msgSize);

    // Шаг 3. Рендерим тело сообщения
    const paragraphLines: string[] = parseLines(content);
    for (const [index, paragraphLine] of paragraphLines.entries()) {
        if (index >= contentIndex) {

            // Шаг 1. Проверка цитирования
            const p: Paragraph = Paragraph.parse(paragraphLine);
            const color: string = matchColor(p.level);
            screen.setForegroudColor(color);

            // Шаг 2. Проверка оригинов и тирлайнов
            if (isOriginLine(paragraphLine)) {
                screen.setForegroudColor(Color.White);
            }

            // Шаг 3. Собственно рендеринг текста
            screen.writeText(0, 7 + index - contentIndex, paragraphLine);
        }
    }

};
