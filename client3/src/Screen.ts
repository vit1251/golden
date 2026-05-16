
export class Screen {
    readonly matrix: string[][];
    readonly back: string[][];
    readonly front: string[][];
    protected fg: string = '#C0C0C0';
    protected bg: string = '#000000';

    constructor(readonly width: number = 80, readonly height: number = 25) {
        this.matrix = Array.from({ length: height }, () => new Array(width).fill(' '));
        this.back = Array.from({ length: height }, () => new Array(width).fill('#000000'));
        this.front = Array.from({ length: height }, () => new Array(width).fill('#C0C0C0'));
    }

    setBackgroundColor(color: string) {
        this.bg = color;
    }

    setForegroudColor(color: string) {
        this.fg = color;
    }

    reset({ foreGround = '#C0C0C0', backGround = '#000000' }: { foreGround?: string , backGround?: string } = {}) {
        // Шаг 2. Сбросили цвета        
        this.fg = foreGround;
        this.bg = backGround;
        // Шаг 1. Буфер обнуляем
        for (let posY = 0; posY < this.height; posY++) {
            for (let posX = 0; posX < this.width; posX++) {
                this.set(posX, posY, ' ');
            }
        }
    }

    set(x: number, y: number, ch: string) {
        if ((y < this.height) && (x < this.width)) {
            this.matrix[y]![x] = ch.length === 1 ? ch : '?';
            this.front[y]![x] = this.fg;
            this.back[y]![x] = this.bg;
        }
    }

    drawLine(y: number) {
        for (let pos = 0; pos < this.width; pos++) {
            this.set(pos, y, '─');
        }
    }

    drawRect(x: number, y: number, width: number, height: number) {
        this.set(x, y, '┌');
        this.set(x + width, y, '┐');
        this.set(x, y + height, '└');
        this.set(x + width, y + height, '┘');

        for (let pos = 1; pos < width; pos++) {
            this.set(x + pos, y, '─');
            this.set(x + pos, y + height, '─');
        }
        for (let pos = 1; pos < height; pos++) {
            this.set(x, y + pos, '│');
            this.set(x + width, y + pos, '│');
        }
    }

    writeText(x: number, y: number, text: string) {
        for (let i = 0; i < text.length; i++) {
            if (x + i < this.width) {
                const ch: string = text[i] ?? ' ';
                this.set(x + i, y, ch);
            }
        }
    };
}
