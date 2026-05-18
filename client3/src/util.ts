
import type { Screen } from "./Screen.ts";
import type { Area, Message } from "./store/appSlice.ts";

/* Утилиты работы с текстом */

/**
 * Пишем текст посредине
 *
 */
export function writeCenter(screen: Screen, y: number, msg: string) {
    const centerX: number = Math.round((80 - msg.length) / 2);
    screen.writeText(centerX, y, msg);
}

/* Работа со строками */

export function fillStart(str: string, size: number): string {
    if (str.length < size) {
        return str.padStart(size, ' ');
    }
    const suffix: string = str.slice(str.length - size + 1);
    return `…${suffix}`;
}

export function fillEnd(str: string, size: number): string {
    if (str.length < size) {
        return str.padEnd(size, ' ');
    }
    const prefix: string = str.slice(0, size - 1);
    return `${prefix}…`;
}

/* Работа с данными */

export function searchArea(areas: Array<Area>, areaIndex: string): Area | null {
    for (const area of areas) {
        if (area.area_index === areaIndex) {
            return area;
        }
    }
    return null;
}

export function searchMessage(messages: Array<Message>, messageIndex: string): Message | null {
    for (const message of messages) {
        if (message.hash === messageIndex) {
            return message;
        }
    }
    return null;
}

/**
 * Возвращает порфядковый номер элемента в списке
 * 
 */
export function searchAreaPosition(areas: Array<Area>, areaIndex: string): number | null {
    for (const [index, area] of areas.entries()) {
        if (area.area_index === areaIndex) {
            return index;
        }
    }
    return null;
}

export function searchMessagePosition(messages: Array<Message>, messageIndex: string): number | null {
    for (const [index, message] of messages.entries()) {
        if (message.hash === messageIndex) {
            return index;
        }
    }
    return null;
}

/* Формирование табличек */

export function stringAdjust(str: string, size: number, adjust: 'left' | 'right'): string {
    if (adjust === 'left') return fillEnd(str, size);
    if (adjust === 'right') return fillStart(str, size);
    throw new Error(`Непонятное выравнивание "${adjust}"`);
}

/**
 * Генерация строки таблицы
 * 
 */
export function renderRow(columns: Array<{ value: string, size: number, adjust: 'left' | 'right' }>, sep: string = ''): string {
    const parts: string[] = [];
    for (const { value, size, adjust } of columns) {
        const str: string = stringAdjust(value, size, adjust);
        parts.push(str);
    }
    return parts.join(sep);
}

/**
 * Создаем область видимости для переменных
 * 
 * 
 */
export const useScope = (callback: () => void) => {
    callback();
}