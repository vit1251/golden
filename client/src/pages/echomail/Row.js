
import React from 'react';
import { Link } from "react-router-dom";

import "./Row.css";

export const Row = (props) => {

    const { areas = [] } = props;

    return (
        <>

            <div className="rowContainer">
               {areas.map((area) => (
                 <Link to={`/echomail/${area.area_index}`}>
                 <div className="row">
                   <div className="rowName">{area.name}</div>
                   <div className="rowMarker">{true ? '•' : null}</div>
                   <div className="rowSummary">{area.summary}</div>
                   <div className="rowCounter">{area.new_message_count > 0 ? area.new_message_count : null}</div>
                 </div>
                 </Link>
               ))}
            </div>

         </>
    );
};
