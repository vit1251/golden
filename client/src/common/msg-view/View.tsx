
import './View.css';

function checkInternet(): boolean {
    return false;
}
function checkEmail(): boolean {
    //     /\b([a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,})\b/ig
    return false;
}
function checkFido(): boolean {
    //     /\b([1-6]:\d{1,5}\/\d{1,5}(\.\d{1,5})?)\b/g
    return false;
}
function checkFido2(): boolean {
    //      /\b(area:\/\/[A-Z0-9.-]+)\b/ig
    return false;
}

function chunkTest(chunk: string): number {
    // Шаг 1. Проверка на Интернет ссылку
    if (chunk.startsWith('https://') || chunk.startsWith('http://')) return 1;

    return 0;
}

class RegExpBundler {
    protected expressions: Array<string> = [];
    addExpression(exp: string): this {
        this.expressions.push(`(${exp})`);
        return this;
    }
    build(): RegExp {
        const regexp: string = this.expressions.join('|');
        return new RegExp(`(?:${regexp})`, 'ig');
    }
}

const Line = ({ line }: { line: string }) => {
    const regExp: RegExp = (new RegExpBundler())
        .addExpression('area:\/\/[A-Za-z0-9.-]+')
        .addExpression('\b[1-6]:\d{1,5}\/\d{1,5}(?:\.\d{1,5})?\b')
        .build();

    const chunks = line.split(regExp);
    console.log(chunks);

    return (
        <>
            {chunks.map((chunk, i) => {
                if (!chunk) return null; // Убираем пустые куски от split

                // Проверяем, является ли часть текста ссылкой
                const chunkType: number = chunkTest(chunk);
                if (chunkType === 1) {
                    return (
                        <a key={i} href={chunk} target="_blank" rel="noreferrer" className="msg-link">
                            {chunk}
                        </a>
                    );
                }

                // Обычный текст
                return <span key={i}>{chunk}</span>;
            })}
        </>
    );
};

class Paragraph {
    protected constructor(readonly line: string, readonly prefix: string, readonly level: number) {
    }
    static parse(line: string) {
        const match = line.match(/^(\s*[A-Za-zА-Яа-я0-9]{1,3}>{1,3}|\s+[>]{1,3}>)/);
        if (match) {
            const prefix = match[0];
            // Считаем количество знаков ">", чтобы понять уровень вложенности
            const level: number = (prefix.match(/>/g) ?? []).length;
            return new Paragraph(line, prefix, level);
        }
        return new Paragraph(line, '', 0);
    }
}

export const View = ({ rawText }: { rawText: string }) => {
    const text: string = rawText.replace(/\r\n|\r|\n/g, '\n');
    const lines = text.split('\n');
    const records: Array<any> = [];

    for (const line of lines) {

        // Шаг 1. Обработка цитат
        const p: Paragraph = Paragraph.parse(line);
        if (p.level > 0) {
            const colorClass: string[] = [
                `msg-line`,
                `msg-quote-${Math.min(p.level, 3)}`,
            ];
            records.push(<div className={colorClass.join(' ')}><Line line={line} /></div>);
            continue;
        }

        if (line.startsWith('---') || line.startsWith(' * Origin:')) {
            const colorClass: string[] = [
                'msg-line',
                'msg-service',
            ];
            records.push(<div className={colorClass.join(' ')}><Line line={line} /></div>);
            continue;
        }

        // Обычный текст автора
        const colorClass: string[] = [
            'msg-line',
            'msg-plain',
        ];
        records.push(<div className={colorClass.join(' ')}><Line line={line} /></div>);
    }

    return (
        <div className="msg">
            {records}
        </div>
    );
}