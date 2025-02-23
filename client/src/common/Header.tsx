
import React from 'react';
import { Link } from "react-router";
import { useSelector } from 'react-redux';
import { eventBus } from '../EventBus';

import './Header.css';

export const Header = () => {

    const {
        NetMessageCount = 0,
        EchoMessageCount = 0,
        FileCount = 0,
    } = useSelector((state: any) => state.summary);

    const items = [
        {
            name: 'Главная',
            path: '/',
        },
        {
            name: 'Личные сообщения',
            path: '/netmail',
            itemCount: NetMessageCount,
        },
        {
            name: 'Телеконференции',
            path: '/echo',
            itemCount: EchoMessageCount,
        },
        {
            name: 'Файлы',
            path: '/files',
            itemCount: FileCount,
        },
        {
            name: 'Обслуживание',
            path: '/service',
        },
        {
            name: 'Люди',
        },
        {
            name: 'Черновики',
        }
    ];

    return (
        <div className="Header">
            <div className="HeaderGroup">
                {items.map((item, index) => (
                <div key={index} className="HeaderItem">
                    <span className="HeaderLabel">
                        <Link to={item.path ?? "#"}>{item.name}</Link>
                    </span>
                    {item.itemCount ? (
                        <>
                            {' '}
                            <span className="badge">{item.itemCount}</span>
                        </>
                    ) : null}
                </div>
                ))}
            </div>
            <div className="HeaderGroup">
                <div className="HeaderItem HeaderIcon">
                    <Link to="/setup">Настройки</Link>
                </div>
            </div>
        </div>
    );
};
