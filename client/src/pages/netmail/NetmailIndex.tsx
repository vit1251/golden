
import { useParams } from "react-router";

export const NetmailIndex = () => {
    const { echoTag, msgId } = useParams();
    console.log(echoTag);
    return (
        <>
            <h1>Not yet implemented.</h1>
        </>
    );
};
