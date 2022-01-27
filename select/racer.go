package _select

import (
	"net/http"
	"time"
)

func Racer(a, b string) (fast string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		fast = a
	} else {
		fast = b
	}
	return
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
