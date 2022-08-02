
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

export const Message = (props) => {
    const { body } = props;
    const rows = parseLines(body);
    return (
        <>
            {rows.map((row) => {
                const quote = parseQuote(row);
                const { who, count, msg } = quote;
                const qp = makeQuoteChar(count);
                const is_quote = !who.empty();
                return (
                   <>
                     <div className={ quote.count % 2 ? 'rowRed' : 'rowGreen' } data-tooltip={JSON.stringify(quote)}>{ is_quote ? (
                        <>
                            <div className="rowQuote">
                                {who.value}{qp} {msg}
                            </div>
                        </>
                     ) : (
                        <>
                            <div>{row}</div>
                        </>
                     )}</div>

                   </>
                );
            })}
        </>
    );
}