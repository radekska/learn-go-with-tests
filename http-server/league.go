package http_server

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
)

type League []Player

func (l League) FindPlayer(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}

func (l League) Sort() {
	sort.Slice(l, func(i, j int) bool {
		return l[i].Wins > l[j].Wins
	})
}

func NewLeague(rdr io.Reader) (League, error) {
	var league League
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}
	return league, err
}
