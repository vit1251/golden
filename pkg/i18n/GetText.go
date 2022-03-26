package i18n

import (
	"os"
	"log"
	"fmt"
	"strings"
	"net/http"
)

var mainTranslation *MainTranslation

type MainTranslation struct {
	mainTranslations map[string]*LangTranslation
}

type LangTranslation struct {
	actionTranslations map[string]*ActionTranslation
}

func (self *LangTranslation) GetActionTranslation(actionName string) *ActionTranslation {
	if result, ok := self.actionTranslations[actionName]; ok {
		return result
	} else {
		result := new(ActionTranslation)
		result.translations = make(map[string]string)
		self.actionTranslations[actionName] = result
		return result
	}
}

type ActionTranslation struct {
	translations map[string]string
}

func (self *ActionTranslation) SetTranslation(codeName string, msg string) {
	self.translations[codeName] = msg
}

func (self *ActionTranslation) GetTranslation(codeName string) string {
	var result string = codeName
	if msg, ok := self.translations[codeName]; ok {
		result = msg
	}
	return result
}

func (self *MainTranslation) GetLangTranslation(lang string) *LangTranslation {
	if result, ok := self.mainTranslations[lang]; ok {
		return result
	} else {
		result := new(LangTranslation)
		result.actionTranslations = make(map[string]*ActionTranslation)
		self.mainTranslations[lang] = result
		return result
	}
}

func init() {

	/* Initialize master mapper */
	mainTranslation = new(MainTranslation)
	mainTranslation.mainTranslations = make(map[string]*LangTranslation)

	/* Russian */
	{
		russianTranslation := mainTranslation.GetLangTranslation("ru-RU")
		/* NetmailIndexAction */
		netmailIndexActionTranslation := russianTranslation.GetActionTranslation("NetmailIndexAction")
		netmailIndexActionTranslation.SetTranslation("action-button-create", "Создать сообщение")
		/* EchoAreaIndexAction */
		echoAreaIndexActionTranslation := russianTranslation.GetActionTranslation("EchoAreaIndexAction")
		echoAreaIndexActionTranslation.SetTranslation("action-button-create", "Создать сообщение")
	}

	/* English */
	{
		englishTranslation := mainTranslation.GetLangTranslation("en-US")
		/* NetmailIndexAction */
		netmailIndexActionTranslation := englishTranslation.GetActionTranslation("NetmailIndexAction")
		netmailIndexActionTranslation.SetTranslation("action-button-create", "Compose")
		/* EchoAreaIndexAction */
		echoAreaIndexActionTranslation := englishTranslation.GetActionTranslation("EchoAreaIndexAction")
		echoAreaIndexActionTranslation.SetTranslation("action-button-create", "Compose")
	}

}

func GetText(langName string, actionName string, codeName string) string {

	var result string = codeName

	newLangTranslation := mainTranslation.GetLangTranslation(langName)
	newActionTranslation := newLangTranslation.GetActionTranslation(actionName)
	result = newActionTranslation.GetTranslation(codeName)

	return result

}

type Lang struct {
	lang1 string
	lang2 string
	charset string
}

func parseLang(lang string) Lang {

	var l Lang

	parts := strings.SplitN(lang, ".", 2)

	if len(parts) >= 2 {
		l.charset = parts[1]
	}

	if len(parts) >= 1 {

		code := parts[0]

		langs := strings.SplitN(code, "_", 2)

		if len(langs) >= 2 {
			l.lang2 = langs[1]
		}

		if len(langs) >= 1 {
			l.lang1 = langs[0]
		}
	}

	return l
}

func GetDefaultLanguage() string {

	var result string = "en-US"

	// LANG=ru_RU.UTF-8
	if lang, exists := os.LookupEnv("LANG"); exists {
		l := parseLang(lang)
		log.Printf("lang = %#v", l)
		result = fmt.Sprintf("%s-%s", l.lang1, l.lang2)
	}

	return result
}

func GetLangNameFromRequest(r *http.Request) string {
	var mainLanguage string = GetDefaultLanguage()
	if acceptLanguages, ok := r.Header["Accept-Language"]; ok {
		for _, acceptLanguage := range acceptLanguages {
			log.Printf("User accept: %s", acceptLanguage)
			// TODO - parse "ru,en-US;q=0.9,en;q=0.8" params...
		}
	}
	return mainLanguage
}
