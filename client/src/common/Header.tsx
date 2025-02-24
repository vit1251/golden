

import { Link } from "react-router";
import { useSelector } from 'react-redux';

import { useTranslation } from 'react-i18next';

import './Header.css';

export const Header = () => {

    const { t, i18n } = useTranslation();

    const {
        NetMessageCount = 0,
        EchoMessageCount = 0,
        FileCount = 0,
    } = useSelector((state: any) => state.summary);

    const items = [
        {
            name: t('Home'),
            path: '/',
        },
        {
            name: t('Netmail'),
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
