
import { useEffect } from "react";

type KeyHandlers = {
    [key: string]: () => void;
};

export const useKeyboard = (handlers: KeyHandlers) => {
    useEffect(() => {
        const handleKeyDown = (event: KeyboardEvent) => {
            // Игнорируем нажатия, если фокус в поле ввода
            if (event.target instanceof HTMLInputElement ||
                event.target instanceof HTMLTextAreaElement ||
                (event.target as HTMLElement).isContentEditable) {
                return;
            }

            // Если для нажатой клавиши есть обработчик — вызываем его
            const { [event.key]: handler } = handlers;
            if (handler) {
                event.preventDefault();
                handler();
            }
        };
        console.log(`Подключили клавиши`);
        window.addEventListener('keydown', handleKeyDown);
        // Чистим за собой
        return () => {
            console.log(`Отключили клавиши`);
            window.removeEventListener('keydown', handleKeyDown);
        };
    }, [handlers]); // Хук обновится, если изменятся обработчики
};
