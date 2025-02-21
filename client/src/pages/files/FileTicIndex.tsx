
import { useParams } from "react-router";
import { Header } from '../../common/Header';

export const FileTicIndex = () => {
    const { echoTag, msgId } = useParams();
    console.log(echoTag);
    return (
        <>
            <Header />
        </>
    );
};