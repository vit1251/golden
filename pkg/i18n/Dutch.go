package i18n

func init() {

	mainTranslation := GetMainTranslation()

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
