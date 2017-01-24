package api

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"github.com/pborman/uuid"
	"timio/persistence"
	"timio/timer"
	"timio/config"
	"strconv"
)

type Api struct {
	persister persistence.Persister
	httpPort int
}

func Initialize(persister persistence.Persister, config config.Config) Api {
	api := Api{}
	api.persister = persister
	api.httpPort = config.HttpPort
	return api
}

func (api Api) Start() {
	router := mux.NewRouter()
	createTimer := func(w http.ResponseWriter, req *http.Request) {
		api.AddTimer(w, req)
	}
	router.HandleFunc("/timers", createTimer).Methods("POST")
	log.Printf("Timio will listen on port %d", api.httpPort)
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(api.httpPort), router))
}

func (api Api) AddTimer(w http.ResponseWriter, req *http.Request) {
	var timer timer.Timer
	_ = json.NewDecoder(req.Body).Decode(&timer)
	timer.ID = uuid.New()
	api.persister.Persist(timer)
	w.Header().Add(TIMER_ID_HEADER, timer.ID)
	w.WriteHeader(http.StatusNoContent)
}
