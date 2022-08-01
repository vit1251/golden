
import { Link } from "react-router-dom";
import { useSelector } from 'react-redux';
import { eventBus } from '../EventBus';

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
        {
            name: 'Netmail',
            path: '/netmail',
            itemCount: NetMessageCount,
        },
        {
            name: 'Echomail',
            path: '/echomail',
            itemCount: EchoMessageCount,
        },
        {
            name: 'Files',
            path: '/files',
            itemCount: FileCount,
        },
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
            <div className="Header-item-group">
                {items.map((item, index) => (
                <div key={index} className="Header-item">
                    <span className="tab-label">
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
{/*
            <div className="Header-item-group">
                <a className="nav-link" href="/setup">
                    <div className="Header-item">
                        <span className="tab-label">
                            <Link to="/setup">Setup</Link>
                        </span>
                        <span className="badge hidden" id="mainMenuSetup"></span>
                    </div>
                </a>
            </div>
*/}
        </div>
    );
};
