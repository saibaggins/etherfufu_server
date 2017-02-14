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

func GetConfigPath() string {
	activeEnv := os.Getenv("ACTIVE_ENV")
	fmt.Println("Current active environment is ", activeEnv)
	absPath, _ := filepath.Abs("./config")
	return filepath.Join(absPath, activeEnv+".yml")
}

func LoadEnvConfig() *Configuration {
	file, err := ioutil.ReadFile(GetConfigPath())
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
	return _configuration
}
