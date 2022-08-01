
import { useParams } from "react-router-dom";
import { Header } from './Header';

export const EchoMailIndex = (props) => {
    const { echoTag, msgId } = useParams();
    console.log(echoTag);
    return (
        <>
            <Header />
        </>
    );
};