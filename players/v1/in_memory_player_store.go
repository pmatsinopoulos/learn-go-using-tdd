package v1

import "github.com/pmatsinopoulos/players/v1/serializers"

type InMemoryPlayerStore struct {
	store map[string]int
}

func (i *InMemoryPlayerStore) GetLeague() League {
	result := make(League, 0, len(i.store))

	for player, wins := range i.store {
		result = append(result, serializers.Player{Name: player, Wins: wins})
	}

	return result
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		store: map[string]int{},
	}
}
