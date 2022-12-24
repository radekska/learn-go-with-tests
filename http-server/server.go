package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerServer struct {
	store PlayerStore
}
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

func (p *PlayerServer) processWin(rw http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	rw.WriteHeader(http.StatusCreated)
}

func (p *PlayerServer) showScore(rw http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		rw.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(rw, score)
}

func (p *PlayerServer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodGet:
		p.showScore(rw, player)
	case http.MethodPost:
		p.processWin(rw, player)
	}
}
