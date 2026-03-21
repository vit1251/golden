
import { useNavigate, useParams } from "react-router";
import { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { Message as MessageComponent } from './Message';
import { Message } from "../../models/Message.model";
import { Area } from "../../models/Area.model";
import { useInput } from "../../Hotkey";
import { playError } from "../../Audio";


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

    const handleBack = () => navigate(`/echo/${echoTag}`);
    const handleMsgRemove = () => {
        sendMessage({ type: 'ECHO_MSG_REMOVE', echoTag, msgId });
        handleBack();
    };

    const handlePrevMessage = () => {
        playError();
    };
    const handleNextMessage = () => {
        playError();
    };

    useEffect(() => {
        const removeHotkeys = useInput((event: KeyboardEvent) => {
            if (event.key === 'Escape') handleBack();
            if (event.key === 'Delete') handleMsgRemove();
            if (event.key === 'ArrowLeft') handlePrevMessage();
            if (event.key === 'ArrowRight') handleNextMessage();
        });
        return () => removeHotkeys();
    }, []);

    return (
        <div className="Page Page-View">
            <MessageComponent />
        </div>
    );
};
