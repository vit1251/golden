
import { Color } from "./color.ts";
import type { Screen } from "./Screen.ts";
import { store } from "./store/index.ts";


export const Oops = (screen: Screen) => {

    const { scene } = store.getState().app;

    screen.setForegroudColor(Color.Yellow);
    screen.writeText(4, 2, 'Ошибка приложeния');
    screen.writeText(4, 3, '=================');

    screen.setForegroudColor(Color.Gray);
    screen.writeText(4, 5, `Произошла внутренняя неустранимая ошибка.`);
    screen.writeText(4, 7, `Перезагрузите приложение.`);
    
};