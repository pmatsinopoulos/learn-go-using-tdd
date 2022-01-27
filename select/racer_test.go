package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("it returns the fastest responding site", func(t *testing.T) {
		slowServer := makeDelayedServer(20)
		defer slowServer.Close()

		fastServer := makeDelayedServer(0)
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("Expected %q, got %q", want, got)
		}
	})

	t.Run("it times out after specific seconds", func(t *testing.T) {
		serverA := makeDelayedServer(600)
		serverB := makeDelayedServer(600)

		defer serverA.Close()
		defer serverB.Close()

		_, err := ConfigurableRacer(serverA.URL, serverB.URL, 500*time.Millisecond)
		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(millis time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if millis > 0 {
			time.Sleep(millis * time.Millisecond)
		}
		w.WriteHeader(http.StatusOK)
	}))
}
