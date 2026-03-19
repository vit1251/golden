
export interface Area {
    readonly name: string,                   // Название конференции. Например: RU.ANEKDOT
    readonly summary: string,                // Описание конфиренции
    readonly message_count: number,          // Количество сообщений
    readonly new_message_count: number,      // Количество непрочитанных сообщений
    readonly order: number,                  // Порядковый номер
    readonly area_index: string,             // Индекс
}
