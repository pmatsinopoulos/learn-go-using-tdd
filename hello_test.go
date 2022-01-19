package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("actual %q, different from expected %q", got, want)
		}
	}

	t.Run("saying hello to a given name", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello uses the default World name when name is not given", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}
