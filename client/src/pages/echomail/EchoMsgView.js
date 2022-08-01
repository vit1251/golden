
import Hotkeys from 'react-hot-keys';

import { useParams } from "react-router-dom";
import { Header } from '../../common/Header';

export const EchoMsgView = (props) => {
    const { echoTag, msgId } = useParams();
    console.log(echoTag);
    const handlePrevMessage = () => {
        console.log(`handlePrevMessage...`);
    };
    return (
        <>
            <Header />

            <Hotkeys 
                keyName="ctrl+left,pgup"
                onKeyUp={handlePrevMessage}
                />
            
            <h1>EchoMailView</h1>
            
            
            
        </>
    );
};
