package config

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type ConfigStruct struct {
	KafkaBrokers []string `yaml:"kafkaBrokers"`
	KafkaTopic   string   `yaml:"kafkaTopic"`
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
