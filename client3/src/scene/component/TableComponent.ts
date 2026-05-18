
import { Color } from "../../color.ts";
import type { Screen } from "../../Screen.ts";
import { renderRow } from "../../util.ts";

export interface Column<T> {
    name: string,                                               // Название столбца
    key: 'index' | keyof T,                                     // Ключ для доступа к данным в элементе массива
    size: number,                                               // Ширина столбца
    adjust: 'left' | 'right',                                   // Выравнивание текста внутри ячейки
    render?: (value: any, record: T) => string,                 // Функция для рендера значения ячейки
}

export interface TableOptions<T> {
    columns: Array<Column<T>>,
    records: Array<T>,
    key: 'index' | keyof T,
    recordIndex: any,
    preRender?: (screen: Screen, record: T, index: number) => void,
    postRender?: (screen: Screen, record: T, index: number) => void
    sep?: string,
}

export const TableComponent = <T>(screen: Screen, { sep = '', key, columns, records, recordIndex, preRender, postRender }: TableOptions<T>) => {
    const componentWidth: number = 80 - 2;  // слева и справа у нас рамка
    const componentHeight: number = 25 - 2; // сверху у нас заголовок, а снизу рамка

    // Шаг 1. Рисуем рамку
    screen.setForegroudColor(Color.LightBlue);
    screen.drawRect(0, 0, 80 - 1, 25 - 1);

    // Шаг 2. Рисуем заголовок
    screen.setForegroudColor(Color.Yellow);
    let posX: number = 1;
    for (const { name, size } of columns) {
        screen.writeText(posX, 0, name);
        posX = posX + size; // позиция X начала столбца
        if (sep) posX = posX + 1; // если есть разделитель
    }

    // Шаг 3. Рисуем данные таблицы
    const startIndex: number = 0;
    const endIndex: number = startIndex + 25 - 2;
    for (let index = startIndex; index < endIndex; index++) {
        const record: T | undefined = records.at(index);
        if (!record) continue;
        const data = {
            ...record,
            index: index,
        };
        const { [key]: recordKey } = data;

        if (recordKey === recordIndex) {
            screen.setForegroudColor(Color.White);
            screen.setBackgroundColor(Color.Blue);
        } else {
            screen.setForegroudColor(Color.Gray);
            screen.setBackgroundColor(Color.Black);
        }

        const rowData: Array<{ value: string, size: number, adjust: 'left' | 'right' }> = [];
        for (const { key, size, adjust, render } of columns) {
            const { [key]: value } = data;
            rowData.push({
                value: render ? render(value, data) : `${value}`,
                size: size,
                adjust: adjust,
            });
        }
        if (preRender) preRender(screen, record, index);
        const row: string = renderRow(rowData, sep);
        screen.writeText(1, index + 1, row);
        if (postRender) postRender(screen, record, index);
    }


};
