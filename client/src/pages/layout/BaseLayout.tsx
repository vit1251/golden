
import { type ReactElement, useEffect, useRef, useState } from "react";
import { Header } from "../../common/Header.tsx";

import "./BaseLayout.css";


export const BaseLayout = ({ content }: { content: ReactElement }) => {

    // Шаг 1. Регистрация часов
    const [ clock, setClock ] = useState('00:00:00');
    useEffect(() => {
        const timerId: number = setInterval(() => {
            const now: Date = new Date();
            const hour: string = String(now.getHours()).padStart(2, '0');
            const min: string = String(now.getMinutes()).padStart(2, '0');
            const sec: string = String(now.getSeconds()).padStart(2, '0');
            const clock: string = `${hour}:${min}:${sec}`;
            setClock(clock);
        }, 500);
        return () => {
            clearInterval(timerId);
        };
    }, []);

    // Шаг 2. Автофокус основного блока страницы
    const scrollRef = useRef<HTMLElement>(null);
    useEffect(() => {
        if (scrollRef.current) {
            scrollRef.current.focus();
        }
    });

    return (
        <div className="Container">
            <Header />
            <main className="Main" ref={scrollRef}>{content}</main>
            <footer className="Status">
                <div className="StatusBar">
                    <div className="StatusBar-Start">Golden Point v1.2.19 - версия разработчика</div>
                    <div className="StatusBar-End">{clock}</div>
                </div>
            </footer>
        </div>
    );
}
