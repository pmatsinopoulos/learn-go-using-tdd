package _select

import (
	"net/http"
)

func Racer(a, b string) (fast string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

func ping(url string) chan struct{} {
	result := make(chan struct{})
	go func() {
		http.Get(url)
		close(result)
	}()
	return result
}
