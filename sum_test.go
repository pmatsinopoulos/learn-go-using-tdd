package main

import "testing"

func TestSum(t *testing.T) {
	assertResults := func(got int, want int) {
		t.Helper()
		if got != want {
			t.Errorf("Expected %v, got %v", want, got)
		}
	}

	t.Run("collection of any size", func(t *testing.T) {
		arrayOfIntegers := []int{8, 3}
		got := Sum(arrayOfIntegers)
		want := 11
		assertResults(got, want)
	})
}
