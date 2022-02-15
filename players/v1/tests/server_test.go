package tests

import (
	"fmt"
	"github.com/pmatsinopoulos/players/v1"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func TestGetPlayerScore(t *testing.T) {
	t.Run("it returns 0 if player given does not exist", func(t *testing.T) {
		playerStore := StubPlayerStore{scores: map[string]int{}}
		got := playerStore.GetPlayerScore("Non-Existing Player")
		want := 0
		if got != want {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{scores: map[string]int{
		"Pepper": 20,
		"Floyd":  10,
	}}
	server := v1.PlayerServer{PlayerStore: store}

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
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

func assertStatus(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}
