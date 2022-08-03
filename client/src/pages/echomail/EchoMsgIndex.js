
import { useNavigate } from "react-router-dom";
import { useSelector, useDispatch } from 'react-redux';
import { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';

import Hotkeys from 'react-hot-keys';

import { Header } from '../../common/Header';
import { eventBus } from '../../EventBus.js';
import { Row } from './Row.js';

import "./EchoMsgIndex.css";

export const EchoMsgIndex = (props) => {

    const navigate = useNavigate();

    const [activeIndex, setActiveIndex] = useState(0);

    const areas = useSelector((state) => state.areas);
    const messages = useSelector((state) => state.messages);

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
    const handleNextMessage = () => {
        console.log(`handlePrevMessage...`);
    };
    const handleAreaIndex = () => {
        navigate(`/echomail`);
    };

    return (
        <>

            <Header />

            <div class="container">
                <h1>Echoarea</h1>

                <Row
                    onRowLink={(row) => {
                        const { hash = '' } = row;
                        return `/echomail/${echoTag}/${hash}/view`;
                    }}
                    columns={[
                       {className: 'rowUserpic', key: ''},
                       {className: 'rowFrom', key: 'from'},
                       {className: 'rowMarker', render: (row) => {
                           const { view_count = 0 } = row;
                           const value = view_count == 0 ? '•' : null;
                           return value;
                       }},
                       {className: 'rowSubject', key: 'subject'},
                       {className: 'rowDate', key: 'date'},
                    ]}
                    data={messages}
                    activeIndex={activeIndex}
                    />

            </div>

            <Hotkeys keyName="up" onKeyDown={handlePrevMessage} />
            <Hotkeys keyName="down" onKeyDown={handleNextMessage} />
            <Hotkeys keyName="esc" onKeyDown={handleAreaIndex} />

        </>
    );
};
