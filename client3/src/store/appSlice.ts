
import { createAction, createSlice, type PayloadAction } from '@reduxjs/toolkit';
import { socketConnected, socketDisconnect } from '../middleware/socketMiddleware.ts';

export type Scene = 'welcome' | 'echo/index' | 'echo/message/index' | 'echo/message/view';

export type Area = {
    name: string,              // "1081.LOCAL"
    summary: string,           // ""
    message_count: number,     // 22
    new_message_count: number, // 0
    order: number,             // 1773873848
    area_index: string,        // "fce8f8e7-5ce9-490b-bae6-bad560faa3b0"
};

export type Message = {
    date: string,                   // "2026-04-24 11:14"
    from: string,                   // "Ivan Zelenyi"
    hash: string,                   // "69eb26c7"
    subject: string,                // "телеграм на узле ВСЁ"
    to: string,                     // ""
    view_count: number,             // 5
};

// 1. Описываем тип для структуры нашего состояния
export interface AppState {
    ready: boolean,                // Установлено соединение с приложением
    scene: Scene,                  // Текущая сцена приложения
    areas: Array<Area>,            // Список эхоконференций
    areaIndex: string,             // Выбранная эхоконференция
    messages: Array<Message>,      // Список сообщений
    messageIndex: string,          // Выбранное (текущее) сообещние
    content: string,               // Текст сообщения
    contentIndex: number,          // Первая отображаемая строка
}

// 2. Задаем начальное значение
const initialState: AppState = {
    ready: false,
    scene: 'welcome',
    areas: [],
    areaIndex: '',
    messages: [],
    messageIndex: '',
    content: '',
    contentIndex: 0,
};

const EchoIndexAction = createAction<{ areas: Array<Area> }>('ECHO_INDEX');
const EchoMessageIndexAction = createAction<{ headers: Array<Message> }>('ECHO_MSG_INDEX');
const EchoMessageViewAction = createAction<{ area: { name: string }, body: string, echo: { from: string, to: string, subject: string, date: string } }>('ECHO_MSG_VIEW');

// 3. Создаем слайс (редьюсер + экшены в одном флаконе)
const appSlice = createSlice({
    name: 'app',
    initialState,
    reducers: {
        changeScene: (state, action: PayloadAction<Scene>) => {
            state.scene = action.payload;
        },
        areaHome: (state) => {
            const candidate: Area | null = state.areas.at(0) ?? null;
            if (candidate) {
                state.areaIndex = candidate.area_index;
            }
        },
        areaEnd: (state) => {
            const candidate: Area | null = state.areas.at(-1) ?? null;
            if (candidate) {
                state.areaIndex = candidate.area_index;
            }
        },
        nextArea: (state) => {
            let candidat: Area | null = null;
            for (const [index, area] of state.areas.entries()) {
                if (area.area_index === state.areaIndex) {
                    if (index + 1 < state.areas.length) {
                        candidat = state.areas.at(index + 1) ?? null;
                    }
                }
            }
            if (candidat) {
                state.areaIndex = candidat.area_index;
            }
        },
        prevArea: (state) => {
            let candidat: Area | null = null;
            for (const [index, area] of state.areas.entries()) {
                if (area.area_index === state.areaIndex) {
                    if (index - 1 >= 0) {
                        candidat = state.areas.at(index - 1) ?? null;
                    }
                }
            }
            if (candidat) {
                state.areaIndex = candidat.area_index;
            }
        },
        messageHome: (state) => {
            const candidate: Message | null = state.messages.at(0) ?? null;
            if (candidate) {
                state.messageIndex = candidate.hash;
            }
        },
        messageEnd: (state) => {
            const candidate: Message | null = state.messages.at(-1) ?? null;
            if (candidate) {
                state.messageIndex = candidate.hash;
            }
        },
        messagePrev: (state) => {
            let candidat: Message | null = null;
            for (const [index, message] of state.messages.entries()) {
                if (message.hash === state.messageIndex) {
                    if (index - 1 >= 0) {
                        candidat = state.messages.at(index - 1) ?? null;
                    }
                }
            }
            if (candidat) {
                state.messageIndex = candidat.hash;
            }
        },
        messageNext: (state) => {
            let candidat: Message | null = null;
            for (const [index, message] of state.messages.entries()) {
                if (message.hash === state.messageIndex) {
                    if (index + 1 < state.messages.length) {
                        candidat = state.messages.at(index + 1) ?? null;
                    }
                }
            }
            if (candidat) {
                state.messageIndex = candidat.hash;
            }
        },
        /* Работа с сообщениями */
        messageScrollUp: (state) => {
            state.contentIndex = state.contentIndex > 0 ? state.contentIndex - 1 : state.contentIndex;
        },
        messageScrollDown: (state) => {
            state.contentIndex = state.contentIndex + 1;
        }
    },
    extraReducers: (builder) => {
        builder
            /* Упраление соединением */
            .addCase(socketConnected, (state) => {
                state.ready = true;
            })
            .addCase(socketDisconnect, (state) => {
                state.ready = false;
            })
            /* Сообщения от сервера с данными */
            .addCase(EchoIndexAction, (state, action) => {
                console.log(`Получен список эхоконференций`);
                state.areas = action.payload.areas;
                const areaIndexes: string[] = state.areas.map(a => a.area_index);
                if (!areaIndexes.includes(state.areaIndex) || (state.areaIndex === '')) {
                    const area1: Area | null = state.areas.at(0) ?? null;
                    state.areaIndex = area1 ? area1.area_index : '';
                }
            })
            .addCase(EchoMessageIndexAction, (state, action) => {
                console.log(`Получен список сообщений в эхоконференции`);
                state.messages = action.payload.headers;
                const messageIndexes: string[] = state.messages.map(m => m.hash);
                if (!messageIndexes.includes(state.messageIndex) || (state.messageIndex === '')) {
                    for (const message of state.messages) {
                        state.messageIndex = message.hash;
                        if (message.view_count === 0) {
                            break;
                        }
                    }
                }
            })
            .addCase(EchoMessageViewAction, (state, action) => {
                console.log(`Получено тело сообщения`);
                state.content = action.payload.body;
                state.contentIndex = 0;
            });
    }
});

// Экспортируем автоматически созданные экшены
export const { changeScene } = appSlice.actions;
export const { prevArea, nextArea } = appSlice.actions;
export const { areaHome, areaEnd } = appSlice.actions;
export const { messageNext, messagePrev } = appSlice.actions;
export const { messageHome, messageEnd } = appSlice.actions;
export const { messageScrollUp, messageScrollDown } = appSlice.actions;

// Экспортируем редьюсер для подключения в стор
export default appSlice.reducer;