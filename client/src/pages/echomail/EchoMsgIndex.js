
import Hotkeys from 'react-hot-keys';

import { useSelector, useDispatch } from 'react-redux';
import { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { Header } from '../../common/Header';
import { eventBus } from '../../EventBus.js';
import { Row } from './Row.js';

export const EchoMsgIndex = (props) => {

    const areas = useSelector((state) => state.areas) ?? [];

    useEffect(() => {
        eventBus.invoke({
            type: 'ECHO_INDEX',
        });
    }, []);

    const { echoTag } = useParams();
    console.log(`echoTag = `, echoTag);

    const area = areas.find((area) => area.area_index === echoTag);
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

    return (
        <>

            <Header />

            <Hotkeys
                keyName="ctrl+left,pgup"
                onKeyUp={handlePrevMessage}
                />

            <div class="container">
                <h1>Echoarea</h1>

            </div>
        </>
    );
};
