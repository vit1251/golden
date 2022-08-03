
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

    return (
        <>
            <Header />

            <div class="container">
                <h1>EchoMailView</h1>

                <Message body={body} />

            </div>

            <Hotkeys keyName="esc" onKeyDown={handleMsgIndex} />

        </>
    );
};
