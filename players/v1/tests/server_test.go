package tests

import (
	"fmt"
	"github.com/pmatsinopoulos/players/v1"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		v1.PlayerServer(response, request)

		want := "20"
		assertResponseBody(t, response.Body.String(), want)
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		v1.PlayerServer(response, request)

		assertResponseBody(t, response.Body.String(), "10")
	})
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func newGetScoreRequest(name string) (request *http.Request) {
	request, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)

	return
}
