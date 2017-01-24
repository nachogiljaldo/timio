package main

import (
	"timio/persistence"
	"timio/api"
	"timio/config"
)

func main() {
	config := config.LoadConfiguration()
	persister := persistence.Persister{}
	api := api.Initialize(persister, config)
	api.Start()
}