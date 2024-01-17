
import React from 'react';
import { Link } from "react-router-dom";

import "./Row.css";

export const Row = (props) => {

    const {
        data = [],
        columns = [],
        onRowLink = (row) => '#',
    } = props;

    return (
        <>

            <div className="rowContainer">
               {data.map((row) => (
                 <Link to={onRowLink(row)}>
                 <div className="row">
                   {columns.map((column) => {
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
