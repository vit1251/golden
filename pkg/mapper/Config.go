package mapper

type Config struct {
	params []*ConfigValue
}

func NewConfig() *Config {
	return new(Config)
}

func (self *Config) Insert(section string, name string, value string) {
	newConfigValue := NewConfigValue()
	newConfigValue.Section = section
	newConfigValue.Name = name
	newConfigValue.SetValue(value)
	newConfigValue.SetUpdate(false)
	self.params = append(self.params, newConfigValue)
}

func (self *Config) Set(section string, name string, value string) {
	for _, param := range self.params {
		if param.Section == section && param.Name == name {
			param.SetValue(value)
		}
	}
}

func (self Config) Get(section string, name string) (string, bool) {

	for _, param := range self.params {
		if param.Section == section && param.Name == name {
			return param.GetValue(), true
		}
	}

	return "", false

}

func (self Config) GetParams() []*ConfigValue {
	return self.params
}
