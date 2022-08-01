
import Hotkeys from 'react-hot-keys';

import { useSelector, useDispatch } from 'react-redux';
import { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { Header } from '../common/Header';
import { eventBus } from '../EventBus.js';
import { Row } from './Row.js';

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

            <div class="container">
                <h1>Echomail</h1>

                <Row areas={areas} />
            </div>
        </>
    );
};