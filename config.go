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

type Mysql struct {
	Host     string `json:"host"`
	DB       string `json:"db"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type ServiceAddress struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type Config struct {
	Service    string         `json:"service"`
	Address    ServiceAddress `json:"address"`
	Connector  ServiceAddress `json:"connector"`
	Clients    []string       `json:"clients"`
	Mongo      Mongo          `json:"mongo"`
	Mysql      Mysql          `json:"mysql"`
	AccessKey  string         `json:"accessKey"`
	RefreshKey string         `json:"refreshKey"`
	OnlyErrors bool           `json:"onlyErrors"`
	Debug      bool           `json:"debug"`
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
