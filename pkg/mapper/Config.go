package mapper

type Config struct {
	params []ConfigValue
}

func NewConfig() *Config {
	return new(Config)
}

func (self Config) Get(section string, name string) (string, bool) {
	for _, param := range self.params {
		if param.section == section && param.name == name {
			return param.value, true
		}
	}
	return "", false
}

func (self Config) GetParams() []ConfigValue {
	return self.params
}

func (self *Config) Set(section string, name string, value string) {
	var newValues []ConfigValue
	var updated bool = false
	for _, param := range self.params {
		if param.section == section && param.name == name {
			newParam := ConfigValue{
				section: section,
				name:    name,
				value:   value,
			}
			newValues = append(newValues, newParam)
			updated = true
		} else {
			newValues = append(newValues, param)
		}
	}
	if !updated {
		newParam := ConfigValue{
			section: section,
			name:    name,
			value:   value,
		}
		newValues = append(newValues, newParam)
	}
	self.params = newValues
}
