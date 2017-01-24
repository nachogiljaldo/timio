package config

import (
	"github.com/ghodss/yaml"
	"log"
	"io/ioutil"
	"io"
)

type Config struct {
	HttpPort int `json:"httpport"`
}

func LoadConfiguration(fileName string) (Config, error) {
	log.Printf("Configuration file is %s", fileName)
	config := Config{HttpPort:-1}
	bytes, err := ioutil.ReadFile(fileName)
	if (err != nil) {
		return config, io.EOF
	}
	yaml.Unmarshal(bytes, &config)
	if (config.HttpPort == -1) {
		log.Printf("Using default port %d", defaultHttpPort())
		config.HttpPort = defaultHttpPort()
	}
	return config, nil
}

func defaultHttpPort() int {
	return 80
}