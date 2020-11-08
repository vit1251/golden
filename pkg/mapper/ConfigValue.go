package mapper

type ParamType string

const (
	ParamString   ParamType = "STRING"
	ParamInt      ParamType = "INT"
	ParamDuration ParamType = "DURATION"
	ParamBool     ParamType = "BOOL"
)

type ConfigValue struct {
	Summary string    /* Parameter summary     */
	Section string    /* Parameter section     */
	Name    string    /* Parameter name        */
	value   string    /* Parameter value       */
	Type    ParamType /* Parameter value type  */
	update  bool
}

func NewConfigValue() *ConfigValue {
	newConfigValue := new(ConfigValue)
	newConfigValue.update = false
	return newConfigValue
}

func (self *ConfigValue) GetValue() string {
	return self.value
}

func (self *ConfigValue) SetValue(value string) {
	self.value = value
	self.update = true
}

func (self *ConfigValue) IsUpdate() bool {
	return self.update
}

func (self *ConfigValue) SetUpdate(update bool) {
	self.update = update
}
