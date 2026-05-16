
import { type ReactElement, useEffect, useRef } from "react";
import { useNavigate } from "react-router";

import "./Row.css";

export interface Column<T> {
    className: string,
    key?: keyof T,
    styles?: (row: T) => {} | undefined,
    render?: (row: T) => number | string | ReactElement,
}

export const Row = <T extends Record<string, any>>({ ref, tabIndex, activeIndex = 0, index, record, columns, onRowLink }: { ref?: React.Ref<HTMLDivElement>, tabIndex: number, activeIndex: number, index: number, record: T, columns: Column<T>[], onRowLink: any } ) => {

    const navigate = useNavigate();

    const classes: string[] = ['row'];
    if (index === activeIndex) {
        classes.push('rowActive');
    }

    return (
        <div ref={ref} tabIndex={tabIndex} className={classes.join(' ')}>
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
    const itemRefs = useRef<(HTMLDivElement | null)[]>([]);

    useEffect(() => {
        const elementToFocus = itemRefs.current[activeIndex];
        if (elementToFocus) {
            elementToFocus.scrollIntoView({
                behavior: 'smooth',
                block: 'nearest', // Прокрутит только если элемента нет на экране
            });
        }
    }, [activeIndex]);

    return (
        <div className={classes.join(' ')}>
            {records.map((record: T, index: number) => (
                <Row
                    key={index}
                    tabIndex={-1}
                    ref={(el) => { itemRefs.current[index] = el; }}
                    activeIndex={activeIndex}
                    index={index}
                    record={record}
                    columns={columns}
                    onRowLink={onRowLink}
                    />))
            }
        </div>
    );
};
