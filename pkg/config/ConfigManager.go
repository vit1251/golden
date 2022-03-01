package config

import (
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/registry"
	"log"
)

type ConfigManager struct {
	registry *registry.Container
}

func NewConfigManager(r *registry.Container) *ConfigManager {
	newConfigManager := new(ConfigManager)
	newConfigManager.registry = r
	return newConfigManager
}

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

func (self *ConfigManager) Store(c *Config) error {

	mapperManager := self.restoreMapperManager()
	configMapper := mapperManager.GetConfigMapper()
	outdateConfig, _ := configMapper.GetConfigFromDatabase()

	/* Main */
	outdateConfig.Set("main", "Address", c.Main.Address)
	outdateConfig.Set("main", "Password", c.Main.Password)
	outdateConfig.Set("main", "Origin", c.Main.Origin)
	outdateConfig.Set("main", "TearLine", c.Main.TearLine)
	outdateConfig.Set("main", "Link", c.Main.Link)
	outdateConfig.Set("main", "StationName", c.Main.StationName)
	outdateConfig.Set("main", "RealName", c.Main.RealName)
	outdateConfig.Set("main", "NetAddr", c.Main.NetAddr)
	outdateConfig.Set("main", "City", c.Main.City)
	outdateConfig.Set("main", "Country", c.Main.Country)

	/* Mailer */
	outdateConfig.Set("mailer", "Interval", c.Mailer.Interval)

	/* Netmail */
	outdateConfig.Set("netmail", "Charset", c.Netmail.Charset)

	/* Echomail */
	outdateConfig.Set("echomail", "Charset", c.Echomail.Charset)

	configMapper.SetConfigToDatabase(outdateConfig)

	return nil
}

func (self *ConfigManager) Restore() (*Config, error) {

	var c Config

	mapperManager := self.restoreMapperManager()
	configMapper := mapperManager.GetConfigMapper()
	outdateConfig, _ := configMapper.GetConfigFromDatabase()

	/* Main */
	c.Main.Address, _ = outdateConfig.Get("main", "Address")
	c.Main.Password, _ = outdateConfig.Get("main", "Password")
	c.Main.Origin, _ = outdateConfig.Get("main", "Origin")
	c.Main.TearLine, _ = outdateConfig.Get("main", "TearLine")
	c.Main.Link, _ = outdateConfig.Get("main", "Link")
	c.Main.StationName, _ = outdateConfig.Get("main", "StationName")
	c.Main.RealName, _ = outdateConfig.Get("main", "RealName")
	c.Main.NetAddr, _ = outdateConfig.Get("main", "NetAddr")
	c.Main.City, _ = outdateConfig.Get("main", "City")
	c.Main.Country, _ = outdateConfig.Get("main", "Country")

	/* Mailer */
	c.Mailer.Interval, _ = outdateConfig.Get("mailer", "Interval")

	/* Netmail */
	c.Netmail.Charset, _ = outdateConfig.Get("netmail", "Charset")

	/* Echomail */
	c.Echomail.Charset, _ = outdateConfig.Get("echomail", "Charset")

	return &c, nil
}

func (self *ConfigManager) GetConfig() *Config {
	if config == nil {
		newConfig, err1 := self.Restore()
		if err1 != nil {
			log.Printf("Restore config with error: err = %#v", err1)
		}
		config = newConfig
	}
	return config
}

func (self *ConfigManager) restoreMapperManager() *mapper.MapperManager {
	managerPtr := self.registry.Get("MapperManager")
	if manager, ok := managerPtr.(*mapper.MapperManager); ok {
		return manager
	} else {
		panic("no mapper manager")
	}
}
