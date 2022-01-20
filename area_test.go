package main

import "testing"

func TestArea(t *testing.T) {
	t.Run("square case", func(t *testing.T) {
		got := Area(6.0, 6.0)
		want := 36.0
		if got != want {
			t.Errorf("Expected: %v, actual: %v", want, got)
		}
	})

	t.Run("rectangle case", func(t *testing.T) {
		got := Area(12.0, 6.0)
		want := 72.0
		if got != want {
			t.Errorf("Expected: %v, actual: %v", want, got)
		}
	})
}
