package i18n

var _mainTranslation_ *MainTranslation

func GetMainTranslation() *MainTranslation {

	if _mainTranslation_ == nil {
		_mainTranslation_ = new(MainTranslation)
		_mainTranslation_.mainTranslations = make(map[string]*LangTranslation)
	}

	return _mainTranslation_

}
