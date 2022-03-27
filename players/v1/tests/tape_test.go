package tests

import (
	v1 "github.com/pmatsinopoulos/players/v1"
	"io/ioutil"
	"testing"
)

func TestTapeWrite(t *testing.T) {
	file, clean := createTempFile(t, "12345")
	defer clean()

	tape := v1.Tape{File: file}

	tape.Write([]byte("abc"))

	file.Seek(0, 0)

	newFileContents, _ := ioutil.ReadAll(file)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
