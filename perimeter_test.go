package main

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("square case", func(t *testing.T) {
		square := Rectangle{height: 10.0, width: 10.0}
		got := Perimeter(square)
		want := 40.0

		if got != want {
			t.Errorf("Expected: %v, got: %v", want, got)
		}
	})

	t.Run("rectangle case", func(t *testing.T) {
		rectangle := Rectangle{height: 12.0, width: 10.0}
		got := Perimeter(rectangle)
		want := 44.0

		if got != want {
			t.Errorf("Expected: %v, got: %v", want, got)
		}
	})
}
