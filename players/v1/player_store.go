package v1

type PlayerStore interface {
	GetLeague() League
	GetPlayerScore(name string) int
	RecordWin(name string)
}
