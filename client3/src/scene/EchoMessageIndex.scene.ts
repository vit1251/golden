
import { useEffect } from "../index.ts";
import { useKeyboard } from "../Keyboard.ts";
import { socketSend } from "../middleware/socketMiddleware.ts";
import type { Screen } from "../Screen.ts";
import { changeScene, messageNext, messagePrev } from "../store/appSlice.ts";
import { store } from "../store/index.ts";
import { fillEnd } from "../util.ts";


export const EchoMessageIndex = (screen: Screen) => {

    //
    const { areas, areaIndex } = store.getState().app;
    const { messages, messageIndex } = store.getState().app;

    // Шаг 1. Регистрируем обработчики клавиатуры
    useEffect(() => {

        useKeyboard({
            Escape: () => store.dispatch(changeScene('echo/index')),
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
            ArrowUp: () => store.dispatch(messagePrev()),
            ArrowDown: () => store.dispatch(messageNext()),
        });

    }, []);


    // Шаг 2. Отображаем список сообщений
    
    for (const [index, message] of messages.entries()) {
        screen.setBackgroundColor('#000000');
        if (message.hash === messageIndex) screen.setBackgroundColor('#C00000');
        const row: string = [
            fillEnd(`${message.subject}`, 40),
            fillEnd(`${message.from}`, 20),
            fillEnd(`${message.date}`, 12),
        ].join('');
        screen.setForegroudColor('#C0C0C0');
        if (message.view_count === 0) screen.setForegroudColor('#D0D000');
        screen.writeText(2, index + 4, row);

    }


};