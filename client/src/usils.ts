
/**
 * Возвращает строку в верхнем регистре
 * 
 * @param str 
 * @returns 
 */
export function stringUpper(str: string) {
    return str.toUpperCase();
}

/**
 * Собирает инициалы имени
 * 
 * Пример:
 *   Vitold Sedyshev => VS
 *   Alexey Pavlov => AP
 * 
 */
export function makeShort(str: string) {
    const parts: string[] = str.split(' ');
    const chars: string[] = [];
    for (const part of parts) {
        if (part.length > 0) {
            chars.push(stringUpper(part[0]));
        }
    }
    return chars.join('');
}

export function stringToHexColor(input: string): string {
    let hash = 0;
    for (let i = 0; i < input.length; i++) {
        hash = input.charCodeAt(i) + ((hash << 5) - hash);
    }

    // Преобразуем хэш в шестнадцатеричный цвет
    const r = (hash & 0xFF0000) >> 16; // Красный компонент
    const g = (hash & 0x00FF00) >> 8;  // Зеленый компонент
    const b = hash & 0x0000FF;         // Синий компонент

    // Преобразуем компоненты в двузначные шестнадцатеричные значения
    const toHex = (value: number): string => {
        const hex = value.toString(16);
        return hex.length === 1 ? '0' + hex : hex; // Добавляем ведущий ноль, если нужно
    };

    // Формируем строку в формате #rrggbb
    return `#${toHex(r)}${toHex(g)}${toHex(b)}`;
}