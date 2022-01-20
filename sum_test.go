package main

import (
	"fmt"
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	assertResults := func(got []int, want []int) {
		errorMessage := fmt.Sprintf("Expected %v, got %v", want, got)
		if !reflect.DeepEqual(got, want) {
			t.Errorf(errorMessage)
		}
	}

	t.Run("returns a slice with the sums of each input slice - case: 1 slice", func(t *testing.T) {
		slice1 := []int{1, 1, 1}
		got := SumAll(slice1)
		want := []int{3}
		assertResults(got, want)
	})

	t.Run("returns a slice with the sums of each input slice - case: 2 slices", func(t *testing.T) {
		slice1 := []int{1, 2}
		slice2 := []int{0, 9}
		got := SumAll(slice1, slice2)
		want := []int{3, 9}
		assertResults(got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	assertResults := func(got []int, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v, got %v", want, got)
		}
	}

	t.Run("case 1: 1 slice given", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3})
		want := []int{5}
		assertResults(got, want)
	})

	t.Run("case 2: 2 slices given", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{2, 8, 20})
		want := []int{5, 28}
		assertResults(got, want)
	})

	t.Run("case 3: 3 slices given, one is empty", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{2, 8, 20}, []int{})
		want := []int{5, 28, 0}
		assertResults(got, want)
	})
}
