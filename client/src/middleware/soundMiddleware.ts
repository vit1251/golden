
import { createAction, type Middleware } from '@reduxjs/toolkit';

export type Note = [number, number];

export type Music = 'SND_SAYBIBI' | 'SND_THEEND' | 'SND_GOTIT' | 'SND_TOOBAD' | 'SND_TOYOU';

const music: Record<Music, Array<Note>> = {
    SND_SAYBIBI: [
        [440, 111],
    ],
    SND_TOYOU: [
        [100, 18],
        [500, 18],
        [100, 18],
    ],
    SND_THEEND: [
        [220, 111],
        [110, 167],
    ],
    SND_GOTIT: [
        [110, 56],
        [220, 56],
        [110, 56],
        [220, 111],
    ],
    SND_TOOBAD: [
        [440, 111],
        [220, 111],
        [110, 167],
    ],
};

const audioCtx: AudioContext = new window.AudioContext();

/**
 * Проигрывам ноты
 * 
 */
async function playSequence(notes: Array<Note>) {
    
    if (audioCtx.state === 'suspended') {
        await audioCtx.resume();
    }

    let startTime = audioCtx.currentTime;

    notes.forEach(([freq, duration]) => {
        const oscillator = audioCtx.createOscillator();
        const gainNode = audioCtx.createGain();

        oscillator.type = 'sine';
        oscillator.frequency.setValueAtTime(freq, startTime);

        // Добавляем микро-затухание (0.01с), чтобы ноты не слипались и не щелкали
        gainNode.gain.setValueAtTime(1, startTime);
        gainNode.gain.exponentialRampToValueAtTime(0.001, startTime + (duration / 1000) - 0.01);

        oscillator.connect(gainNode);
        gainNode.connect(audioCtx.destination);

        oscillator.start(startTime);
        oscillator.stop(startTime + (duration / 1000));

        // Сдвигаем время старта для следующей ноты
        startTime += (duration / 1000);
    });
};

export const soundEvent = createAction<Music>('SOUND_EVENT');

export const soundMiddleware: Middleware = (store) => (next) => (action) => {

    // 1. Сначала пропускаем экшен дальше, чтобы состояние обновилось
    const result = next(action);

    // 1. Если это ваше конкретное событие
    if (soundEvent.match(action)) {
        const soundName: Music = action.payload; 
        const { [soundName]: notes } = music;
        if (notes) playSequence(notes);
    }
    
    return result;
};
