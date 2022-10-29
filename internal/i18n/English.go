package i18n

func init() {

	mainTranslation := GetMainTranslation()

	englishTranslation := mainTranslation.GetLangTranslation("en-US")

	/* Main menu */
	actionTranslation := englishTranslation.GetActionTranslation("Action")
	actionTranslation.SetTranslation("menu-home", "Home")
	actionTranslation.SetTranslation("menu-netmail", "Netmail")
	actionTranslation.SetTranslation("menu-echomail", "Echomail")
	actionTranslation.SetTranslation("menu-files", "Files")
	actionTranslation.SetTranslation("menu-service", "Service")
	actionTranslation.SetTranslation("menu-people", "People")
	actionTranslation.SetTranslation("menu-draft", "Draft")
	actionTranslation.SetTranslation("menu-setup", "Setup")

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
