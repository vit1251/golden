
import { Link } from "react-router";

import "./Row.css";

export const Row = (props: { data: any, columns: any, onRowLink: any } ) => {

    const {
        data = [],
        columns = [],
        onRowLink = (row: string): string => '#',
    } = props;

    return (
        <>

            <div className="rowContainer">
               {data.map((row: any) => (
                 <Link to={onRowLink(row)}>
                 <div className="row">
                   {columns.map((column: { className: string, key: string, render: any }) => {
                       const { className = '', key, render } = column;
                       const { [key]: raw = '' } = row;
                       const value = render ? render(row) : raw;
                       return (
                           <div className={className}>{value}</div>
                       );
                   })}
                 </div>
                 </Link>
               ))}
            </div>

         </>
    );
};
