
import { useDispatch, useSelector } from 'react-redux';
import { useEffect } from 'react';

import { Rows } from './Row';

import "./EchoIndex.css";
import { Area } from '../../models/Area.model';
import { useInput } from '../../Hotkey';

export const EchoIndex = () => {
    const dispatch = useDispatch();

    const sendMessage = (payload: any) => {
        dispatch({
            type: 'SOCKET_SEND',
            payload: payload,
        });
    };

    const areas: Area[] = useSelector((state: any) => state.areas.records) ?? [];

    useEffect(() => {
        sendMessage({
            type: 'ECHO_INDEX',
        });
    }, []);

    const handlePrevArea = () => {
        console.log(`handlePrevMessage...`);
    };
    const handleNextArea = () => {
        console.log(`handlePrevMessage...`);
    };
    const handleOpenArea = () => {
        console.log(`openArea...`);
    };

    useEffect(() => {
        const removeHotkeys = useInput((event: KeyboardEvent) => {
            if (event.key === 'ArrowUp') handlePrevArea();
            if (event.key === 'ArrowDown') handleNextArea();
            if (event.key === 'Enter') handleOpenArea();
        });
        return () => removeHotkeys();
    }, []);

    return (
        <div className="Page Page-Areas">

            <Rows<Area>
                onRowLink={(row: Area) => `/echo/${row.area_index}`}
                columns={[
                    {
                        className: "rowMarker", render: (row: any): string => {
                            const { new_message_count = 0 } = row;
                            //const value = new_message_count > 0 ? '•' : '';
                            const value = new_message_count > 0 ? '+' : '';
                            return value;
                        }
                    },
                    { className: "rowSummary", key: "summary" },
                    {
                        className: "rowMessageCount", render: (row: Area): string => {
                            const { message_count = 0 } = row;
                            return `${message_count}`;
                        }
                    },
                    {
                        className: "rowMessageNewCount", render: (row: Area): string => {
                            const { new_message_count = 0 } = row;
                            return `${new_message_count}`;
                        }
                    },
                    { className: "rowName", key: "name" },
                ]}
                records={areas} />

        </div>
    );
};