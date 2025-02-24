
import { useNavigate } from "react-router";

import "./Row.css";
import { Area } from "./EchoMsgIndex";
import { ReactElement, useEffect, useState } from "react";
import { stringToHexColor } from "../../usils";
import { useInput } from "../../Hotkey";

export interface Column<T> {
    className: string,
    key?: keyof T,
    styles?: (row: T) => {} | undefined,
    render?: (row: T) => number | string | ReactElement,
}

export const Row = <T extends object>({ index, record, columns, onRowLink }: { index:number, record: T, columns: Column<T>[], onRowLink: any } ) => {

    const [state, setState] = useState({
        activeIndex: 0,
    });
    const navigate = useNavigate();

    useEffect(() => {
        const removeHotkeys = useInput((event: KeyboardEvent) => {
            if (event.key === 'ArrowUp') {
                setState((prev) => ({
                    ...prev,
                    activeIndex: prev.activeIndex > 0 ? prev.activeIndex - 1 : prev.activeIndex,
                }))
            }
            if (event.key === 'ArrowDown') {
                setState((prev) => ({
                    ...prev,
                    activeIndex: prev.activeIndex + 1,
                }))
            }
        });
        return () => removeHotkeys();
    }, [])

    const handleDoubleClick = (row: T) => {
        console.log(`Открываем полноэкранный просмотр`);
        const linkAddr: string = onRowLink(row);
        navigate(linkAddr);
    };

    const handleClick = (row: T) => {
        console.log(`Открываем предварительный просмотр`);
    };

    const classes: string[] = ['row'];
    if (index === state.activeIndex) {
        classes.push('rowActive');
    }

    return (
        <div className={classes.join(' ')} onClick={() => handleClick(record)} onDoubleClick={() => handleDoubleClick(record)}>
            {columns.map((column: Column<T>) => {
                const { className = '', key, styles, render } = column;
                const { [key]: raw = '' } = record;
                const value = render ? render(record) : `${raw}`;
                const userStyle = styles ? styles(record) : {};
                return (
                    <div style={userStyle} className={className}>{value}</div>
                );
            })}
        </div>
    );
};

export const Rows = <T extends object>({ records, columns, onRowLink }: { records: T[], columns: Column<T>[], onRowLink: any } ) => {
    const classes: string[] = ['rowContainer'];
    return (
        <div className={classes.join(' ')}>
            {records.map((record: T, index: number) => (<Row index={index} record={record} columns={columns} onRowLink={onRowLink} />))}
        </div>
    );
};
