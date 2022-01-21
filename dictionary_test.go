package main

import "testing"

func TestSearch(t *testing.T) {
    t.Run("-", func(t *testing.T) {
        dictionary := map[string]string{"test": "this is just a test"}

        got := Search(dictionary, "test")
        want := "this is just a test"

        if got != want {
            t.Errorf("Expected %q, actual %q", want, got)
        }
    })
}