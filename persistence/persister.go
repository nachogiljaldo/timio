package persistence

import (
	"log"
	"timio/timer"
)

type Persister struct {
}

func (persister Persister) Persist(timer timer.Timer) int8 {
	log.Printf("Received timer to persist: %s", timer.ID)
	return 0
}
