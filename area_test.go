package main

import "testing"

func TestArea(t *testing.T) {
	t.Run("square case", func(t *testing.T) {
		square := Rectangle{height: 6.0, width: 6.0}
		got := Area(square)
		want := 36.0
		if got != want {
			t.Errorf("Expected: %v, actual: %v", want, got)
		}
	})

	t.Run("rectangle case", func(t *testing.T) {
		rectangle := Rectangle{height: 12.0, width: 6.0}
		got := Area(rectangle)
		want := 72.0
		if got != want {
			t.Errorf("Expected: %v, actual: %v", want, got)
		}
	})
}
