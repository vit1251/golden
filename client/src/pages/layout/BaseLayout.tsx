
import { ReactElement } from "react";
import { Header } from "../../common/Header";

import "./BaseLayout.css";


export const BaseLayout = ({ content }: { content: ReactElement }) => {
    return (
        <div className="Container">
            <Header />
            <main className="Main">{content}</main>
            <footer className="Status">Golden Point v1.2.19 - версия разработчика</footer>
        </div>
    );
}
