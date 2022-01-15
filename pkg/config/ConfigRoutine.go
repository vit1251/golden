package config

type Param struct {
	Section string
	Name    string
	Prt     *string
}

func GetParams() []Param {
	var params []Param = []Param{
		/* Networking */
		{"main", "Address", nil},
		{"main", "NetAddr", nil},
		{"main", "Password", nil},
		{"main", "Link", nil},
		/* User options */
		{"main", "RealName", nil},
		{"main", "Country", nil},
		{"main", "City", nil},
		{"main", "Origin", nil},
		{"main", "TearLine", nil},
		/* Netmail */
		{"netmail", "Charset", nil},
		/* Echomail */
		{"echomail", "Charset", nil},
		/* Other */
		{"main", "StationName", nil},
		{"mailer", "Interval", nil},
	}
	return params
}

func GetValuePtrByPath(c *Config, section string, name string) *string {
	var params []Param = []Param{
		/* Networking */
		{"main", "Address", &c.Main.Address},
		{"main", "NetAddr", &c.Main.NetAddr},
		{"main", "Password", &c.Main.Password},
		{"main", "Link", &c.Main.Link},
		/* User options */
		{"main", "RealName", &c.Main.RealName},
		{"main", "Country", &c.Main.Country},
		{"main", "City", &c.Main.City},
		{"main", "Origin", &c.Main.Origin},
		{"main", "TearLine", &c.Main.TearLine},
		/* Netmail */
		{"netmail", "Charset", &c.Netmail.Charset},
		/* Echomail */
		{"echomail", "Charset", &c.Echomail.Charset},
		/* Other */
		{"main", "StationName", &c.Main.StationName},
		{"mailer", "Interval", &c.Mailer.Interval},
	}
	for _, a := range params {
		if a.Section == section && a.Name == name {
			return a.Prt
		}
	}
	return nil
}

func GetByPath(c *Config, section string, name string) string {
	value := GetValuePtrByPath(c, section, name)
	if value != nil {
		return *value
	} else {
		return ""
	}
}

func SetByPath(c *Config, section string, name string, value string) {
	addr := GetValuePtrByPath(c, section, name)
	if addr != nil {
		*addr = value
	}
}
