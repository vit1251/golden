
import { useParams, useNavigate } from "react-router";
import { useDispatch, useSelector } from 'react-redux';
import { useCallback, useEffect } from 'react';

import { Rows } from './Row';

import "./EchoMsgIndex.css";
import { adjustBrightness, makeShort, stringToHexColor } from "../../usils";
import { Area } from "../../models/Area.model";
import { Message } from "../../models/Message.model";
import { useInput } from "../../Hotkey";
import { firstMessage, lastMessage, nextMessage, prevMessage } from "../../features/messageSlice";
import { RootState } from "../../app/store";


export const EchoMsgIndex = () => {
    const dispatch = useDispatch();

    const sendMessage = (payload: any) => {
        dispatch({
            type: 'SOCKET_SEND',
            payload: payload,
        });
    };

    const navigate = useNavigate();

    const { echoTag } = useParams();
    console.log(`echoTag = `, echoTag);

    const areas: Array<Area> = useSelector((state: RootState) => state.areas.records);
    const messages: Array<Message> = useSelector((state: RootState) => state.messages.records);
    const area: Area | undefined = areas.find((area: Area) => area.area_index === echoTag);
    const activeIndex: number = useSelector((state: RootState) => state.messages.activeIndex);
    const msg: Message | null = messages.at(activeIndex) ?? null;

    useEffect(() => {
        sendMessage({
            type: 'ECHO_INDEX',
        });
    }, []);
    useEffect(() => {
        sendMessage({
            type: 'ECHO_MSG_INDEX',
            echoTag,
        });
    }, [echoTag]);


    console.log(`area = `, area);

    //    const handleCreateMessage = () => {
    //        navigate(`/echo/${echoTag}/create`);
    //    };
    const handleFirstMessage = () => dispatch(firstMessage());
    const handlePrevMessage = () => dispatch(prevMessage());
    const handleNextMessage = () => dispatch(nextMessage());
    const handleLastMessage = () => dispatch(lastMessage());
    const handleBack = () => navigate(`/echo`);
    const handleOpenMessage = useCallback(() => {
        console.log(`open`);
        if (msg) {
            navigate(`/echo/${echoTag}/${msg.hash}/view`);
        }
    }, [ activeIndex ]);

    useEffect(() => {
        const removeHotkeys = useInput((event: KeyboardEvent) => {
            if (event.key === 'Escape') handleBack();
            if (event.key === 'Home') handleFirstMessage();
            if (event.key === 'ArrowUp') handlePrevMessage();
            if (event.key === 'ArrowDown') handleNextMessage();
            if (event.key === 'End') handleLastMessage();
            if (event.key === 'Enter') handleOpenMessage();
            //            if (event.key === `Ctrl+C`) handleCreateMessage();
        });
        return () => removeHotkeys();
    }, [ handleOpenMessage ]);

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
