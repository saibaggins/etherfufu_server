package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type DBConfig struct {
	URI string `yaml:"uri"`
}

type Configuration struct {
	Database DBConfig `yaml:"database"`
}

var _configuration *Configuration

func LoadEnvConfig() *Configuration {

	getConfigPath := func() string {
		activeEnv := os.Getenv("ACTIVE_ENV")
		fmt.Println("Current active environment is ", activeEnv)
		absPath, _ := filepath.Abs("./config")
		return filepath.Join(absPath, activeEnv+".yml")
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

	_configuration = &config
	return _configuration
}

func GetEnvConfig() *Configuration {
	if _configuration == nil {
		return LoadEnvConfig()
	}

	return _configuration
}
