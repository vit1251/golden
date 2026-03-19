

import { Link } from "react-router";
import { useSelector } from 'react-redux';

import { useTranslation } from 'react-i18next';

import './Header.css';
import { IconTune } from "../IconTune";

interface Item {
    name: string,
    path?: string,
}

export const Header = () => {

    const { t, i18n } = useTranslation();

    const items: Array<Item> = [
        {
            name: t('Home'),
            path: '/',
        },
        {
            name: t('Netmail'),
            path: '/netmail',
        },
        {
            name: t('Echomail'),
            path: '/echo',
        },
        {
            name: t('Files'),
            path: '/files',
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
        <header className="Header">
            <div className="HeaderGroup">
                {items.map((item, index) => (
                <div key={index} className="HeaderItem">
                    <span className="HeaderLabel">
                        <Link to={item.path ?? "#"}>{item.name}</Link>
                    </span>
                </div>
                ))}
            </div>
            <div className="HeaderGroup">
                <div className="HeaderItem HeaderIcon">
                    <Link to="/setup"><IconTune /></Link>
                </div>
            </div>
        </header>
    );
};
