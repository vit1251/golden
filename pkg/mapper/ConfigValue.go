package mapper

type ConfigValue struct {
	section string /* Parameter section     */
	name    string /* Parameter name        */
	value   string /* Parameter value       */
}

func (self *ConfigValue) GetValue() string {
	return self.value
}
