
import React from "react";
import { HashRouter, Routes, Route } from 'react-router';

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

export const App = () => {
    return (
        <HashRouter>
            <Routes>
                
                <Route index element={<Welcome />} />
                
                <Route path="netmail">
                    <Route index element={<NetmailIndex />} />
                    <Route path=":msgId">
                        <Route path="view" element={<NetmailView />} />
                    </Route>
                </Route>

                <Route path="echomail">
                    <Route index element={<EchoIndex />} />
                    <Route path=":echoTag">
                        <Route index element={<EchoMsgIndex />} />
                        <Route path="create" element={<EchoMsgCompose />} />
                        <Route path=":msgId">
                            <Route path="view" element={<EchoMsgView />} />
                        </Route>
                    </Route>
                </Route>
                
                <Route path="files">
                    <Route index element={<FileIndex />} />
                    <Route path=":echoTag">
                        <Route path="tics" element={<FileTicIndex />} />
                        <Route path=":fileId">
                            <Route path="view" element={<FileTicView />} />
                        </Route>
                    </Route>
                </Route>

                <Route path="setup" element={<Setup />} />

            </Routes>
        </HashRouter>
    );
};
