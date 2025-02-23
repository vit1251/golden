
import { useNavigate } from "react-router";

import "./Row.css";

export const Row = ({ record, columns, onRowLink }: { record: any, columns: any, onRowLink: any } ) => {

    const navigate = useNavigate();

    const handleDoubleClick = (row: any) => {
        console.log(`Открываем полноэкранный просмотр`);
        const linkAddr: string = onRowLink(row);
        navigate(linkAddr);
    };

    const handleClick = (row: any) => {
        console.log(`Открываем предварительный просмотр`);
    };

    return (
        <div className="row" onClick={() => handleClick(record)} onDoubleClick={() => handleDoubleClick(record)}>
            {columns.map((column: { className: string, key: string, render: any }) => {
                const { className = '', key, render } = column;
                const { [key]: raw = '' } = record;
                const value = render ? render(record) : raw;
                return (
                    <div className={className}>{value}</div>
                );
            })}
        </div>
    );
};

export const Rows = ({ records, columns, onRowLink }: { records: any[], columns: any[], onRowLink: any } ) => {
    return (
        <>
            <div className="rowContainer">
               {records.map((record: any) => (<Row record={record} columns={columns} onRowLink={onRowLink} />))}
            </div>
        </>
    );
};
