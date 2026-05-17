
import { Color } from "../color.ts";
import { useEffect } from "../index.ts";
import { useKeyboard } from "../Keyboard.ts";
import { socketSend } from "../middleware/socketMiddleware.ts";
import { soundEvent } from "../middleware/soundMiddleware.ts";
import type { Screen } from "../Screen.ts";
import { changeScene, messageEnd, messageHome, messageNext, messagePrev } from "../store/appSlice.ts";
import { store } from "../store/index.ts";
import { fillEnd, fillStart, renderRow, searchMessagePosition } from "../util.ts";


export const EchoMessageIndex = (screen: Screen) => {

    //
    const { areas, areaIndex } = store.getState().app;
    const { messages, messageIndex } = store.getState().app;

    // Шаг 1. Регистрируем обработчики клавиатуры
    useEffect(() => {

        useKeyboard({
            Escape: () => {
                // Шаг 1. Обновляем список конференций
                store.dispatch(socketSend({
                    msg: {
                        type: 'ECHO_INDEX',
                    },
                }));
                // Шаг 2. Перехзодим в список конференций
                store.dispatch(changeScene('echo/index'));
            },
            Enter: () => {
                // Шаг 1. Подгружаем сообщение
                store.dispatch(socketSend({
                    msg: {
                        type: 'ECHO_MSG_VIEW',
	                    echoTag: areaIndex,
	                    msgId: messageIndex,
                    },
                }));
                // Шаг 2. Переходим к чтению сообщения
                store.dispatch(changeScene('echo/message/view'));
            },
            Home: () => {
                store.dispatch(messageHome());
            },
            End: () => {
                store.dispatch(messageEnd());
            },
            ArrowUp: () => {
                const messagePosition: number = searchMessagePosition(messages, messageIndex) ?? Number.NaN;
                if (messagePosition === 0) store.dispatch(soundEvent('SND_THEEND'));
                store.dispatch(messagePrev());
            },
            ArrowDown: () => {
                const messagePosition: number = searchMessagePosition(messages, messageIndex) ?? Number.NaN;
                if (messagePosition + 1 === messages.length) store.dispatch(soundEvent('SND_THEEND'));
                store.dispatch(messageNext());
            },
        });

    }, []);


    // Шаг 2. Оформление
    screen.setForegroudColor(Color.LightBlue);
    screen.drawRect(0, 0, 80 - 1, 25 - 1);
    screen.setForegroudColor(Color.Yellow);
    screen.writeText(2, 0, '№п/п');
    screen.writeText(8, 0, 'Тема');
    screen.writeText(42, 0, 'Отправитель');
    screen.writeText(62, 0, 'Дата');

    // Шаг 2. Отображаем список сообщений
    for (const [index, message] of messages.entries()) {
        
        if (message.hash === messageIndex) {
            screen.setForegroudColor(Color.White);
            screen.setBackgroundColor(Color.Blue);
        } else {
            screen.setForegroudColor(Color.Gray);
            screen.setBackgroundColor(Color.Black);
        }
        const row: string = renderRow([
            { value: `${index + 1}`, size: 5, adjust: 'right'},
            { value: `${message.subject}`, size: 33, adjust: 'left' },
            { value: `${message.from}`, size: 20, adjust: 'left' },
            { value: `${message.date}`, size: 17, adjust: 'left' },
        ]);
        if (message.view_count === 0) {
            screen.setForegroudColor(Color.Yellow);
        }
        screen.writeText(1, index + 1, row);

    }


};