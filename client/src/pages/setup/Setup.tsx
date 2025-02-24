
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
        <>
            <Header />


            Язык
            
            <div onClick={handleRussian}>Русский</div>
            <div onClick={handleEnglish}>English</div>
        </>
    );

};
