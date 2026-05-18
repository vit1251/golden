
import { Color } from "../color.ts";
import { useEffect } from "../index.ts";
import { useKeyboard } from "../Keyboard.ts";
import { socketSend } from "../middleware/socketMiddleware.ts";
import { soundEvent } from "../middleware/soundMiddleware.ts";
import type { Screen } from "../Screen.ts";
import { changeScene, messageEnd, messageHome, messageNext, messagePrev, type Message } from "../store/appSlice.ts";
import { store } from "../store/index.ts";
import { renderRow, searchMessagePosition } from "../util.ts";
import { TableComponent } from "./component/TableComponent.ts";


export const EchoMessageIndex = (screen: Screen) => {

    // Шаг 0. Данные из хранилища
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


    // Шаг 2. Рисуем таблицу с данными
    TableComponent(screen, {
        columns: [
            { name: '№п/п', key: 'index', size: 6, adjust: 'right', render: (value: number) => `${value + 1} ` },
            { name: 'Тема', key: 'subject', size: 32, adjust: 'left' },
            { name: 'Отправитель', key: 'from', size: 20, adjust: 'left' },
            { name: 'Дата', key: 'date', size: 17, adjust: 'left' },
        ],
        key: 'hash',
        recordIndex: messageIndex,
        records: messages,
        preRender: (screen: Screen, message: Message, index: number) => {
            if (message.view_count === 0) {
                screen.setForegroudColor(Color.Yellow);
            }
        },
        sep: ' ', // разделитель пробел
    });

};