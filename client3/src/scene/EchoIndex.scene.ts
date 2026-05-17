
import { Color } from "../color.ts";
import { useEffect } from "../index.ts";
import { useKeyboard } from "../Keyboard.ts";
import { socketSend } from "../middleware/socketMiddleware.ts";
import { soundEvent } from "../middleware/soundMiddleware.ts";
import { Oops } from "../Oops.ts";
import type { Screen } from "../Screen.ts";
import { areaEnd, areaHome, changeScene, nextArea, prevArea, type Area } from "../store/appSlice.ts";
import { store } from "../store/index.ts";
import { renderRow, searchArea, searchAreaPosition, useScope } from "../util.ts";



export const EchoIndex = (screen: Screen) => {

    const { areas, areaIndex } = store.getState().app;
    const area: Area | null = searchArea(areas, areaIndex);
    if (area === null) {
        return Oops(screen);
    }

    useEffect(() => {

        useKeyboard({
            Enter: () => {

                // Шаг 1. Запросим обновление списка сообщений
                store.dispatch(socketSend({
                    msg: {
                        type: 'ECHO_MSG_INDEX',
                        echoTag: area.area_index,
                    },
                }))

                // Шаг 3. Переходим на сцену сообещний
                store.dispatch(changeScene('echo/message/index'));
                
            },
            Home: () => {
                store.dispatch(areaHome());
            },
            End: () => {
                store.dispatch(areaEnd());
            },
            ArrowUp: () => {
                // Шаг 1. Проверка границ и пользовательское уведомление
                const areaPosition: number = searchAreaPosition(areas, areaIndex) ?? Number.NaN;
                if (areaPosition === 0) store.dispatch(soundEvent('SND_THEEND'));
                // Шаг 2. Операция по переходу
                store.dispatch(prevArea());
            },
            ArrowDown: () => {
                // Шаг 1. Проверка границ и пользовательское уведомление
                const areaPosition: number = searchAreaPosition(areas, areaIndex) ?? Number.NaN;
                if (areaPosition + 1 === areas.length) store.dispatch(soundEvent('SND_THEEND'));
                // Шаг 2. Операция по переходу
                store.dispatch(nextArea());
            },
        });

    }, []);

    // Шаг 2. Оформление
    screen.setForegroudColor(Color.LightBlue);
    screen.drawRect(0, 0, 80 - 1, 25 - 1);

    screen.setForegroudColor(Color.Yellow);
    screen.writeText(2, 0, `№п/п`);              //
    screen.writeText(8, 0, `Описание`);
    screen.writeText(40, 0, `Сообщений`);
    screen.writeText(50, 0, `Новых`);
    screen.writeText(60, 0, `Название эхи`);
            

    // Шаг 3. Отрисуем список эхоконференций
    for (const [index, area] of areas.entries()) {
        // Шаг 1. Выставить цвета на выделенном элементе
        if (area.area_index === areaIndex) {
            screen.setForegroudColor(Color.White);
            screen.setBackgroundColor(Color.Blue);
        } else {
            screen.setForegroudColor(Color.Gray);
            screen.setBackgroundColor(Color.Black);
        }
        const areaSummary: string = area.summary ? area.summary : area.name;
        const row: string = renderRow([
            { value: `${index + 1} `, size: 4, adjust: 'right' },
            { value: `${areaSummary}`, size: 30, adjust: 'left' },
            { value: `${area.message_count}`, size: 10, adjust: 'right' },
            { value: `${area.new_message_count}`, size: 10, adjust: 'right'},
            { value: `${area.name}`, size: 20, adjust: 'left'},
        ]);
        
        screen.writeText(1, index + 1, row);
        // Шаг 3. РОисуем маркер
        if (area.new_message_count > 0) {
            screen.setForegroudColor(Color.White);
            screen.writeText(4, index + 1, '>');
        }
    }

};
