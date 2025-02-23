
import { useSelector } from 'react-redux';
import { useEffect } from 'react';

import { Header } from '../../common/Header';
import { eventBus } from '../../EventBus';
import { Rows } from './Row';

import "./EchoIndex.css";

export const EchoIndex = () => {

    const areas = useSelector((state: any) => state.areas) ?? [];

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

                <Rows
                     onRowLink={(row: any) => `/echomail/${row.area_index}`}
                     columns={[
                        { className: "rowName", key: "name" },
                        { className: "rowMarker", render: (row: any) => {
                            const { new_message_count = 0 } = row;
                            const value = new_message_count > 0 ? '•' : null;
                            return value;
                        }},
                        { className: "rowSummary", key: "summary" },
                        { className: "rowCounter", render: (row: any) => {
                            const { new_message_count = 0 } = row;
                            const value = new_message_count > 0 ? new_message_count : null;
                            return value;
                        }},
                     ]}
                     records={areas} />
                     
            </div>
        </>
    );
};