package config

//import "gopkg.in/yaml.v2"

type Config struct {
	HttpPort int
}

func LoadConfiguration() Config {
	config := Config{}
	config.HttpPort = defaultHttpPort()
	return config
}

func defaultHttpPort() int {
	return 8080
}