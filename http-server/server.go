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
}

func (p *PlayerServer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	score := p.store.GetPlayerScore(player)
	if score == 0 {
		rw.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(rw, score)
}

//Start from here https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server#write-the-test-first-3
