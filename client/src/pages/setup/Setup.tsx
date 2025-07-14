
import { Header } from '../../common/Header';

import i18n from '../../i18n';

export const Setup = () => {

    const handleRussian = () => {
        i18n.changeLanguage('ru-RU');
    }

    const handleEnglish = () => {
        i18n.changeLanguage('en-US');
    }

    return (
        <div>

            <h1>Настройка</h1>

            <h2>Параметры подключения</h2>

            <h2>Выбор языка</h2>

            <div onClick={handleRussian}>Русский</div>
            <div onClick={handleEnglish}>English</div>

        </div>
    );

};
