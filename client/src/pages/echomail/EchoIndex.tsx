
import { useNavigate } from 'react-router';

import { useDispatch, useSelector } from 'react-redux';
import { useCallback, useEffect } from 'react';

import { Rows } from './Row.tsx';

import "./EchoIndex.css";

import { type Area } from '../../models/Area.model.ts';
import { useKeyboard } from '../../Hotkey.tsx';
import { type RootState } from '../../app/store.ts';
import { echoIndex, nextArea, prevArea } from '../../features/areaSlice.ts';
import { soundEvent } from '../../middleware/soundMiddleware.ts';
import { socketSend } from '../../middleware/socketMiddleware.ts';

export const EchoIndex = () => {
    
    const navigate = useNavigate();

    const dispatch = useDispatch();

    const currentPage = useSelector((state: RootState) => state.areas.currentPage);
    const pageSize = useSelector((state: RootState) => state.areas.pageSize);

    const areas: Area[] = useSelector((state: RootState) => state.areas.records) ?? [];
    const activeIndex: number = useSelector((state: RootState) => state.areas.activeIndex) ?? 0;
    console.log(areas);
    console.log(activeIndex);

    useEffect(() => {
        // 1. Запросим список эхоконференций
        dispatch(socketSend({
            msg: {
                type: 'ECHO_INDEX',
            },
        }));
    }, []);

    const handlePrevArea = () => {
        console.log(`handlePrevMessage...`);
        // Шаг 1. Проверяем в начале мы списка
        if (activeIndex === 0) {
            dispatch(soundEvent('SND_THEEND'));
        }
        // Шаг 2. Образабтываем запрос
        dispatch(prevArea());
    };
    const handleNextArea = () => {
        console.log(`handleNextMessage...`);
        // Шаг 1. Проверяем в конце ли мы списка
        if (activeIndex + 1 === areas.length) {
            dispatch(soundEvent('SND_THEEND'));
        }
        // Шаг 2. Обрабатываем запрос
        dispatch(nextArea());
    };
    const handleOpenArea = () => {
        const area: Area | null = areas.at(activeIndex) ?? null;
        console.log(`openArea: areaIndex = ${activeIndex}`);
        console.log(area);
        if (area) {
            navigate(`/echo/${area.area_index}`);
        }
    };

    useKeyboard({
        ArrowUp: () => handlePrevArea(),
        ArrowDown: () => handleNextArea(),
        Enter: () => handleOpenArea(),
    });

    const totalPages = Math.ceil(areas.length / pageSize);
    const startIndex = (currentPage - 1) * pageSize;
    const paginatedAreas = areas.slice(startIndex, startIndex + pageSize);

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
                records={paginatedAreas} />

        </div>
    );
};