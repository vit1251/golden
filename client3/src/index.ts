
import { Oops } from './Oops.ts';
import { EchoIndex } from './scene/EchoIndex.scene.ts';
import { EchoMessageIndex } from './scene/EchoMessageIndex.scene.ts';
import { EchoMessageView } from './scene/EchoMessageView.scene.ts';
import { Welcome } from './scene/Welcome.scene.ts';
import { Screen } from './Screen.ts';
import { store } from './store/index.ts';

import './styles.css';

import { socketConnect } from './middleware/socketMiddleware.ts';

const CHAR_WIDTH = 10;
const CHAR_HEIGHT = 18;
const FONT_SIZE = 16;


function render(canvas: HTMLCanvasElement, screen: Screen) {
    const ctx = canvas.getContext('2d');
    if (!ctx) return;

    // Шаг 1. Диагностическое сообщение
    console.log(`Перерисовка жэкрана`);

    // 2. Рассчитываем размеры холста на основе матрицы
    const rowsCount = screen.height;
    const colsCount = screen.width;

    canvas.width = colsCount * CHAR_WIDTH;
    canvas.height = rowsCount * CHAR_HEIGHT;

    // 3. Очищаем холст и задаем базовые стили
    ctx.fillStyle = '#000000'; // Цвет фона
    ctx.fillRect(0, 0, canvas.width, canvas.height);

    ctx.fillStyle = '#ffffff'; // Цвет текста псевдографики
    ctx.font = `${FONT_SIZE}px Courier`;

    // Математическое выравнивание текста внутри ячейки
    ctx.textAlign = 'center';
    ctx.textBaseline = 'middle';

    // 4. Отрисовка символов в один проход
    for (let row = 0; row < rowsCount; row++) {
        for (let col = 0; col < colsCount; col++) {
            const char: string = screen.matrix[row]![col] ?? ' ';
          
            // Вычисляем точный центр ячейки
            const x: number = col * CHAR_WIDTH;
            const y: number = row * CHAR_HEIGHT;

            // Рендерим фон
            ctx.fillStyle = screen.back[row]![col] ?? '#000000'; // Цвет фона
            ctx.fillRect(x, y, CHAR_WIDTH, CHAR_HEIGHT);
            // Рендерим символ
            if (char !== ' ') {
                ctx.fillStyle = screen.front[row]![col] ?? '#C0C0C0'; // цвет символа
                const chX: number = x + CHAR_WIDTH / 2;
                const chY: number = y + CHAR_HEIGHT / 2;
                ctx.fillText(char, chX, chY);
            }
        }
    }
}

function handlePaint(canvas: HTMLCanvasElement, screen: Screen) {
    const { scene } = store.getState().app;
    // Шаг 0. Диагностическая информация
    console.log(`Выполенние сцены: scene = `, scene);
    // Шаг 1. Очистка экрана
    screen.reset();
    // Шаг 2. Выполнение сцены
    if (scene === 'welcome') Welcome(screen);
    else if (scene === 'echo/index') EchoIndex(screen);
    else if (scene === 'echo/message/index') EchoMessageIndex(screen);
    else if (scene === 'echo/message/view') EchoMessageView(screen);
    else Oops(screen);
    // Шаг 3. Отрисовка полученного экрана
    render(canvas, screen);
}

const app = document.getElementById('app');
if (app) {

    // Шаг 1. Создаем Canvas для рисования
    const canvas: HTMLCanvasElement = document.createElement('canvas');
    app.appendChild(canvas);

    // Шаг 2. Подключаем соедниение с сервером
    store.dispatch(socketConnect());

    // Шаг 3. Дисплей
    const screen: Screen = new Screen(80, 25);

    // Шаг 4. Подписываемся на обновления
    store.subscribe(() => handlePaint(canvas, screen));

    // Шаг 5. Начальная инициализация
    handlePaint(canvas, screen);

}

export function useEffect(callback: () => void, depends: Array<any>) {
    console.log(`Работает обработчик эффекта`);
    callback();
}

