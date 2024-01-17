
import { HashRouter, Routes, Route } from "react-router-dom";

import { Welcome } from "./Welcome";

import { EchoIndex } from './echomail/EchoIndex';
import { EchoMsgIndex } from './echomail/EchoMsgIndex';
import { EchoMsgView } from './echomail/EchoMsgView';
import { EchoMsgCompose } from './echomail/EchoMsgCompose';
import { FileIndex } from './files/FileIndex';
import { FileTicIndex } from './files/FileTicIndex';
import { FileTicView } from './files/FileTicView';
import { NetmailIndex } from './netmail/NetmailIndex';
import { NetmailView } from './netmail/NetmailView';
import { Setup } from './setup/Setup';

import '../themes/custom.css';
import '../themes/theme_black.css';

export const App = (props) => {
    return (
        <>
            <HashRouter>
                <Routes>
                    <Route path="/" element={<Welcome />} />
                    <Route path="/netmail" element={<NetmailIndex />} />
                    <Route path="/netmail/:msgId/view" element={<NetmailView />} />
                    <Route path="/echomail" element={<EchoIndex />} />
                    <Route path="/echomail/:echoTag" element={<EchoMsgIndex />} />
                    <Route path="/echomail/:echoTag/:msgId/view" element={<EchoMsgView />} />
                    <Route path="/echomail/:echoTag/create" element={<EchoMsgCompose />} />
                    <Route path="/files" element={<FileIndex />} />
                    <Route path="/files/:echoTag/tics" element={<FileTicIndex />} />
                    <Route path="/files/:echoTag/:fileId/view" element={<FileTicView />} />
                    <Route path="/setup" element={<Setup />} />
                </Routes>
            </HashRouter>
        </>
    );
};
