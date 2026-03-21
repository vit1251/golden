
import { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useNavigate, useParams } from 'react-router';
import { useInput } from '../../Hotkey';

import './Message.css';
import { playModemBeep, playSequence } from '../../Audio';

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

const Body = ({ rawText }: { rawText: string }) => {
    const text: string = rawText.replace(/\r\n|\r|\n/g, '\n');
    const lines = text.split('\n');
    const records: Array<any> = [];

    for (const line of lines) {

        // Шаг 1. Обработка цитат
        const match = line.match(/^(\s*[A-Za-zА-Яа-я0-9]{1,3}>{1,3}|\s+[>]{1,3}>)/);
        if (match) {
            const prefix = match[0];
            // Считаем количество знаков ">", чтобы понять уровень вложенности
            const level: number = (prefix.match(/>/g) || []).length;
            // Ограничиваем уровень (например, до 3), чтобы не плодить бесконечные стили
            const colorClass: string[] = [
                `msg-line`,
                `msg-quote-${Math.min(level, 3)}`,
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
        <div>
            {records}
        </div>
    );
}

export const Message = () => {
    const dispatch = useDispatch();

    const sendMessage = (payload: any) => {
        dispatch({
            type: 'SOCKET_SEND',
            payload: payload,
        });
    };

    const navigate = useNavigate();

    const msgArea: string = useSelector((state: any) => state.view.echo);
    const msgFrom: string = useSelector((state: any) => state.view.from);
    const msgTo: string = useSelector((state: any) => state.view.to);
    const msgSubject: string = useSelector((state: any) => state.view.subject);
    const msgDate: string = useSelector((state: any) => state.view.date);
    const content: string = useSelector((state: any) => state.view.content);

    const { echoTag, msgId } = useParams();
    console.log(`echoTag = ${echoTag} msgId = ${msgId}`);



    const handleSound = (type: 'SND_SAYBIBI' | 'SND_THEEND' | 'SND_GOTIT' | 'SND_TOOBAD' | 'SND_TOYOU' = 'SND_TOYOU') => {
        playModemBeep();
        if (type === 'SND_SAYBIBI') {
            playSequence([
                [440, 111],
            ]);
        }
        if (type === 'SND_TOYOU') {
            playSequence([
                [100, 18],
                [500, 18],
                [100, 18],
            ]);
        }
        if (type === 'SND_THEEND') {
            playSequence([
                [220, 111],
                [110, 167],
            ]);
        }
        if (type === 'SND_GOTIT') {
            playSequence([
                [110, 56],
                [220, 56],
                [110, 56],
                [220, 111],
            ]);
        }
        if (type === 'SND_TOOBAD')
            playSequence([
                [440, 111],
                [220, 111],
                [110, 167],
            ]);


    };
    

    const handleBack = () => {
        navigate(`/echo/${echoTag}`);
    };
    const handlePreviousMessage = () => {
        console.log(`Переход на предыдущее сообщение.`);
    }
    const handleNextMessage = () => {
        console.log(`Переход на следующее сообщение.`);
    }

    useEffect(() => {
        sendMessage({
            type: 'ECHO_MSG_VIEW',
            echoTag,
            msgId,
        });
    }, []);

    useEffect(() => {
        const removeHandler = useInput((event: KeyboardEvent) => {
            if (event.key === 'Escape') handleBack();
            if (event.key === 'ArrowLeft') handlePreviousMessage();
            if (event.key === 'ArrowRight') handleNextMessage();
            /* Home */
            if (event.key === 'Home') handleSound();
        });
        return () => removeHandler();
    }, []);



    return (
        <div className="Page-View">
            <div className="View-Header echo-msg-view-header-wrapper">
                <table className="echo-msg-view-header">
                    <tbody><tr className="" title="">
                        <td className="echo-msg-view-header-name">
                            <span className="">Area:</span>
                        </td>
                        <td className="echo-msg-view-header-value">
                            <span className="">{msgArea}</span>
                        </td>
                    </tr>
                        <tr className="" title="">
                            <td className="echo-msg-view-header-name">
                                <span className="">From:</span></td>
                            <td className="echo-msg-view-header-value">
                                <span className="">{msgFrom}</span></td>
                        </tr>
                        <tr className="" title="">
                            <td className="echo-msg-view-header-name">
                                <span className="">To:</span></td>
                            <td className="echo-msg-view-header-value">
                                <span className="">{msgTo}</span></td>
                        </tr>
                        <tr className="" title="">
                            <td className="echo-msg-view-header-name">
                                <span className="">Subject:</span></td>
                            <td className="echo-msg-view-header-value">
                                <span className="">{msgSubject ?? '-'}</span></td>
                        </tr>
                        <tr className="" title="">
                            <td className="echo-msg-view-header-name">
                                <span className="">Date:</span></td>
                            <td className="echo-msg-view-header-value">
                                <span className="">{msgDate}</span></td>
                        </tr>
                    </tbody></table>
            </div>
            <div className="View-Body echo-msg-view-body">
                <Body rawText={content} />
            </div>
        </div>
    );
}