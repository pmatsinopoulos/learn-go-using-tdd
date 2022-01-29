package sync_tutorial

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	assertCounter := func(t testing.TB, counter Counter, numberOfIncrements int) {
		t.Helper()

		if counter.Val() != numberOfIncrements {
			t.Errorf("Expected %v, got %v", numberOfIncrements, counter.Val())
		}
	}

	t.Run("#Val() returns the internal counter", func(t *testing.T) {
		counter := Counter{8}
		got := counter.Val()
		if got != 8 {
			t.Errorf("Expected 8 got %v", got)
		}
	})

	t.Run("incrementing the counter by 3 leaves it to 3", func(t *testing.T) {
		counter := Counter{}
		numberOfTimesToIncrement := 3
		for i := 0; i < numberOfTimesToIncrement; i++ {
			counter.Inc()
		}

		assertCounter(t, counter, numberOfTimesToIncrement)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		counter := Counter{}
		numberOfTimesToIncrement := 1000
		wg := sync.WaitGroup{}
		wg.Add(numberOfTimesToIncrement)

		for i := 0; i < numberOfTimesToIncrement; i++ {
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				counter.Inc()
			}(&wg)
		}
		wg.Wait()

		assertCounter(t, counter, numberOfTimesToIncrement)
	})
}
