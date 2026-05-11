
import { useNavigate, useParams } from "react-router";
import { useCallback, useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { Message as MessageComponent } from './Message.tsx';
import { type Message } from "../../models/Message.model.ts";
import { type Area } from "../../models/Area.model.ts";
import { useKeyboard } from "../../Hotkey.tsx";
import { socketSend } from "../../middleware/socketMiddleware.ts";
import { soundEvent } from "../../middleware/soundMiddleware.ts";


export const EchoMsgView = () => {
    const dispatch = useDispatch();

    const navigate = useNavigate();

    const areas: Array<Area> = useSelector((state: any) => state.areas.records);
    const messages: Array<Message> = useSelector((state: any) => state.messages.records);
    const content: string = useSelector((state: any) => state.view.content);
    const msgTo: string = useSelector((state: any) => state.view.to);

    const yours: string[] = [ 'All' ]; // Варианты написания имени

    const { echoTag, msgId } = useParams();
    console.log(echoTag);

    useEffect(() => {
        dispatch(socketSend({
            msg: {
                type: 'ECHO_MSG_VIEW',
                echoTag: echoTag,
                msgId: msgId,
            },
        }));
    }, [echoTag, msgId]);

    useEffect(() => {
        if (yours.includes(msgTo)) {
            dispatch(soundEvent('SND_TOYOU'));
        }
    }, [ msgTo ]);

    const handleMsgRemove = useCallback(() => {
        // Шаг 1. Удаляем сообщение
        dispatch(socketSend({
            msg: {
                type: 'ECHO_MSG_REMOVE',
                echoTag: echoTag,
                msgId: msgId,
            },
        }));
        // Шаг 2. Возвращаемся на список сообещний
        handleBack();
    
    }, [ echoTag, msgId ]);

    /* Home */
    //if (event.key === '1') playMusic('SND_SAYBIBI');
    //if (event.key === '2') playMusic('SND_TOYOU');
    //if (event.key === '3') playMusic('SND_THEEND');
    //if (event.key === '4') playMusic('SND_GOTIT');
    //if (event.key === '5') playMusic('SND_TOOBAD');


    const activeIndex: number = 0;
    

    const handleBack = () => {
        navigate(`/echo/${echoTag}`);
    };

    const handlePrevMessage = () => {
        console.log(`Переход на предыдущее сообщение.`);
        dispatch(soundEvent('SND_THEEND'));
    };
    const handleNextMessage = () => {
        console.log(`Переход на следующее сообщение.`);
        dispatch(soundEvent('SND_THEEND'));
    };

    useKeyboard({
        Escape: () => handleBack(),
        Delete: () => handleMsgRemove(),
        ArrowLeft: () => handlePrevMessage(),
        ArrowRight: () => handleNextMessage(),
    });
   
    return (
        <div className="Page Page-View">
            <MessageComponent />
        </div>
    );
};
