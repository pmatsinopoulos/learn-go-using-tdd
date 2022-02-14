package v1

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, scoreForPlayer(player))
}

func scoreForPlayer(player string) string {
	score := ""
	switch player {
	case "Floyd":
		score = "10"
	case "Pepper":
		score = "20"
	default:
		score = "0"
	}
	return score
}
