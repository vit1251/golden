package style

import "fmt"

type CSSStyleSheet struct {
	Rules     CSSRuleList
}

func NewCSSStyleSheet() *CSSStyleSheet {
	s := new(CSSStyleSheet)
	return s
}

func (self *CSSStyleSheet) InsertRule(rule *CSSRule) {
	self.Rules = append(self.Rules, *rule)
}

func (self *CSSStyleSheet) String() string {
	var result string
	for _, rule := range self.Rules {
		//
		fmt.Printf("rule = %+q", rule)
		//
		result += fmt.Sprintf("%s {\n", rule.selectorText)
		for _, property := range rule.styleMap.propertyMap {
			result += fmt.Sprintf("    %s: %s;\n", property.Name, property.Value)
		}
		result += "}\n"
	}
	return result
}
