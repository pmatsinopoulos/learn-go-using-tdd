package v1

import "github.com/pmatsinopoulos/players/v1/serializers"

type PlayerStore interface {
	GetLeague() []serializers.Player
	GetPlayerScore(name string) int
	RecordWin(name string)
}
