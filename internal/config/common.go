package config

import (
	"github.com/alhaos/cliker/internal/webService"
	"github.com/ilyakaznacheev/cleanenv"
)

type Configuration struct {
	WebServer *webService.WebServerConfig `yaml:"webService"`
}

func FromFile(confFilename string) (*Configuration, error) {

	conf := Configuration{}

	err := cleanenv.ReadConfig(confFilename, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, err
}
