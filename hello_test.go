package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"
	if got != want {
		t.Errorf("actual %q, different from expected %q", got, want)
	}
}
