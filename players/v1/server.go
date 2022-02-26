package v1

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerServer struct {
	PlayerStore PlayerStore
	router      *http.ServeMux
}

// PlayerServer Factory function

func NewPlayerServer(store PlayerStore) PlayerServer {
	p := PlayerServer{
		PlayerStore: store,
		router:      http.NewServeMux(),
	}
	p.router.Handle("/league", p.leagueHandler())

	p.router.Handle("/players/", p.playersHandler())
	return p
}

// public methods

func (p PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.router.ServeHTTP(w, r)
}

// private methods

func (p PlayerServer) leagueHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
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
