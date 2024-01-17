
import { Link } from "react-router-dom";
import { useSelector } from 'react-redux';
import { eventBus } from '../EventBus';

import { Settings } from 'react-feather';

import './Header.css';

export const Header = (props) => {

    const {
        NetMessageCount = 0,
        EchoMessageCount = 0,
        FileCount = 0,
    } = useSelector((state) => state.summary);

    const items = [
        {
            name: 'Home',
            path: '/',
        },
//        {
//            name: 'Netmail',
//            path: '/netmail',
//            itemCount: NetMessageCount,
//        },
        {
            name: 'Echomail',
            path: '/echomail',
            itemCount: EchoMessageCount,
        },
//        {
//            name: 'Files',
//            path: '/files',
//            itemCount: FileCount,
//        },
//        {
//            name: 'Service',
//            path: '/service',
//        },
//        {
//            name: 'People',
//        },
//        {
//            name: 'Draft',
//        }
    ];

    return (
        <div className="Header">
            <div className="HeaderGroup">
                {items.map((item, index) => (
                <div key={index} className="HeaderItem">
                    <span className="HeaderLabel">
                        <Link to={item.path}>{item.name}</Link>
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
                    <Link to="/setup"><Settings size={20} /></Link>
                </div>
            </div>
        </div>
    );
};
