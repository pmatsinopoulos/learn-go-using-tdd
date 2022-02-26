package v1

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type PlayerServer struct {
	PlayerStore PlayerStore
}

func (p PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		p.handlePOST(w, r)
	case http.MethodGet:
		p.handleGET(w, r)
	}
}

func (p PlayerServer) handlePOST(w http.ResponseWriter, r *http.Request) {
	player := playerName(r)

	w.WriteHeader(http.StatusAccepted)
	p.PlayerStore.RecordWin(player)
}

type allowedPath string

func (ap allowedPath) String() string {
	switch ap {
	case league:
		return "/league"
	case playersIndividualPlayer:
		return "/players/"
	}
	return "unknown"
}

const (
	league                  allowedPath = "/league"
	playersIndividualPlayer allowedPath = "/players/"
)

func path(r *http.Request) string {
	return r.URL.Path
}

func (p PlayerServer) handleGET(w http.ResponseWriter, r *http.Request) {
	if match, _ := regexp.MatchString(string(playersIndividualPlayer), path(r)); match {
		player := playerName(r)

		score := p.PlayerStore.GetPlayerScore(player)
		if score == 0 {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)
		}

		fmt.Fprint(w, score)
	} else if match, _ := regexp.MatchString(string(league), path(r)); match {
		w.WriteHeader(http.StatusOK)
	}

}

func playerName(r *http.Request) (player string) {
	player = strings.TrimPrefix(r.URL.Path, "/players/")

	return
}
