
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
import { TableComponent } from "./component/TableComponent.ts";



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

    // Шаг 2. Рендерим данные
    TableComponent(screen, {
        key: 'area_index',
        recordIndex: areaIndex,
        columns: [
            { name: `№п/п`, key: 'index', size: 5, adjust: 'right', render: (value: number) => `${value + 1} ` },
            { name: `Описание`, key: 'summary', size: 30, adjust: 'left', render: (value, record) => value ? value : record.name },
            { name: `Сообщений`, key: 'message_count', size: 10, adjust: 'right', render: (value) => `${value} ` },
            { name: `Новых`, key: 'new_message_count', size: 10, adjust: 'right', render: (value) => `${value} ` },
            { name: `Название эхи`, key: 'name', size: 20, adjust: 'left'},
        ],
        records: areas,
        sep: ' ',
        postRender: (screen: Screen, area: Area, index: number) => {
            if (area.new_message_count > 0) {
                screen.setForegroudColor(Color.White);
                screen.writeText(5, index + 1, '>');
            }
        }
    })


};
