package config

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type ConfigStruct struct {
	Port                string   `yaml:"port"`
	Db                  string   `yaml:"db"`
	OrderExpirationTime int      `yaml:"orderExpirationTime"`
	KafkaBrokers        []string `yaml:"kafkaBrokers"`
	KafkaTopic          string   `yaml:"kafkaTopic"`
	TracesUrl           string   `yaml:"tracesUrl"`
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
