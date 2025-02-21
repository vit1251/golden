
import React from 'react';
import { Provider } from 'react-redux';
import { createRoot } from "react-dom/client";

import { App } from "./pages/App";

import { store } from './Storage';

const rootElement: string = 'root';
const container: Element | null = document.getElementById(rootElement);
if (container) {
    const root = createRoot(container);
    root.render(
        <Provider store={store}>
            <App />
        </Provider>
    );
} else {
    console.log(`Не удалось найти корневой элемент "${rootElement}".`);
}
