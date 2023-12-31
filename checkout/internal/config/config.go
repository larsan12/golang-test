package config

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type ConfigStruct struct {
	Token    string `yaml:"token"`
	Services struct {
		Loms    string `yaml:"loms"`
		Product string `yaml:"product"`
	} `yaml:"services"`
	Port                     string `yaml:"port"`
	Db                       string `yaml:"db"`
	ProductServiceRateLiming uint   `yaml:"productServiceRateLiming"`
	GetProductPoolAmount     int    `yaml:"getProductPoolAmount"`
	TracesUrl                string `yaml:"tracesUrl"`
}

var ConfigData ConfigStruct

func Init() error {
	rawYAML, err := os.ReadFile("config.yml")
	if err != nil {
		return errors.WithMessage(err, "reading config file")
	}

	err = yaml.Unmarshal(rawYAML, &ConfigData)
	if err != nil {
		return errors.WithMessage(err, "parsing yaml")
	}

	return nil
}
