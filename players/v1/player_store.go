package v1

type PlayerStore interface {
	GetPlayerScore(name string) int
}
