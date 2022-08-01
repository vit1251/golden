
import { useParams } from "react-router-dom";
import { Header } from '../../common/Header';

export const NetmailView = (props) => {
    const { echoTag, msgId } = useParams();
    console.log(echoTag);
    return (
        <>
            <Header />
        </>
    );
};