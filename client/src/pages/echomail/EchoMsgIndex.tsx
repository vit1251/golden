
import { useParams, useNavigate } from "react-router";
import { useDispatch, useSelector } from 'react-redux';
import { useEffect } from 'react';

import { Rows } from './Row';

import "./EchoMsgIndex.css";
import { adjustBrightness, makeShort, stringToHexColor } from "../../usils";
import { Area } from "../../models/Area.model";
import { Message } from "../../models/Message.model";


export const EchoMsgIndex = () => {
    const dispatch = useDispatch();

    const sendMessage = (payload: any) => {
        dispatch({
            type: 'SOCKET_SEND',
            payload: payload,
        });
    };

    const navigate = useNavigate();

    const areas: Array<Area> = useSelector((state: any) => state.areas.records);
    const messages: Array<Message> = useSelector((state: any) => state.messages.records);

    const { echoTag } = useParams();
    console.log(`echoTag = `, echoTag);

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

    const area: Area | undefined = areas.find((area: Area) => area.area_index === echoTag);
    console.log(`area = `, area);

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
