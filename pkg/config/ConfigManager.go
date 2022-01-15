package config

import (
	"bytes"
	"errors"
	"github.com/BurntSushi/toml"
	ctx "github.com/vit1251/golden/pkg/common"
	"log"
	"os"
	"path"
)

type ConfigManager struct {
}

func NewConfigManager() *ConfigManager {
	return new(ConfigManager)
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

func (self *ConfigManager) GetConfigPath() string {
	baseDir := ctx.GetFidoDirectory()
	return path.Join(baseDir, "config.toml")
}

func (self *ConfigManager) Store(c *Config) error {

	/* Encode settings */
	buf := new(bytes.Buffer)
	if err1 := toml.NewEncoder(buf).Encode(c); err1 != nil {
		return err1
	}
	data := buf.Bytes()

	/* Prepare ou directory */
	newConfigPath := self.GetConfigPath()

	/* Store output */
	err2 := os.WriteFile(newConfigPath, data, 0644)
	if err2 != nil {
		return err2
	}

	return nil
}

func (self *ConfigManager) Restore() (*Config, error) {

	var c Config

	/* Prepare ou directory */
	newConfigPath := self.GetConfigPath()

	/* Restore output */
	data, err1 := os.ReadFile(newConfigPath)
	if err1 != nil {
		return &c, err1
	}

	err2 := toml.Unmarshal(data, &c)
	if err2 != nil {
		return &c, err2
	}

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

func (self *ConfigManager) IsNotExists() bool {
	newConfigPath := self.GetConfigPath()
	_, err1 := os.Stat(newConfigPath)
	log.Printf("err = %#v", err1)
	if errors.Is(err1, os.ErrNotExist) {
		return true
	} else {
		return false
	}
}

func (self *ConfigManager) Reset() {
	config = nil
}
