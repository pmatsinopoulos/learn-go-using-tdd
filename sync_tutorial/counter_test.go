package sync_tutorial

import "testing"

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter by 3 leaves it to 3", func(t *testing.T) {
		counter := Counter{}
		numberOfTimesToIncrement := 3
		for i := 0; i < numberOfTimesToIncrement; i++ {
			counter.Inc()
		}

		got := counter.Val()
		if got != numberOfTimesToIncrement {
			t.Errorf("Expected %v, got %v", numberOfTimesToIncrement, got)
		}
	})
}
