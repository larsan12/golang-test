package config

import (
	"log"
	"os"

	_ "embed"

	"github.com/pkg/errors"
	"go.uber.org/zap"
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

func init() {
	// config init
	err := Init()
	if err != nil {
		log.Fatal("Unable to connect init config", zap.Error(err))
	}
}

//go:embed config.yml
var TestConfig string

func Init() error {
	isTest, exists := os.LookupEnv("IS_TEST")

	if exists && isTest == "true" {
		err := yaml.Unmarshal([]byte(TestConfig), &ConfigData)
		if err != nil {
			return errors.WithMessage(err, "parsing yaml")
		}
		return nil
	}

	rawYAML, err := os.ReadFile("./config.yml")
	if err != nil {
		return errors.WithMessage(err, "reading config file")
	}

	err = yaml.Unmarshal(rawYAML, &ConfigData)
	if err != nil {
		return errors.WithMessage(err, "parsing yaml")
	}

	return nil
}
