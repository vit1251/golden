
import Hotkeys from 'react-hot-keys';

import { useSelector, useDispatch } from 'react-redux';
import { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { Header } from '../../common/Header';
import { eventBus } from '../../EventBus.js';
import { Row } from './Row.js';

import "./EchoIndex.css";

export const EchoIndex = (props) => {

    const areas = useSelector((state) => state.areas) ?? [];

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

            <Hotkeys
                keyName="ctrl+left,pgup"
                onKeyUp={handlePrevMessage}
                />

            <div className="container">
                <h1>Echomail</h1>

                <Row 
                     onRowLink={(row) => `/echomail/${row.area_index}`}
                     columns={[
                        { className: "rowName", key: "name" },
                        { className: "rowMarker", render: (row) => {
                            const { new_message_count = 0 } = row;
                            const value = new_message_count > 0 ? '•' : null;
                            return value;
                        }},
                        { className: "rowSummary", key: "summary" },
                        { className: "rowCounter", render: (row) => {
                            const { new_message_count = 0 } = row;
                            const value = new_message_count > 0 ? new_message_count : null;
                            return value;
                        }},
                     ]}
                     data={areas} />
            </div>
        </>
    );
};