package http_server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// PlayerServer is a HTTP interface for player information
type PlayerServer struct {
	Store        PlayerStore
	http.Handler // embedding http.Handler interface as it consist of ServeHTTP method
}

// NewPlayerServer creates a PlayerServer with routing configured
func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.Store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	p.Handler = router // here we assign concrete implementation of ServeHTTP as http.NewServeMux has this method

	return p
}

// PlayerStore stores score information about players
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() League
}

func (p *PlayerServer) processWin(rw http.ResponseWriter, player string) {
	p.Store.RecordWin(player)
	rw.WriteHeader(http.StatusCreated)
}

func (p *PlayerServer) showScore(rw http.ResponseWriter, player string) {
	score := p.Store.GetPlayerScore(player)
	if score == 0 {
		rw.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(rw, score)
}

func (p *PlayerServer) leagueHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("content-type", "application/json")
	json.NewEncoder(rw).Encode(p.Store.GetLeague())
	rw.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) playersHandler(rw http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodGet:
		p.showScore(rw, player)
	case http.MethodPost:
		p.processWin(rw, player)
	}
}
