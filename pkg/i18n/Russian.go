package i18n

func init() {

	mainTranslation := GetMainTranslation()

	russianTranslation := mainTranslation.GetLangTranslation("ru-RU")
	/* Main menu */
	actionTranslation := russianTranslation.GetActionTranslation("Action")
	actionTranslation.SetTranslation("menu-home", "Главная")
	actionTranslation.SetTranslation("menu-netmail", "Личные сообщения")
	actionTranslation.SetTranslation("menu-echomail", "Телеконференции")
	actionTranslation.SetTranslation("menu-files", "Файлы")
	actionTranslation.SetTranslation("menu-service", "Обслуживание")
	actionTranslation.SetTranslation("menu-people", "Люди")
	actionTranslation.SetTranslation("menu-draft", "Черновики")
	actionTranslation.SetTranslation("menu-setup", "Настройки")

	/* NetmailIndexAction */
	netmailIndexActionTranslation := russianTranslation.GetActionTranslation("NetmailIndexAction")
	netmailIndexActionTranslation.SetTranslation("action-button-create", "Создать сообщение")
	/* EchoAreaIndexAction */
	echoAreaIndexActionTranslation := russianTranslation.GetActionTranslation("EchoAreaIndexAction")
	echoAreaIndexActionTranslation.SetTranslation("action-button-create", "Создать")
	/* EchoMsgIndexAction */
	echoMsgIndexActionTranslation := russianTranslation.GetActionTranslation("EchoMsgIndexAction")
	echoMsgIndexActionTranslation.SetTranslation("action-compose-button", "Создать сообщение")
	echoMsgIndexActionTranslation.SetTranslation("action-tree-button", "Дерево")
	echoMsgIndexActionTranslation.SetTranslation("action-mark-as-read-button", "Пометить как прочитанное")
	echoMsgIndexActionTranslation.SetTranslation("action-settings-button", "Настроить")
	/* FileEchoIndexAction */
	fileEchoIndexActionTranslation := russianTranslation.GetActionTranslation("FileEchoIndexAction")
	fileEchoIndexActionTranslation.SetTranslation("action-button-create", "Создать")

}
