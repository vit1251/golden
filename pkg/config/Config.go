package config

type Main struct {
	NetAddr     string
	Password    string
	Address     string /* Point FTN address */
	Country     string
	City        string
	RealName    string
	StationName string
	Link        string /* Boss FTN address */
	TearLine    string /* Message TearLine */
	Origin      string /* Message Origin */
}

type Mailer struct {
	Interval string
}

type Netmail struct {
	Charset string
}

type Echomail struct {
	Charset string
}

type Config struct {
	Main     Main     `toml:"Main"`
	Mailer   Mailer   `toml:"Mailer"`
	Netmail  Netmail  `toml:"Netmail"`
	Echomail Echomail `toml:"Echomail"`
}

var config *Config
