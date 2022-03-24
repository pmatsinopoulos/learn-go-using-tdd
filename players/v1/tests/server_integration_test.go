package tests

import (
	"encoding/json"
	v1 "github.com/pmatsinopoulos/players/v1"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := v1.NewInMemoryPlayerStore()
	server := v1.NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "3")
}

func TestRecordingWinsAndGettingLeague(t *testing.T) {
	store := v1.NewInMemoryPlayerStore()
	server := v1.NewPlayerServer(store)

	player := "Peter"
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	player = "Sam"
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	// Get the League

	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	var got v1.League

	json.NewDecoder(response.Body).Decode(&got)

	want := v1.League{
		{Name: "Peter", Wins: 3},
		{Name: "Sam", Wins: 4},
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected %v, got %v", want, got)
	}
}
