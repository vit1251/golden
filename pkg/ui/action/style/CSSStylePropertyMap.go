package style

type CSSStylePropertyMap struct {
	propertyMap map[string]string
}

func NewCSSStylePropertyMap() *CSSStylePropertyMap {
	pm := new(CSSStylePropertyMap)
	pm.propertyMap = make(map[string]string)
	return pm
}

func (self *CSSStylePropertyMap) Set(name string, value string) {
	self.propertyMap[name] = value
}

