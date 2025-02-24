
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
    /* Больше двух инициалов берем первый и последний */
    const charLength: number = chars.length;
    if (charLength > 2) {
        return `${chars[0]}${chars[charLength - 1]}`;
    }
    /* В остальных случаях: 0, 1, 2 склеиваем как есть */
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

export function adjustBrightness(hexColor: string, factor: number): string {
    // Проверяем, что hexColor начинается с # и имеет правильную длину
    if (!/^#[0-9A-Fa-f]{6}$/.test(hexColor)) {
        throw new Error("Некорректный формат цвета. Ожидается #rrggbb.");
    }

    // Извлекаем компоненты цвета
    const r = parseInt(hexColor.slice(1, 3), 16);
    const g = parseInt(hexColor.slice(3, 5), 16);
    const b = parseInt(hexColor.slice(5, 7), 16);

    // Функция для изменения яркости компонента
    const adjustComponent = (value: number): number => {
        const newValue = value * factor;
        return Math.min(255, Math.max(0, Math.round(newValue))); // Ограничиваем значение от 0 до 255
    };

    // Применяем изменение яркости к каждому компоненту
    const newR = adjustComponent(r);
    const newG = adjustComponent(g);
    const newB = adjustComponent(b);

    // Преобразуем компоненты обратно в шестнадцатеричный формат
    const toHex = (value: number): string => {
        const hex = value.toString(16);
        return hex.length === 1 ? '0' + hex : hex; // Добавляем ведущий ноль, если нужно
    };

    // Формируем новый цвет в формате #rrggbb
    return `#${toHex(newR)}${toHex(newG)}${toHex(newB)}`;
}
