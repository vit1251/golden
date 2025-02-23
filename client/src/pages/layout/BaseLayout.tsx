
import { ReactElement } from "react";
import { Header } from "../../common/Header";

export const BaseLayout = ({ content }: { content: ReactElement }) => {

    return (
        <>
            <Header />
            {content}
        </>
    );

}
