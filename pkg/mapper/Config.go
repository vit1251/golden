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
