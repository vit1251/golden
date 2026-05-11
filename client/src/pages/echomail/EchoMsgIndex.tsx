
import { useParams, useNavigate } from "react-router";
import { useDispatch, useSelector } from 'react-redux';
import { useCallback, useEffect } from 'react';

import { Rows } from './Row.tsx';

import "./EchoMsgIndex.css";

import { adjustBrightness, makeShort, stringToHexColor } from "../../usils.ts";
import { type Area } from "../../models/Area.model.ts";
import { type Message } from "../../models/Message.model.ts";
import { useKeyboard } from "../../Hotkey.tsx";
import { firstMessage, lastMessage, nextMessage, prevMessage } from "../../features/messageSlice.ts";
import { type RootState } from "../../app/store.ts";
import { socketSend } from "../../middleware/socketMiddleware.ts";
import { soundEvent } from "../../middleware/soundMiddleware.ts";


export const EchoMsgIndex = () => {
    const dispatch = useDispatch();

    const navigate = useNavigate();

    const { echoTag } = useParams();
    console.log(`echoTag = `, echoTag);

    const areas: Array<Area> = useSelector((state: RootState) => state.areas.records);
    const messages: Array<Message> = useSelector((state: RootState) => state.messages.records);
    const area: Area | undefined = areas.find((area: Area) => area.area_index === echoTag);
    const activeIndex: number = useSelector((state: RootState) => state.messages.activeIndex);
    const msg: Message | null = messages.at(activeIndex) ?? null;

    useEffect(() => {
        // 1. Запросим список эхо конференций
        dispatch(socketSend({
            msg: {
                type: 'ECHO_INDEX',
            },
        }));
    }, []);
    useEffect(() => {
        dispatch(socketSend({
            msg: {
                type: 'ECHO_MSG_INDEX',
                echoTag: echoTag,
            }
        }));
    }, [echoTag]);

    console.log(`area = `, area);

    //    const handleCreateMessage = () => {
    //        navigate(`/echo/${echoTag}/create`);
    //    };
    const handleFirstMessage = () => dispatch(firstMessage());
    const handlePrevMessage = () => {
        // Шаг 1. Проверяем в конце ли мы списка
        if (activeIndex === 0) {
            dispatch(soundEvent('SND_THEEND'));
        }
        //
        dispatch(prevMessage());
    };
    const handleNextMessage = () => {
        // Шаг 1. Проверяем в конце ли мы списка
        if (activeIndex + 1 === messages.length) {
            dispatch(soundEvent('SND_THEEND'));
        }
        //
        dispatch(nextMessage());
    };
    const handleLastMessage = () => dispatch(lastMessage());
    const handleBack = () => navigate(`/echo`);
    const handleOpenMessage = () => {
        console.log(`open`);
        if (msg) {
            navigate(`/echo/${echoTag}/${msg.hash}/view`);
        }
    };

    useKeyboard({
        Escape: () => handleBack(),
        Home: () => handleFirstMessage(),
        ArrowUp: () => handlePrevMessage(),
        ArrowDown: () => handleNextMessage(),
        End: () => handleLastMessage(),
        Enter: () => handleOpenMessage(),
    });

    return (
        <div className="Page Page-Message-Index">

            <Rows<Message>
                activeIndex={activeIndex}
                onRowLink={(row: Message): string => {
                    const { hash } = row;
                    if (hash) {
                        return `/echo/${echoTag}/${hash}/view`;
                    } else {
                        return `#`;
                    }
                }}
                columns={[
                    {
                        className: 'rowUserpic', styles: (row: Message) => {
                            const color: string = stringToHexColor(`${row.from}`);
                            const darkColor: string = adjustBrightness(color, 0.5);
                            return {
                                backgroundColor: darkColor,
                            }
                        }, render: (row: Message): string => makeShort(row.from)
                    },
                    { className: 'rowFrom', key: 'from' },
                    {
                        className: 'rowMarker', render: (row: Message): string => {
                            const { view_count = 0 } = row;
                            const value = view_count === 0 ? '•' : '';
                            return value;
                        }
                    },
                    { className: 'rowSubject', key: 'subject' },
                    { className: 'rowDate', key: 'date' },
                ]}
                records={messages}
            />

        </div>
    );
};
