package config

import (
	"log"
	"os"
	"path/filepath"

	_ "embed"

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

// func init() {
// 	// config init
// 	err := Init()
// 	if err != nil {
// 		log.Fatal("Unable to connect init config", zap.Error(err))
// 	}
// }

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

	wd, err := os.Getwd()
	if err != nil {
		return errors.WithMessage(err, "getting working directory")
	}
	log.Printf("current working directory: %s", wd)

	// prefer config from checkout directory regardless of CWD
	resolveFromCheckoutDir := func(start string) (string, bool) {
		dir := start
		for {
			base := filepath.Base(dir)
			candidate := filepath.Join(dir, "config.yml")
			if base == "checkout" {
				if _, statErr := os.Stat(candidate); statErr == nil {
					return candidate, true
				}
			}
			parent := filepath.Dir(dir)
			if parent == dir {
				return "", false
			}
			dir = parent
		}
	}

	cfgPath, ok := resolveFromCheckoutDir(wd)
	if !ok {
		// fallback to local relative path
		cfgPath = "./config.yml"
	}

	absConfigPath, _ := filepath.Abs(cfgPath)
	log.Printf("resolved config path: %s", absConfigPath)

	rawYAML, err := os.ReadFile(cfgPath)
	if err != nil {
		return errors.WithMessage(err, "reading config file")
	}

	err = yaml.Unmarshal(rawYAML, &ConfigData)
	if err != nil {
		return errors.WithMessage(err, "parsing yaml")
	}

	return nil
}
