
import { useSelector } from 'react-redux';
import { useEffect } from 'react';

import { Header } from '../../common/Header';
import { eventBus } from '../../EventBus';
import { Rows } from './Row';

import "./EchoIndex.css";
import { Area } from './EchoMsgIndex';

export const EchoIndex = () => {

    const areas: Area[] = useSelector((state: any) => state.areas) ?? [];

    useEffect(() => {
        /* Step 1. Ask echos */
        eventBus.invoke({
            type: 'ECHO_INDEX',
        });
    }, []);

    const handlePrevMessage = () => {
        console.log(`handlePrevMessage...`);
    };

    return (
        <div>
            <h1>Echomail</h1>

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