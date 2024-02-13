package i18n

func init() {

	mainTranslation := GetMainTranslation()

	spanishTranslation := mainTranslation.GetLangTranslation("es-AR")

	/* Main menu */
	actionTranslation := spanishTranslation.GetActionTranslation("Action")
	actionTranslation.SetTranslation("menu-home", "Inicio")
	actionTranslation.SetTranslation("menu-netmail", "Netmail")
	actionTranslation.SetTranslation("menu-echomail", "Echomail")
	actionTranslation.SetTranslation("menu-files", "Archivos")
	actionTranslation.SetTranslation("menu-service", "Servicios")
	actionTranslation.SetTranslation("menu-people", "Personas")
	actionTranslation.SetTranslation("menu-draft", "Borradores")
	actionTranslation.SetTranslation("menu-setup", "Configuración")

	/* NetmailIndexAction */
	netmailIndexActionTranslation := spanishTranslation.GetActionTranslation("NetmailIndexAction")
	netmailIndexActionTranslation.SetTranslation("action-button-create", "Redactar")
	/* EchoAreaIndexAction */
	echoAreaIndexActionTranslation := spanishTranslation.GetActionTranslation("EchoAreaIndexAction")
	echoAreaIndexActionTranslation.SetTranslation("action-button-create", "Crear")
	/* EchoMsgIndexAction */
	echoMsgIndexActionTranslation := spanishTranslation.GetActionTranslation("EchoMsgIndexAction")
	echoMsgIndexActionTranslation.SetTranslation("action-compose-button", "Redactar")
	echoMsgIndexActionTranslation.SetTranslation("action-tree-button", "Árbol")
	echoMsgIndexActionTranslation.SetTranslation("action-mark-as-read-button", "Marcar como leído")
	echoMsgIndexActionTranslation.SetTranslation("action-settings-button", "Opciones")
	/* FileEchoIndexAction */
	fileEchoIndexActionTranslation := spanishTranslation.GetActionTranslation("FileEchoIndexAction")
	fileEchoIndexActionTranslation.SetTranslation("action-button-create", "Crear")

}
