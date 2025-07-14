
import { useParams, useNavigate } from "react-router";
import { useSelector } from 'react-redux';
import { useEffect } from 'react';

import { eventBus } from '../../EventBus';
import { Rows } from './Row';

import "./EchoMsgIndex.css";
import { adjustBrightness, makeShort, stringToHexColor } from "../../usils";

export interface Area {
    name: string                   /* Имя конференции. Пример "RU.ANEKDOT" */
    summary: string; 
    message_count: number;
    new_message_count: number;
    order: number;
    area_index: string;
}

export interface Message {
    from: string;
    to: string;
    view_count: number;
    subject: string;
    date: string;
    hash: string;
}

export const EchoMsgIndex = () => {

    const navigate = useNavigate();

    const areas = useSelector((state: any) => state.areas);
    const messages = useSelector((state: any) => state.messages);

    useEffect(() => {
        eventBus.invoke({
            type: 'ECHO_INDEX',
        });
    }, []);

    const { echoTag } = useParams();
    console.log(`echoTag = `, echoTag);

    const area: Area = areas.find((area: Area) => area.area_index === echoTag);
    console.log(`area = `, area);

    useEffect(() => {
        /* Step 1. Ask echos */
        eventBus.invoke({
            type: 'ECHO_MSG_INDEX',
            echoTag,
        });
        eventBus.invoke({
            type: 'SUMMARY',
        });
    }, [echoTag]);

    const handlePrevMessage = () => {
        console.log(`handlePrevMessage...`);
    };
    const handleNextMessage = () => {
        console.log(`handlePrevMessage...`);
    };
    const handleAreaIndex = () => {
        navigate(`/echo`);
    };
    const handleCreateMessage = () => {
        navigate(`/echo/${echoTag}/create`);
    };

    return (
        <div>
            <h1>Echoarea</h1>

            <Rows<Message>
                    onRowLink={(row: Message): string => {
                        const { hash } = row;
                        if (hash) {
                            return `/echo/${echoTag}/${hash}/view`;
                        } else {
                            return `#`;
                        }
                    }}
                    columns={[
                       {className: 'rowUserpic', styles: (row: Message) => {
                            const color: string = stringToHexColor(`${row.from}`);
                            const darkColor: string = adjustBrightness(color, 0.5);
                            return {
                                backgroundColor: darkColor,
                            }
                       }, render: (row: Message): string => makeShort(row.from)},
                       {className: 'rowFrom', key: 'from'},
                       {className: 'rowMarker', render: (row: Message): string => {
                           const { view_count = 0 } = row;
                           const value = view_count === 0 ? '•' : '';
                           return value;
                       }},
                       {className: 'rowSubject', key: 'subject'},
                       {className: 'rowDate', key: 'date'},
                    ]}
                    records={messages}
                    />

        </div>
    );
};
