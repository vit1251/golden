
import { useParams } from "react-router";

export const FileIndex = () => {

    const { echoTag, msgId } = useParams();

    console.log(echoTag);
    return (
        <div>
            Показать список файлов
        </div>
    );
};
