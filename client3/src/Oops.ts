
import type { Screen } from "./Screen.ts";
import { store } from "./store/index.ts";
import { writeCenter } from "./util.ts";

export const Oops = (screen: Screen) => {

    const { scene } = store.getState().app;

    screen.setForegroudColor('#C00000');
    screen.writeText(4, 2, 'Ошибка прилоежния');
    screen.writeText(4, 4, `Нет реализации сцены с идентификатором ${scene}`);
    
};