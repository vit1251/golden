
import { useParams, useNavigate } from "react-router";
import { useSelector } from 'react-redux';
import { useEffect } from 'react';

import { Header } from '../../common/Header';
import { eventBus } from '../../EventBus';
import { Rows } from './Row';

import "./EchoMsgIndex.css";

export const EchoMsgIndex = () => {

    const navigate = useNavigate();

    const areas = useSelector((state: any) => state.areas);
    const messages = useSelector((state: any) => state.messages);

    useEffect(() => {
        eventBus.invoke({
            type: 'ECHO_INDEX',
        });
    }, []);

    const { echoTag } = useParams();
    console.log(`echoTag = `, echoTag);

    const area = areas.find((area: any) => area.area_index === echoTag);
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
    const handleCreateMessage = () => {
        navigate(`/echomail/${echoTag}/create`);
    };

    return (
        <>

            <Header />

            <div className="container">
                <h1>Echoarea</h1>

                <Rows
                    onRowLink={(row: any) => {
                        const { hash = '' } = row;
                        return `/echomail/${echoTag}/${hash}/view`;
                    }}
                    columns={[
                       {className: 'rowUserpic', key: ''},
                       {className: 'rowFrom', key: 'from'},
                       {className: 'rowMarker', render: (row: any) => {
                           const { view_count = 0 } = row;
                           const value = view_count === 0 ? '•' : null;
                           return value;
                       }},
                       {className: 'rowSubject', key: 'subject'},
                       {className: 'rowDate', key: 'date'},
                    ]}
                    records={messages}
                    />

            </div>

        </>
    );
};
