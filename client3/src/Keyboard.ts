
class KeyboardHandlers {
    protected handlers: Partial<Record<Key, () => void>>;
    constructor() {
        this.handlers = {};
    }
    reset() {
        this.handlers = {};
    }
    register(key: Key, callback: () => void) {
        this.handlers[key] = callback;
    }
    setup(handlers: Partial<Record<Key, () => void>>) {
        const keyNames: string[] = Object.keys(handlers);
        console.log(`Регистриуем новые обработчики клавиатуры на клавиши: keys = `, keyNames);
        this.handlers = handlers;
    }
    fireEvent(eventName: 'keyup' | 'keydown', key: Key) {
        // Шаг 1. Обработка отпускания кнопки
        if (eventName === 'keydown') {
            console.log(`Событие обработчика`);
            console.log(`Обработчики: `, this.handlers);
            const { [key]: keyHandlers } = this.handlers;
            if (keyHandlers) {
                const newKeyHandlers = Array.isArray( keyHandlers ) ? keyHandlers : [ keyHandlers ];
                for (const keyHandler of newKeyHandlers) {
                    keyHandler();
                }
            }
        }
        // Шаг 2. Обработка отпускания кнопки
        if (eventName === 'keyup') {
            return;
        }
    }
}

const keyboardHandlers: KeyboardHandlers = new KeyboardHandlers();

function handleKeyboard(eventName: 'keyup' | 'keydown', key: Key) {
    
    // Шаг 1. Диагностическая информация
    console.log(`Событие "${eventName}" с клавишей `, key);

    // Шаг 2. Отправка сообытия
    keyboardHandlers.fireEvent(eventName, key);

}

window.addEventListener('keydown', (event: KeyboardEvent) => handleKeyboard('keydown', event.key as Key));
window.addEventListener('keyup', (event: KeyboardEvent) => handleKeyboard('keyup', event.key as Key));

export type Key = 'Enter' | 'Escape' | 'ArrowUp' | 'ArrowDown' | 'ArrowLeft'| 'ArrowRight';

export function useKeyboard(curKeyboardHanderls: Partial<Record<Key, () => void>>) {
    keyboardHandlers.setup(curKeyboardHanderls);
}
