package config

import (
	"encoding/json"
	"errors"
	"os"
)

type Mongo struct {
	Host string `json:"host"`
	Port string `json:"port"`
	DB   string `json:"db"`
}

type ServiceAddress struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type Config struct {
	Service    string         `json:"service"`
	Connector  ServiceAddress `json:"connector"`
	Mongo      Mongo          `json:"mongo"`
	Debug      bool           `json:"debug"`
	OnlyErrors bool           `json:"onlyErrors"`
	Address    ServiceAddress `json:"address"`
	Clients    []string       `json:"clients"`
}

func GetConfig(rootDir string) (*Config, error) {
	f, err := os.OpenFile(rootDir+"/config.json", os.O_RDONLY, 0)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	c := Config{}

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&c)

	if err != nil {
		return nil, err
	}

	if c.Address.Host == "" || c.Address.Port == 0 {
		return nil, errors.New("empty address params in config")
	}

	return &c, nil
}
