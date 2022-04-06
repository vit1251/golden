package i18n

func init() {

	/* Initialize master mapper */
	mainTranslation = new(MainTranslation)
	mainTranslation.mainTranslations = make(map[string]*LangTranslation)

	/* Russian */
	{
		russianTranslation := mainTranslation.GetLangTranslation("ru-RU")
		/* Main menu */
		mainTranslation := russianTranslation.GetActionTranslation("Action")
		mainTranslation.SetTranslation("menu-home", "Главная")
		mainTranslation.SetTranslation("menu-netmail", "Личные сообщения")
		mainTranslation.SetTranslation("menu-echomail", "Телеконференции")
		mainTranslation.SetTranslation("menu-files", "Файлы")
		mainTranslation.SetTranslation("menu-service", "Обслуживание")
		mainTranslation.SetTranslation("menu-people", "Люди")
		mainTranslation.SetTranslation("menu-draft", "Черновики")
		mainTranslation.SetTranslation("menu-setup", "Настройки")

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

	/* English */
	{
		englishTranslation := mainTranslation.GetLangTranslation("en-US")

		/* Main menu */
		mainTranslation := englishTranslation.GetActionTranslation("Action")
		mainTranslation.SetTranslation("menu-home", "Home")
		mainTranslation.SetTranslation("menu-netmail", "Netmail")
		mainTranslation.SetTranslation("menu-echomail", "Echomail")
		mainTranslation.SetTranslation("menu-files", "Files")
		mainTranslation.SetTranslation("menu-service", "Service")
		mainTranslation.SetTranslation("menu-people", "People")
		mainTranslation.SetTranslation("menu-draft", "Draft")
		mainTranslation.SetTranslation("menu-setup", "Setup")

		/* NetmailIndexAction */
		netmailIndexActionTranslation := englishTranslation.GetActionTranslation("NetmailIndexAction")
		netmailIndexActionTranslation.SetTranslation("action-button-create", "Compose")
		/* EchoAreaIndexAction */
		echoAreaIndexActionTranslation := englishTranslation.GetActionTranslation("EchoAreaIndexAction")
		echoAreaIndexActionTranslation.SetTranslation("action-button-create", "Create")
		/* EchoMsgIndexAction */
		echoMsgIndexActionTranslation := englishTranslation.GetActionTranslation("EchoMsgIndexAction")
		echoMsgIndexActionTranslation.SetTranslation("action-compose-button", "Compose")
		echoMsgIndexActionTranslation.SetTranslation("action-tree-button", "Tree")
		echoMsgIndexActionTranslation.SetTranslation("action-mark-as-read-button", "Mark as read")
		echoMsgIndexActionTranslation.SetTranslation("action-settings-button", "Settings")
		/* FileEchoIndexAction */
		fileEchoIndexActionTranslation := englishTranslation.GetActionTranslation("FileEchoIndexAction")
		fileEchoIndexActionTranslation.SetTranslation("action-button-create", "Create")
	}

	/* Dutch */
	{
		englishTranslation := mainTranslation.GetLangTranslation("nl-BE")
		/* NetmailIndexAction */
		netmailIndexActionTranslation := englishTranslation.GetActionTranslation("NetmailIndexAction")
		netmailIndexActionTranslation.SetTranslation("action-button-create", "Nieuw bericht")
		/* EchoAreaIndexAction */
		echoAreaIndexActionTranslation := englishTranslation.GetActionTranslation("EchoAreaIndexAction")
		echoAreaIndexActionTranslation.SetTranslation("action-button-create", "Maak een nieuw gebied")
		/* EchoMsgIndexAction */
		echoMsgIndexActionTranslation := englishTranslation.GetActionTranslation("EchoMsgIndexAction")
		echoMsgIndexActionTranslation.SetTranslation("action-compose-button", "Nieuw bericht")
		echoMsgIndexActionTranslation.SetTranslation("action-tree-button", "Boomstructuur")
		echoMsgIndexActionTranslation.SetTranslation("action-mark-as-read-button", "Markeer als gelezen")
		echoMsgIndexActionTranslation.SetTranslation("action-settings-button", "Instellingen")
	}	
}
