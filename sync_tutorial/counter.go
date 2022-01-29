package sync_tutorial

import "sync"

var mu sync.Mutex

type Counter struct {
	count int
}

func (counter *Counter) Inc() {
	mu.Lock()
	defer mu.Unlock()
	counter.count++
}

func (counter *Counter) Val() int {
	return counter.count
}
