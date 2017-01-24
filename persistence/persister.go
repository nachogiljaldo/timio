package persistence

import (
	"log"
	"timio/timer"
)

type Persister struct {
	timers []timer.Timer
}

func (persister *Persister) Persist(timer timer.Timer) error {
	persister.timers = append(persister.timers, timer)
	log.Printf("Have already %d timers", len(persister.timers))
	return nil
}
