package config

import (
	"io/ioutil"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Service Service `yaml:"service"`
	Server  Server  `yaml:"server"`
	Mqtt    Mqtt    `yaml:"mqtt"`
}

type Service struct {
	Name string `yaml:"name"`
}

type Server struct {
	Address string `yaml:"address"`
}

type Mqtt struct {
	Broker   string    `yaml:"broker"`
	Port     int       `yaml:"port"`
	ClientID string    `yaml:"client_id"`
	Event    MqttEvent `ymal:"event"`
}
type MqttEvent struct {
	Calendar string `yaml:"calendar"`
}

func LoadConfig(configFile string) (*Config, error) {
	var cfg Config
	if err := loadConfigFile(configFile, &cfg); err != nil {
		return nil, err
	}

	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	return nil, nil
}

func loadConfigFile(configFile string, cfg *Config) error {
	_, err := os.Stat(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	ymlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(ymlFile, &cfg)
}
