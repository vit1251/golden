
import { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useParams } from 'react-router';
import { useInput } from '../../Hotkey';

import './Message.css';

const Body = ({ rawText }: { rawText: string }) => {
    const text: string = rawText.replace(/\r\n|\r|\n/g, '\n');
    const lines = text.split('\n');
    const records: Array<any> = [];

    for (const line of lines) {
        
        // Шаг 1. Обработка цитат
        const match = line.match(/^(\s+[A-Za-zА-Яа-я0-9]{1,3}>{1,3}|\s+[>]{1,3}>)/);
        if (match) {
            const prefix = match[0];
            // Считаем количество знаков ">", чтобы понять уровень вложенности
            const level: number = (prefix.match(/>/g) || []).length;
            // Ограничиваем уровень (например, до 3), чтобы не плодить бесконечные стили
            const colorClass: string[] = [
                `msg-line`,
                `msg-quote-${Math.min(level, 3)}`,
            ];
            
            records.push( <div className={colorClass.join(' ')}>{line}</div> );
            continue;
        }

        if (line.startsWith('---') || line.startsWith(' * Origin:')) {
            const colorClass: string[] = [
                'msg-line',
                'msg-service',
            ];
            records.push( <div className={colorClass.join(' ')}>{line}</div> );
            continue;
        }

        // Обычный текст автора
        const colorClass: string[] = [
            'msg-line',
            'msg-plain',
        ];
        records.push( <div className={colorClass.join(' ')}>{line}</div> );
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

    const msgArea: string = useSelector((state: any) => state.view.echo);
    const msgFrom: string = useSelector((state: any) => state.view.from);
    const msgTo: string = useSelector((state: any) => state.view.to);
    const msgSubject: string = useSelector((state: any) => state.view.subject);
    const msgDate: string = useSelector((state: any) => state.view.date);
    const content: string = useSelector((state: any) => state.view.content);

    const { echoTag, msgId } = useParams();
    console.log(`echoTag = ${echoTag} msgId = ${msgId}`);

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



    return (
        <div>
            <div className="echo-msg-view-header-wrapper">
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
            <div className="echo-msg-view-body">
                <Body rawText={content} />
            </div>
        </div>
    );
}