package main

import (
	"timio/persistence"
	"timio/api"
	"timio/config"
	"os"
	"log"
	"timio/arguments"
)

func main() {
	configFile := arguments.GetConfigFile(os.Args[1:])
	configuration, err := config.LoadConfiguration(configFile)
	if err != nil {
		log.Fatalf("Could not read configuration file %s", configFile)
	}
	persister := persistence.Persister{}
	api.Initialize(persister, configuration).Start()
}