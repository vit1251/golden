
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
    return '…'.concat(str.slice(str.length - size + 1, size));
}

export function fillEnd(str: string, size: number): string {
    if (str.length < size) {
        return str.padEnd(size, ' ');
    }
    return str.slice(0, size - 1).concat('…');
}

/* Работа с данными */

export function searchArea(areas: Array<Area>, areaIndex: string): Area {
    for (const area of areas) {
        if (area.area_index === areaIndex) {
            return area;
        }
    }
    throw new Error(`Внутренняя ошибка`);
}

export function searchMessage(messages: Array<Message>, messageIndex: string): Message {
    for (const message of messages) {
        if (message.hash === messageIndex) {
            return message;
        }
    }
    throw new Error(`Внутренняя ошибка`);
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