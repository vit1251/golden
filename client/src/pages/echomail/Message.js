
import { useState, useEffect } from 'react';
import { useNavigate } from "react-router-dom";

import Hotkeys from 'react-hot-keys';

import './Message.css';

const parseLines = (content) => {
    const result = [];
    let row = '';
    for (const ch of content) {
        if (ch === '\r') {
           result.push(row);
           row = '';
        } else {
            row += ch;
        }
    }
    return result;
};

const STATE_START = 0;
const STATE_QUOTE = 1;
const STATE_BODY  = 2;

class QuoteOwner {
    constructor() {
        this.value = '';
    }
    register(ch) {
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

const parseQuote = (row) => {
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

const makeQuoteChar = (count) => {
    let result = '';
    for (let i=0; i < count; i++) {
        result += '>';
    }
    return result;
}

const makeClassNameByCounter = (counter) => {

    if (counter === 0) {
        return '';
    }
    if (counter % 2 === 0) {
        return 'rowRed';
        } else {
        return 'rowGreen'; }

};

export const Message = (props) => {

    const [line, setLine] = useState(0);

    const { body } = props;
    const rows = parseLines(body);
    const msgLineCount = rows.length;

    const handlePrevLine = () => {
        if (line > 0) {
            setLine(line - 1);
        }
    };
    const handleNextLine = () => {
        if ((line + 1) < msgLineCount) {
            setLine(line + 1);
        }
    };

    return (
        <>


            <Hotkeys keyName="up" onKeyDown={handlePrevLine} />
            <Hotkeys keyName="down" onKeyDown={handleNextLine} />

            {rows.map((row, index) => {
                const quote = parseQuote(row);
                const { who, count, msg } = quote;
                const qp = makeQuoteChar(count);
                const is_quote = !who.empty();
                return (
                   <div key={`msg-line-${index}`} className={index === line ? 'rowActive' : ''}>
                     <div className={makeClassNameByCounter(quote.count) } data-tooltip={JSON.stringify(quote)}>{ is_quote ? (
                        <>
                            <div className="rowQuote">
                                {who.value}{qp} {msg === '' ? <br /> : msg }
                            </div>
                        </>
                     ) : (
                        <>
                            <div>{row === '' ? <br /> : row }</div>
                        </>
                     )}</div>
                   </div>
                );
            })}
        </>
    );
}