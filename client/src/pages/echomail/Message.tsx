
import { useEffect, useState } from 'react';

import './Message.css';
import { useInput } from '../../Hotkey';
import { useParams } from 'react-router';
import { Area, Message as MessageModel } from './EchoMsgIndex';
import { useSelector } from 'react-redux';

function detectLineEnding(content: string): "win" | "unix" | "mac" | undefined {
    if (/\r\n/.test(content)) {
        return 'win';
    } else if (/\n/.test(content)) {
        return 'unix';
    } else if (/\r/.test(content)) {
        return 'mac';
    }
    return undefined;
}

/**
 * Разбираем содержимое сообщения на строки
 * 
 * @param content Содержимое сообщения
 * @returns Строки сообщения
 */
function parseLines(content: string): string[] {
    const lineEnding = detectLineEnding(content);
    console.log(`Использован режим ${lineEnding} перевода строк`);
    if (lineEnding === 'win') {
        return content.split(/\r\n/);
    } else if (lineEnding === 'unix') {
        return content.split(/\n/);
    } else if (lineEnding === 'mac') {
        return content.split(/\r/);
    } else {
        console.log(`Подозрительное содержимое "${JSON.stringify(content)}"`);
        return [content ?? ''];
    }
};

const STATE_START = 0;
const STATE_QUOTE = 1;
const STATE_BODY = 2;

class QuoteOwner {
    public value: string;

    constructor() {
        this.value = '';
    }

    register(ch: string) {
        if (this.value.length === 0) {
            if (ch === ' ') {
            } else {
                this.value += ch;
            }
        } else {
            this.value += ch;
        }
    }

    empty() {
        return this.value === '';
    }

    reset() {
        this.value = '';
    }

}

const parseQuote = (row: string): {
    row: string,
    count: number,
    who: QuoteOwner,
    msg: string,
} => {
    let state = STATE_START;
    let who = new QuoteOwner();
    let msg = '';
    let count = 0;
    for (const ch of row) {
        if (state === STATE_START) {
            if (ch === '>') {
                count += 1;
                state = STATE_QUOTE;
                msg = '';
            } else {
                who.register(ch);
                msg += ch;
            }
        } else if (state === STATE_QUOTE) {
            if (ch === '>') {
                count += 1;
            } else {
                state = STATE_BODY;
                msg += ch;
            }
        } else if (state === STATE_BODY) {
            msg += ch;
        }
    }
    if (state === STATE_START) {
        who.reset();
    }

    const result = {
        row,
        count,
        who,
        msg,
    };
    return result;
};

const makeQuoteChar = (count: number) => {
    let result = '';
    for (let i = 0; i < count; i++) {
        result += '>';
    }
    return result;
}

const makeClassNameByCounter = (counter: number) => {

    if (counter === 0) {
        return '';
    }
    if (counter % 2 === 0) {
        return 'rowRed';
    } else {
        return 'rowGreen';
    }

};

const Quoting = ({ who, qp, msg }: { who: any, qp: any, msg: any }) => {
    return (
        <div className="rowQuote">
            {who.value}{qp} {msg === '' ? <br /> : msg}
        </div>
    );
}

const Line = ({ value, index, active }: { value: string, index: number, active: boolean }) => {

    const quote = parseQuote(value);
    const { who, count, msg } = quote;

    const qp = makeQuoteChar(count);
    const is_quote = !who.empty();

    const colorClass: string[] = [
        'messageLine',
        makeClassNameByCounter(quote.count),
    ];
    if (active) {
        colorClass.push('rowActive');
    }

    const lineLngth: number = value.length;
    const numberLine: number = Math.floor(lineLngth / 80) + 1; // Предполагаем 80 символов в строке
    const rowHeight: number = 14 * numberLine;

    return (
        <div className={colorClass.join(' ')} style={{ height: `${rowHeight}pt` }}>
            {is_quote ? <Quoting who={who} qp={qp} msg={msg} /> : value}
        </div>
    );
};

export interface MsgView {
    area: {
        name: string,
    },
    echo: {
        from: string, /* Пользователь отправивший сообщение */
        to: string,   /* Пользователь получивший сообщение */
        subject: string, /* Тема сообщения */
        date: string,   /* Дата сообщения */
    },
    body: string,
}

export const Message = () => {

    const view: MsgView = useSelector((state: any) => state.view) ?? {};
    const { body, area, echo } = view;

    const { echoTag, msgId } = useParams();
    console.log(`echoTag = ${echoTag} msgId = ${msgId}`);

    const [state, setState] = useState<{ records: string[], maxLine: number, line: number }>({
        records: [],
        maxLine: 0,
        line: 0,
    });

    const handlePreviousMessage = () => {
        console.log(`Переход на предыдущее сообщение.`);
    }
    const handleNextMessage = () => {
        console.log(`Переход на следующее сообщение.`);
    }

    useEffect(() => {
        const removeHandler = useInput((event: KeyboardEvent) => {
            if (event.key === 'ArrowLeft') {
                handlePreviousMessage();
                event.stopPropagation();
            }
            if (event.key === 'ArrowRight') {
                handleNextMessage();
                event.stopPropagation();
            }
        });
        return () => {
            removeHandler();
        }
    }, []);

    useEffect(() => {
        const records: string[] = parseLines(body);
        setState((prev) => ({
            ...prev,
            records,
            maxLine: records.length,
            line: 0,
        }));
    }, [body]);

    console.log(`[A] Текущая строка ${state.line} всего строк в тексте ${state.maxLine}`);

    const handlePrevLine = () => {
        console.log(`[B] Текущая строка ${state.line} всего строк в тексте ${state.maxLine}`);
        setState((prev) => ({
            ...prev,
            line: prev.line > 0 ? prev.line - 1 : prev.line,
        }));
    };

    const handleNextLine = () => {
        console.log(`[B] Текущая строка ${state.line} всего строк в тексте ${state.maxLine}`);
        setState((prev) => ({
            ...prev,
            line: prev.line < prev.maxLine ? prev.line + 1 : prev.line,
        }));
    };

    useEffect(() => {
        const removeHotkeys = useInput((event: KeyboardEvent) => {
            if (event.key === 'ArrowUp') {
                handlePrevLine();
            }
            if (event.key === 'ArrowDown') {
                handleNextLine();
            }
        });
        return () => {
            removeHotkeys();
        }
    }, []);

    return (
        <>
            <div className="echo-msg-view-header-wrapper">
                <table className="echo-msg-view-header">
                    <tbody><tr className="" title="">
                        <td className="echo-msg-view-header-name">
                            <span className="">Area:</span>
                        </td>
                        <td className="echo-msg-view-header-value">
                            <span className="">{ area?.name ?? '-' }</span>
                        </td>
                    </tr>
                        <tr className="" title="">
                            <td className="echo-msg-view-header-name">
                                <span className="">From:</span></td>
                            <td className="echo-msg-view-header-value">
                                <span className="">{ echo?.from ?? '-' }</span></td>
                        </tr>
                        <tr className="" title="">
                            <td className="echo-msg-view-header-name">
                                <span className="">To:</span></td>
                            <td className="echo-msg-view-header-value">
                                <span className="">{ echo?.to ?? '-' }</span></td>
                        </tr>
                        <tr className="" title="">
                            <td className="echo-msg-view-header-name">
                                <span className="">Subject:</span></td>
                            <td className="echo-msg-view-header-value">
                                <span className="">{ echo?.subject ?? '-' }</span></td>
                        </tr>
                        <tr className="" title="">
                            <td className="echo-msg-view-header-name">
                                <span className="">Date:</span></td>
                            <td className="echo-msg-view-header-value">
                                <span className="">{ echo?.date ?? '-' }</span></td>
                        </tr>
                    </tbody></table>
            </div>
            <div className="echo-msg-view-body">
                {state.records.map((row, index) => (<Line key={index} active={index === state.line} index={index} value={row} />))}
            </div>
        </>
    );
}