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
	Value   string    /* Parameter value       */
	Type    ParamType /* Parameter value type  */
}

func (self *ConfigValue) SetValue(value string) {
	self.Value = value
}

