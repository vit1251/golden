
import { ReactElement } from "react";
import { Header } from "../../common/Header";

import "./BaseLayout.css";


export const BaseLayout = ({ content }: { content: ReactElement }) => {
    return (
        <div className="Container">
            <Header />
            <div className="Main">{content}</div>
        </div>
    );
}
