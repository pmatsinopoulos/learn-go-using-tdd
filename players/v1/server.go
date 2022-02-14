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
	fmt.Fprint(w, p.PlayerStore.GetPlayerScore(player))
}
