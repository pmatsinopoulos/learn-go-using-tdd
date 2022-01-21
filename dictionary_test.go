package main

import (
	"fmt"
	"testing"
)

var assertStrings = func(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Expected %q, actual %q", want, got)
	}
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{map[string]string{"test": "this is just a test"}}

	t.Run("existing word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("foo")
		if err == nil {
			t.Fatal("should raise an error, but it didn't")
		}
		errString := fmt.Sprintf("%v", err)
		expectedError := "word not found"
		assertStrings(t, errString, expectedError)
	})
}
