package main

import "testing"

func TestSum(t *testing.T) {
	arrayOfIntegers := [4]int{1, 2, 3, 4}
	got := Sum(arrayOfIntegers)
	want := 10
	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}
