
import React, { useEffect } from "react";
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
import { BaseLayout } from "./layout/BaseLayout";
import { eventBus } from "../EventBus";

const updateSummary = () => {
    
    /* Step 1. Invoke update summery */
    eventBus.invoke({
        type: 'SUMMARY',
    });

    /* Step 2. Rewind counter */
    setTimeout(updateSummary, 10_000);

};

export const App = () => {

    useEffect(() => {
        updateSummary();
    }, [])

    return (
        <HashRouter>
            <Routes>

                <Route index element={<BaseLayout content={<Welcome />} />} />

                <Route path="netmail">
                    <Route index element={<BaseLayout content={<NetmailIndex />} />} />
                    <Route path=":msgId">
                        <Route path="view" element={<BaseLayout content={<NetmailView />} />} />
                    </Route>
                </Route>

                <Route path="echo">
                    <Route index element={<BaseLayout content={<EchoIndex />} />} />
                    <Route path=":echoTag">
                        <Route index element={<BaseLayout content={<EchoMsgIndex />} />} />
                        <Route path="create" element={<EchoMsgCompose />} />
                        <Route path=":msgId">
                            <Route path="view" element={<BaseLayout content={<EchoMsgView />} />} />
                        </Route>
                    </Route>
                </Route>

                <Route path="files">
                    <Route index element={<BaseLayout content={<FileIndex />} />} />
                    <Route path=":echoTag">
                        <Route path="tics" element={<FileTicIndex />} />
                        <Route path=":fileId">
                            <Route path="view" element={<FileTicView />} />
                        </Route>
                    </Route>
                </Route>

                <Route path="setup" element={<BaseLayout content={<Setup />} />} />

            </Routes>
        </HashRouter>
    );
};
