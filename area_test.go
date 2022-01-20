package main

import (
	"fmt"
	"testing"
)

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

	t.Run("circle case", func(t *testing.T) {
		circle := Circle{radius: 1.5}
		got := Area(circle)
		want := 7.06858

		gotStr := fmt.Sprintf("%.5f", got)
		wantStr := fmt.Sprintf("%.5f", want)

		if gotStr != wantStr {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})
}
