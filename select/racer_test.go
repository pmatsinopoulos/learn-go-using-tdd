package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := makeDelayedServer(20)
	defer slowServer.Close()

	fastServer := makeDelayedServer(0)
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("Expected %q, got %q", want, got)
	}
}

func makeDelayedServer(millis time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if millis > 0 {
			time.Sleep(millis * time.Millisecond)
		}
		w.WriteHeader(http.StatusOK)
	}))
}
