
import { useEffect } from "../index.ts";
import { useKeyboard } from "../Keyboard.ts";
import { socketSend } from "../middleware/socketMiddleware.ts";
import { soundEvent } from "../middleware/soundMiddleware.ts";
import type { Screen } from "../Screen.ts";
import { changeScene, nextArea, prevArea, type Area } from "../store/appSlice.ts";
import { store } from "../store/index.ts";
import { fillEnd, fillStart } from "../util.ts";


export const EchoIndex = (screen: Screen) => {

    const { areas, areaIndex } = store.getState().app;

    useEffect(() => {

        useKeyboard({
            Enter: () => {
                // Шаг 0. Находим конференцию
                const area: Area | null = areas.find(a => a.area_index === areaIndex) ?? null;
                if (area === null) {
                    console.error(`Нет эхоконференции с идентификатором "${areaIndex}"`);
                    return;
                }

                // Шаг 1. Запросим обновление списка сообщений
                store.dispatch(socketSend({
                    msg: {
                        type: 'ECHO_MSG_INDEX',
                        echoTag: area.area_index,
                    },
                }))

                // Шаг 2. Переходим на сцену сообещний
                store.dispatch(changeScene('echo/message/index'));
            },
            ArrowUp: () => {
                store.dispatch(prevArea());
                store.dispatch(soundEvent('SND_THEEND'));
            },
            ArrowDown: () => {
                store.dispatch(soundEvent('SND_THEEND'));
                store.dispatch(nextArea());
            },
        });

    }, []);

    // Шаг 1. Отрисуем список эхоконференций
    screen.setForegroudColor('#C0C0C0');
    for (const [index, area] of areas.entries()) {
        screen.setBackgroundColor('#000000');
        if (area.area_index === areaIndex) screen.setBackgroundColor('#C00000');
        const row: string = [
            fillEnd(`${area.name}`, 20),
            fillEnd(`${area.summary}`, 30),
            fillStart(`${area.message_count}`, 6),
            fillStart(`${area.new_message_count}`, 6),
        ].join('');
        screen.writeText(2, index + 4, row);

    }

};
