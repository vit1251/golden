
import { ReactElement, useEffect, useState } from "react";
import { Header } from "../../common/Header";

import "./BaseLayout.css";


export const BaseLayout = ({ content }: { content: ReactElement }) => {

    const [ clock, setClock ] = useState('00:00:00');
    useEffect(() => {
        setInterval(() => {
            const now: Date = new Date();
            const hour: string = String(now.getHours()).padStart(2, '0');
            const min: string = String(now.getMinutes()).padStart(2, '0');
            const sec: string = String(now.getSeconds()).padStart(2, '0');
            const clock: string = `${hour}:${min}:${sec}`;
            setClock(clock);
        }, 500);
    }, []);

    return (
        <div className="Container">
            <Header />
            <main className="Main">{content}</main>
            <footer className="Status">
                <div className="StatusBar">
                    <div className="StatusBar-Start">Golden Point v1.2.19 - версия разработчика</div>
                    <div className="StatusBar-End">{clock}</div>
                </div>
            </footer>
        </div>
    );
}
