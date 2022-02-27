package v1

import (
	"encoding/json"
	"fmt"
	"github.com/pmatsinopoulos/players/v1/serializers"
	"net/http"
	"strings"
)

type PlayerServer struct {
	PlayerStore  PlayerStore
	http.Handler // embedding
}

// PlayerServer Factory function

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.PlayerStore = store

	router := http.NewServeMux()

	router.Handle("/league", p.leagueHandler())

	router.Handle("/players/", p.playersHandler())

	p.Handler = router

	return p
}

// private methods

func (p PlayerServer) leagueHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(p.getLeagueTable())
	}
}

func (p PlayerServer) getLeagueTable() []serializers.Player {
	return p.PlayerStore.GetLeague()
}

func (p PlayerServer) playersHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			p.handlePOSTRecordWin(w, r)
		case http.MethodGet:
			p.handleGETPlayerScore(w, r)
		}
	}
}

func (p PlayerServer) handlePOSTRecordWin(w http.ResponseWriter, r *http.Request) {
	player := playerName(r)

	w.WriteHeader(http.StatusAccepted)
	p.PlayerStore.RecordWin(player)
}

func (p PlayerServer) handleGETPlayerScore(w http.ResponseWriter, r *http.Request) {
	player := playerName(r)

	score := p.PlayerStore.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	fmt.Fprint(w, score)

}

func playerName(r *http.Request) (player string) {
	player = strings.TrimPrefix(r.URL.Path, "/players/")

	return
}
