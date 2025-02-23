
import { useNavigate } from "react-router";

import "./Row.css";
import { Area } from "./EchoMsgIndex";
import { ReactElement } from "react";
import { stringToHexColor } from "../../usils";

export interface Column<T> {
    className: string,
    key?: keyof T,
    styles?: (row: T) => {} | undefined,
    render?: (row: T) => number | string | ReactElement,
}

export const Row = <T extends object>({ record, columns, onRowLink }: { record: T, columns: Column<T>[], onRowLink: any } ) => {

    const navigate = useNavigate();

    const handleDoubleClick = (row: T) => {
        console.log(`Открываем полноэкранный просмотр`);
        const linkAddr: string = onRowLink(row);
        navigate(linkAddr);
    };

    const handleClick = (row: T) => {
        console.log(`Открываем предварительный просмотр`);
    };

    return (
        <div className="row" onClick={() => handleClick(record)} onDoubleClick={() => handleDoubleClick(record)}>
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
    return (
        <>
            <div className="rowContainer">
               {records.map((record: T) => (<Row record={record} columns={columns} onRowLink={onRowLink} />))}
            </div>
        </>
    );
};
