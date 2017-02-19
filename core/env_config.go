package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Configuration struct {
	Database struct {
		URI string `yaml:"uri"`
	} `yaml:"database"`

	AWS struct {
		AccessKey    string `yaml:"access_key"`
		APISecretKey string `yaml:"api_secret_key"`
	} `yaml:"aws"`
}

var _configuration *Configuration

func ActiveENV() string {
	return os.Getenv("ACTIVE_ENV")
}

func EnvConfig() *Configuration {
	if _configuration == nil {
		_configuration = loadEnvConfig()
	}

	return _configuration
}

func loadEnvConfig() *Configuration {

	getBaseConfigPath := func() string {
		if defaultConfigPath := os.Getenv("CONFIG_PATH"); len(defaultConfigPath) > 0 {
			return defaultConfigPath
		}
		return "./config"
	}

	getConfigPath := func() string {
		fmt.Println("Current active environment is ", ActiveENV())
		absPath, _ := filepath.Abs(getBaseConfigPath())
		return filepath.Join(absPath, ActiveENV()+".yml")
	}

	file, err := ioutil.ReadFile(getConfigPath())
	if err != nil {
		log.Fatal("Config File Missing. ", err)
	}

	var config Configuration
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Config Parse Error: ", err)
	}

	return &config
}
