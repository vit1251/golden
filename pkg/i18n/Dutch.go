package i18n

func init() {

	mainTranslation := GetMainTranslation()

	dutchTranslation := mainTranslation.GetLangTranslation("nl-BE")

	/* Main menu */
	actionTranslation := englishTranslation.GetActionTranslation("Actie")
	actionTranslation.SetTranslation("menu-home", "Thuis")
	actionTranslation.SetTranslation("menu-netmail", "Netmail")
	actionTranslation.SetTranslation("menu-echomail", "Echomail")
	actionTranslation.SetTranslation("menu-files", "Bestanden")
	actionTranslation.SetTranslation("menu-service", "Diensten")
	actionTranslation.SetTranslation("menu-people", "Adresboek")
	actionTranslation.SetTranslation("menu-draft", "Voorlopige versie")
	actionTranslation.SetTranslation("menu-setup", "Instellingen")

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

	/* FileEchoIndexAction */
	fileEchoIndexActionTranslation := englishTranslation.GetActionTranslation("FileEchoIndexAction")
	fileEchoIndexActionTranslation.SetTranslation("action-button-create", "Creëren")

}


