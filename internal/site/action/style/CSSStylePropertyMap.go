package style

type CSSStyleProperty struct {
	Name  string
	Value string
}

type CSSStylePropertyMap struct {
	propertyMap []CSSStyleProperty
}

func NewCSSStylePropertyMap() *CSSStylePropertyMap {
	pm := new(CSSStylePropertyMap)
	return pm
}

func (self *CSSStylePropertyMap) Set(name string, value string) {
	property := CSSStyleProperty{
		Name: name,
		Value: value,
	}
	self.propertyMap = append(self.propertyMap, property)
}
