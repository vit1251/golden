
import { useParams } from "react-router";
import { Header } from '../../common/Header';

export const FileTicView = () => {
    const { echoTag, msgId } = useParams();
    console.log(echoTag);
    return (
        <>
            <Header />
        </>
    );
};