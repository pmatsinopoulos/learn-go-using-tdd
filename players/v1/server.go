package v1

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerServer struct {
	PlayerStore PlayerStore
}

func (p PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	score := p.PlayerStore.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	fmt.Fprint(w, score)
}
