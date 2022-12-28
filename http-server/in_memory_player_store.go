package poker

import "sync"

//func NewInMemoryPlayerStore() *InMemoryPlayerStore {
//	return &InMemoryPlayerStore{store: map[string]int{}, mu: sync.Mutex{}}
//}

type InMemoryPlayerStore struct {
	store map[string]int
	mu    sync.Mutex
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	defer i.mu.Unlock()
	i.mu.Lock()
	return i.store[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mu.Lock()
	i.store[name]++
	i.mu.Unlock()
}
