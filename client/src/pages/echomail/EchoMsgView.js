
import { useNavigate } from "react-router-dom";
import { useState, useEffect } from 'react';
import { useParams } from "react-router-dom";
import { useSelector } from 'react-redux';

import Hotkeys from 'react-hot-keys';

import { eventBus } from '../../EventBus.js';
import { Header } from '../../common/Header';
import { Message } from './Message.js';

export const EchoMsgView = (props) => {

    const navigate = useNavigate();

    const { body } = useSelector((state) => state.message);

    const { echoTag, msgId } = useParams();
    console.log(echoTag);

    useEffect(() => {
        eventBus.invoke({
            type: 'ECHO_MSG_VIEW',
            echoTag,
            msgId,
        });
    }, [echoTag, msgId]);

    const handleMsgIndex = () => {
        console.log(`Back on message index..`);
        navigate(`/echomail/${echoTag}`);
    };

    const handleMsgRemove = () => {
        /* Step 1. Remove message */
        eventBus.invoke({
            type: 'ECHO_MSG_REMOVE',
            echoTag,
            msgId,
        });
        /* Step 2. Message index */
        navigate(`/echomail/${echoTag}`);
    };

    return (
        <>
            <Header />

            <div class="container">
                <h1>EchoMailView</h1>

                <Message body={body} />

            </div>

            <Hotkeys keyName="esc" onKeyDown={handleMsgIndex} />
            <Hotkeys keyName="del" onKeyDown={handleMsgRemove} />

        </>
    );
};
