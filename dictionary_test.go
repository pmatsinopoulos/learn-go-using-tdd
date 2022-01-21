package main

import "testing"

var assertStrings = func(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Expected %q, actual %q", want, got)
	}
}

func TestSearch(t *testing.T) {
	t.Run("-", func(t *testing.T) {
		dictionary := Dictionary{map[string]string{"test": "this is just a test"}}

		got := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})
}
