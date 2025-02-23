
import { useNavigate, useParams } from "react-router";
import { useEffect } from 'react';
import { useSelector } from 'react-redux';

import { eventBus } from '../../EventBus.js';
import { Header } from '../../common/Header.js';
import { Message } from './Message';

export const EchoMsgView = () => {

    const navigate = useNavigate();

    const messages = useSelector((state: any) => state.messages);

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
        navigate(`/echo/${echoTag}`);
    };

    const handleMsgRemove = () => {
        /* Step 1. Remove message */
        eventBus.invoke({
            type: 'ECHO_MSG_REMOVE',
            echoTag,
            msgId,
        });
        /* Step 2. Message index */
        navigate(`/echo/${echoTag}`);
    };

    const handlePrevMessage = () => {
        const msgIndex = messages.findIndex((msg: any) => msg.hash === msgId);
        console.log(`Your index ${msgIndex}`);
        if ((msgIndex - 1) >= 0) {
            const { hash: prevHash } = messages[msgIndex - 1];
            navigate(`/echo/${echoTag}/${prevHash}/view`);
        } else {
            // TODO - play blump...
        }
    };
    const handleNextMessage = () => {
        const msgIndex = messages.findIndex((msg: any) => msg.hash === msgId);
        console.log(`Your index ${msgIndex}`);
        if ((msgIndex + 1) < messages.length) {
            const { hash: nextHash } = messages[msgIndex + 1];
            navigate(`/echo/${echoTag}/${nextHash}/view`);
        } else {
            // TODO - play blump...
        }
    };

    return (
        <>
            <Header />

            <div className="container">
                <h1>EchoMailView</h1>

                <Message />

            </div>

        </>
    );
};
