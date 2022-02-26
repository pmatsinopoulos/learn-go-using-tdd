package tests

import (
	"fmt"
	"github.com/pmatsinopoulos/players/v1"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
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

func TestRecordWin(t *testing.T) {
	t.Run("it records the win", func(t *testing.T) {
		playerStore := StubPlayerStore{scores: map[string]int{"foo": 1, "bar": 2}, winCalls: []string{"Mary"}}
		name := "Pepper"
		playerStore.RecordWin(name)
		got := len(playerStore.winCalls)
		want := 2
		if got != want {
			t.Errorf("Expected %v got %v", want, got)
		}
	})
}

//------------------------------------
// Endpoint: GET /players/:playerName
//------------------------------------

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{scores: map[string]int{
		"Pepper": 20,
		"Floyd":  10,
	}}
	server := v1.NewPlayerServer(&store)

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

//------------------------------------
// Endpoint: POST /players/:playerName
//------------------------------------

func TestStoreWins(t *testing.T) {
	t.Run("it returns accepted on POST", func(t *testing.T) {
		store := StubPlayerStore{
			scores:   map[string]int{},
			winCalls: []string{},
		}
		server := v1.NewPlayerServer(&store)

		request, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)
	})

	t.Run("it records wins when POST", func(t *testing.T) {
		store := StubPlayerStore{
			scores:   map[string]int{},
			winCalls: []string{},
		}
		server := v1.NewPlayerServer(&store)

		player := "Pepper"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := len(store.winCalls)
		want := 1

		if got != want {
			t.Errorf("Expected %v, go %v", want, got)
		}

		if store.winCalls[0] != player {
			t.Errorf("Expected the win to be for %q, but it was for %q", player, store.winCalls[0])
		}
	})
}

//------------------------------------
// Endpoint: GET /league
//------------------------------------

func TestLeague(t *testing.T) {
	t.Run("it returns 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()
		store := StubPlayerStore{}

		server := v1.NewPlayerServer(&store)

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
	})
}

//------------------------------------
// test helper methods
//------------------------------------

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

func newPostWinRequest(name string) (request *http.Request) {
	request, _ = http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)

	return
}
