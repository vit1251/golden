package style

type CSSRule struct {
	selectorText string
	styleMap     CSSStylePropertyMap
}

func NewCSSRule() *CSSRule {
	r := new(CSSRule)
	r.styleMap = *NewCSSStylePropertyMap()
	return r
}

func (self *CSSRule) SetSelectorText(selectorText string) {
	self.selectorText = selectorText
}

func (self *CSSRule) Set(name string, value string) {
	self.styleMap.Set(name, value)
}
