
let handlers: any = [];

const userHandler = ( event: KeyboardEvent ) => {
    //
    for (const handler of handlers) {
        handler(event);
    }
    //
    if (event.code == 'KeyZ' && (event.ctrlKey || event.metaKey)) {
        console.log('Нажат Ctrl+Z');
    }

}
    
document.addEventListener('keydown', (event) => userHandler(event));
//document.removeEventListener('keydown', userHandler);
   

export function useInput(handler: any): () => void {
    console.log(`--- Присоединили горячие клавиши ---`);
    handlers.push(handler);
    return () => {
        console.log(`--- Отсоединили горячие клавиши ---`);
        const newHandlers = [];
        for (const curHandler of handlers) {
            if (curHandler !== handler) {
                newHandlers.push(curHandler);
            }
        }
        handlers = newHandlers;
    }
}
