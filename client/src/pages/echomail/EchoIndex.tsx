
import { useDispatch, useSelector } from 'react-redux';
import { useCallback, useEffect } from 'react';

import { Rows } from './Row';

import "./EchoIndex.css";
import { Area } from '../../models/Area.model';
import { useInput } from '../../Hotkey';
import { RootState } from '../../app/store';
import { nextArea, prevArea } from '../../features/areaSlice';
import { useNavigate } from 'react-router';

export const EchoIndex = () => {
    
    const navigate = useNavigate();

    const dispatch = useDispatch();

    const sendMessage = (payload: any) => {
        dispatch({
            type: 'SOCKET_SEND',
            payload: payload,
        });
    };

    const areas: Area[] = useSelector((state: RootState) => state.areas.records) ?? [];
    const activeIndex: number = useSelector((state: RootState) => state.areas.activeIndex) ?? 0;
    console.log(areas);
    console.log(activeIndex);

    useEffect(() => {
        sendMessage({
            type: 'ECHO_INDEX',
        });
    }, []);

    const handlePrevArea = () => {
        console.log(`handlePrevMessage...`);
        dispatch(prevArea());
    };
    const handleNextArea = () => {
        console.log(`handleNextMessage...`);
        dispatch(nextArea());
    };
    const handleOpenArea = useCallback(() => {
        const area: Area | null = areas.at(activeIndex) ?? null;
        console.log(`openArea: areaIndex = ${activeIndex}`);
        console.log(area);
        if (area) {
            navigate(`/echo/${area.area_index}`);
        }
    }, [ areas, activeIndex ]);

    useEffect(() => {
        const removeHotkeys = useInput((event: KeyboardEvent) => {
            if (event.key === 'ArrowUp') handlePrevArea();
            if (event.key === 'ArrowDown') handleNextArea();
            if (event.key === 'Enter') handleOpenArea();
        });
        return () => removeHotkeys();
    }, [ handlePrevArea, handleNextArea, handleOpenArea ]);

    return (
        <div className="Page Page-Areas">

            <Rows<Area>
                activeIndex={activeIndex}
                onRowLink={(row: Area) => `/echo/${row.area_index}`}
                columns={[
                    {
                        className: "rowMarker", render: (row: any): string => {
                            const { new_message_count = 0 } = row;
                            //const value = new_message_count > 0 ? '•' : '';
                            const value = new_message_count > 0 ? '+' : '';
                            return value;
                        }
                    },
                    { className: "rowSummary", key: "summary" },
                    {
                        className: "rowMessageCount", render: (row: Area): string => {
                            const { message_count = 0 } = row;
                            return `${message_count}`;
                        }
                    },
                    {
                        className: "rowMessageNewCount", render: (row: Area): string => {
                            const { new_message_count = 0 } = row;
                            return `${new_message_count}`;
                        }
                    },
                    { className: "rowName", key: "name" },
                ]}
                records={areas} />

        </div>
    );
};