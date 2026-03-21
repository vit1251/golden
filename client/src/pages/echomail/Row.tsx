
import { ReactElement } from "react";
import { useNavigate } from "react-router";

import "./Row.css";

export interface Column<T> {
    className: string,
    key?: keyof T,
    styles?: (row: T) => {} | undefined,
    render?: (row: T) => number | string | ReactElement,
}

export const Row = <T extends object>({ activeIndex = 0, index, record, columns, onRowLink }: { activeIndex: number, index: number, record: T, columns: Column<T>[], onRowLink: any } ) => {

    const navigate = useNavigate();

    const handleDoubleClick = (row: T) => {
        console.log(`Открываем полноэкранный просмотр`);
        const linkAddr: string = onRowLink(row);
        navigate(linkAddr);
    };

    const handleClick = (row: T) => {
        console.log(`Открываем предварительный просмотр`);
    };

    const classes: string[] = ['row'];
    if (index === activeIndex) {
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

export const Rows = <T extends object>({ activeIndex = 0, records, columns, onRowLink }: { activeIndex?: number, records: T[], columns: Column<T>[], onRowLink: any } ) => {
    const classes: string[] = ['rowContainer'];
    return (
        <div className={classes.join(' ')}>
            {records.map((record: T, index: number) => (<Row activeIndex={activeIndex} index={index} record={record} columns={columns} onRowLink={onRowLink} />))}
        </div>
    );
};
