package main

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("square case", func(t *testing.T) {
		got := Perimeter(10.0, 10.0)
		want := 40.0

		if got != want {
			t.Errorf("Expected: %v, got: %v", want, got)
		}
	})

	t.Run("rectangle case", func(t *testing.T) {
		got := Perimeter(12.0, 10.0)
		want := 44.0

		if got != want {
			t.Errorf("Expected: %v, got: %v", want, got)
		}
	})
}
