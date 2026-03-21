
type Note = [number, number];

let audioCtx: AudioContext = new (window.AudioContext || window.webkitAudioContext)();

export const playSequence = async (notes: Array<Note>) => {
//    if (!audioCtx) {
//        audioCtx = new (window.AudioContext || window.webkitAudioContext)();
//    }
    
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

export function playError() {
    playSequence([
        [ 440, 50 ],
    ]);
}

const ctx = audioCtx;
export function playModemBeep() {
    const now = ctx.currentTime;
    const master = ctx.createGain();
    master.connect(ctx.destination);

    // 1. Dial Tone (425 Гц — стандарт в СССР/РФ)
    // Это тот самый непрерывный гудок "линия свободна"
    const dialTone = ctx.createOscillator();
    dialTone.frequency.value = 425;
    const dialGain = ctx.createGain();
    
    dialGain.gain.setValueAtTime(0.1, now);
    dialGain.gain.setValueAtTime(0, now + 1.0); // Гудит секунду и прерывается набором
    
    dialTone.connect(dialGain).connect(master);
    dialTone.start(now);
    dialTone.stop(now + 1.0);

    // 2. Pulse Dialing (Щелчки разрыва линии)
    // Цифра "3" = 3 щелчка. Частота пульсации — 10 Гц (стандарт)
    const pulseStart = now + 1.2;
    const numPulses = 7;
    const pulseWidth = 0.05; // 50мс — длительность разрыва
    const pulseInterval = 0.1; // 100мс — цикл (10 Гц)

    for (let i = 0; i < numPulses; i++) {
        const t = pulseStart + (i * pulseInterval);
        
        // Щелчок — это резкий низкочастотный удар + белый шум (треск контактов)
        const click = ctx.createOscillator();
        click.type = 'sawtooth';
        click.frequency.setValueAtTime(60, t); // Глухой удар реле
        
        const clickGain = ctx.createGain();
        clickGain.gain.setValueAtTime(0.3, t);
        clickGain.gain.setValueAtTime(0, t + pulseWidth);

        click.connect(clickGain).connect(master);
        click.start(t);
        click.stop(t + pulseWidth + 0.01);
    }

    // 3. Пилот-тон модема (2100 Гц) — ответ после набора
    const pilotStart = pulseStart + (numPulses * pulseInterval) + 0.5;
    const pilot = ctx.createOscillator();
    pilot.frequency.value = 2100;
    const pilotGain = ctx.createGain();
    
    pilotGain.gain.setValueAtTime(0.2, pilotStart);
    pilotGain.gain.setValueAtTime(0, pilotStart + 1.5); // Просто рубим в конце

    pilot.connect(pilotGain).connect(master);
    pilot.start(pilotStart);
    pilot.stop(pilotStart + 1.5);
}