package i18n

func init() {

	mainTranslation := GetMainTranslation()

	dutchTranslation := mainTranslation.GetLangTranslation("nl-BE")

	/* Main menu */
	actionTranslation := dutchTranslation.GetActionTranslation("Action")
	actionTranslation.SetTranslation("menu-home", "Start")
	actionTranslation.SetTranslation("menu-netmail", "Netmail")
	actionTranslation.SetTranslation("menu-echomail", "Echomail")
	actionTranslation.SetTranslation("menu-files", "Bestanden")
	actionTranslation.SetTranslation("menu-service", "Diensten")
	actionTranslation.SetTranslation("menu-people", "Adresboek")
	actionTranslation.SetTranslation("menu-draft", "Concept")
	actionTranslation.SetTranslation("menu-setup", "Instellingen")

	/* NetmailIndexAction */
	netmailIndexActionTranslation := dutchTranslation.GetActionTranslation("NetmailIndexAction")
	netmailIndexActionTranslation.SetTranslation("action-button-create", "Nieuw bericht")
	/* EchoAreaIndexAction */
	echoAreaIndexActionTranslation := dutchTranslation.GetActionTranslation("EchoAreaIndexAction")
	echoAreaIndexActionTranslation.SetTranslation("action-button-create", "Maak een nieuw gebied")
	/* EchoMsgIndexAction */
	echoMsgIndexActionTranslation := dutchTranslation.GetActionTranslation("EchoMsgIndexAction")
	echoMsgIndexActionTranslation.SetTranslation("action-compose-button", "Nieuw bericht")
	echoMsgIndexActionTranslation.SetTranslation("action-tree-button", "Boomstructuur")
	echoMsgIndexActionTranslation.SetTranslation("action-mark-as-read-button", "Markeer als gelezen")
	echoMsgIndexActionTranslation.SetTranslation("action-settings-button", "Instellingen")

	/* FileEchoIndexAction */
	fileEchoIndexActionTranslation := dutchTranslation.GetActionTranslation("FileEchoIndexAction")
	fileEchoIndexActionTranslation.SetTranslation("action-button-create", "Creëren")

}
