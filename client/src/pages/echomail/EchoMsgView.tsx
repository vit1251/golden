
import { useNavigate, useParams } from "react-router";
import { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { Message as MessageComponent } from './Message';
import { Message } from "../../models/Message.model";
import { Area } from "../../models/Area.model";

export const EchoMsgView = () => {
    const dispatch = useDispatch();

    const sendMessage = (payload: any) => {
        dispatch({
            type: 'SOCKET_SEND',
            payload: payload,
        });
    };

    const navigate = useNavigate();

    const areas: Array<Area> = useSelector((state: any) => state.areas.records);
    const messages: Array<Message> = useSelector((state: any) => state.messages.records);
    const content: string = useSelector((state: any) => state.view.content);

    const { echoTag, msgId } = useParams();
    console.log(echoTag);

    useEffect(() => {
        sendMessage({
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
        // Шаг 1. Отправляем команду удаления сообщения
        sendMessage({
            type: 'ECHO_MSG_REMOVE',
            echoTag,
            msgId,
        });
        // Шаг 2. Возвращаемся в директорию телеконференций
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
        <div>
            <MessageComponent />
        </div>
    );
};
