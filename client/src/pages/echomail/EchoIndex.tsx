
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
        eventBus.invoke({
            type: 'SUMMARY',
        });
    }, []);

    const handlePrevMessage = () => {
        console.log(`handlePrevMessage...`);
    };

    return (
        <>

            <Header />

            <div className="container">
                <h1>Echomail</h1>

                <Rows<Area>
                     onRowLink={(row: Area) => `/echo/${row.area_index}`}
                     columns={[
                        { className: "rowName", key: "name" },
                        { className: "rowMarker", render: (row: any): string => {
                            const { new_message_count = 0 } = row;
                            const value = new_message_count > 0 ? '•' : '';
                            return value;
                        }},
                        { className: "rowSummary", key: "summary" },
                        { className: "rowCounter", render: (row: Area): string => {
                            const { new_message_count = 0 } = row;
                            const value = new_message_count > 0 ? `${new_message_count}` : '';
                            return value;
                        }},
                     ]}
                     records={areas} />
                     
            </div>
        </>
    );
};