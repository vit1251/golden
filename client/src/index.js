
import { Provider } from 'react-redux';
import { createRoot } from "react-dom/client";

import { App } from "./pages";

import { store } from './Storage.js';

const container = document.getElementById('app');
const root = createRoot(container);
root.render(
    <Provider store={store}>
        <App />
    </Provider>
);
