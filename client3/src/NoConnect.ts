
import { Color } from "./color.ts";
import type { Screen } from "./Screen.ts";

export const NoConnect = (screen: Screen) => {

    screen.setForegroudColor(Color.Yellow);
    screen.writeText(4, 2, 'Ошибка связи');
    screen.writeText(4, 3, '============');

    screen.setForegroudColor(Color.Gray);
    screen.writeText(4, 5, `Произошла внутренняя неустранимая ошибка соединения с приложением.`);
    screen.writeText(4, 7, `Перезагрузите сервер приложения.`);

}