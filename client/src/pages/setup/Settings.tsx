
import { useDispatch, useSelector } from 'react-redux';

import i18n from '../../i18n';
import { setCode } from '../../features/settingsSlice';

interface Language {
    name: string,
    code: string,
}

export const Settings = () => {
    const dispatch = useDispatch()

    const code: string = useSelector((state: any) => state.settings.code);

    const handleChange = (event: any) => {
        const value: string = event.target.value;
        dispatch(setCode(value));
        i18n.changeLanguage(value);
    }

    const languages: Array<Language> = [
        {name: 'Русский', code: 'ru-RU' },
        {name: 'English', code: 'en-US' },
    ];

    return (
        <div className="Page Page-Settings">

            <h1>Настройка</h1>

            <h2>Параметры подключения</h2>

            <h2>Выбор языка</h2>

            <select onChange={handleChange} value={code}>
                {languages.map(l => <option value={l.code}>{l.name}</option>)}
            </select>

        </div>
    );

};
