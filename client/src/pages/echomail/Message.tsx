
import { useDispatch, useSelector } from 'react-redux';
import { useNavigate, useParams } from 'react-router';

import { View } from '../../common/msg-view/View.tsx';
import { View2 } from '../../common/msg-view/View2.tsx';


export const Message = () => {

    const dispatch = useDispatch();

    const navigate = useNavigate();

    const msgArea: string = useSelector((state: any) => state.view.echo);
    const msgFrom: string = useSelector((state: any) => state.view.from);
    const msgTo: string = useSelector((state: any) => state.view.to);
    const msgSubject: string = useSelector((state: any) => state.view.subject);
    const msgDate: string = useSelector((state: any) => state.view.date);
    const content: string = useSelector((state: any) => state.view.content);

    const { echoTag, msgId } = useParams();
    console.log(`echoTag = ${echoTag} msgId = ${msgId}`);

    return (
        <div className="Page-View">
            <div className="View-Header echo-msg-view-header-wrapper">
                <table className="echo-msg-view-header">
                    <tbody><tr className="" title="">
                        <td className="echo-msg-view-header-name">
                            <span className="">Area:</span>
                        </td>
                        <td className="echo-msg-view-header-value">
                            <span className="">{msgArea}</span>
                        </td>
                    </tr>
                        <tr className="" title="">
                            <td className="echo-msg-view-header-name">
                                <span className="">From:</span></td>
                            <td className="echo-msg-view-header-value">
                                <span className="">{msgFrom}</span></td>
                        </tr>
                        <tr className="" title="">
                            <td className="echo-msg-view-header-name">
                                <span className="">To:</span></td>
                            <td className="echo-msg-view-header-value">
                                <span className="">{msgTo}</span></td>
                        </tr>
                        <tr className="" title="">
                            <td className="echo-msg-view-header-name">
                                <span className="">Subject:</span></td>
                            <td className="echo-msg-view-header-value">
                                <span className="">{msgSubject ?? '-'}</span></td>
                        </tr>
                        <tr className="" title="">
                            <td className="echo-msg-view-header-name">
                                <span className="">Date:</span></td>
                            <td className="echo-msg-view-header-value">
                                <span className="">{msgDate}</span></td>
                        </tr>
                    </tbody></table>
            </div>
            <div className="View-Body echo-msg-view-body">
                <View rawText={content} />
            </div>
        </div>
    );
}