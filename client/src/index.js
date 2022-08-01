
import { Provider } from 'react-redux';
import { createRoot } from "react-dom/client";
import { HashRouter, Routes, Route } from "react-router-dom";
import { Welcome } from "./pages";
import { EchoIndex } from './pages/EchoIndex';
import { EchoAreaIndex } from './pages/EchoAreaIndex';
import { EchoMailView } from './pages/EchoMailView';
import { FileIndex } from './pages/FileIndex';
import './custom.css';
import './theme_black.css';
import { store } from './Storage.js';

const container = document.getElementById('app');
const root = createRoot(container);
root.render(
    <Provider store={store}>
        <HashRouter>
            <Routes>
                <Route path="/" element={<Welcome />} />
                <Route path="/echomail" element={<EchoIndex />} />
                <Route path="/echomail/:echoTag" element={<EchoAreaIndex />} />
                <Route path="/echomail/:echoTag/view" element={<EchoMailView />} />
                <Route path="/files" element={<FileIndex />} />
            </Routes>
        </HashRouter>
    </Provider>
);
