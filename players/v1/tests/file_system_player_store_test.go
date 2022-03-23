package tests

import (
	v1 "github.com/pmatsinopoulos/players/v1"
	"github.com/pmatsinopoulos/players/v1/serializers"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestFileSystemPlayerStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
          {"Name": "Cleo", "Wins": 10},
          {"Name": "Chris", "Wins": 33}
        ]`)
		defer cleanDatabase()

		store := v1.FileSystemPlayerStore{Database: database}

		got := store.GetLeague()

		want := []serializers.Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		assertLeagueMatching(t, got, want)

		got = store.GetLeague()

		assertLeagueMatching(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `
          [{"Name": "Cleo", "Wins": 10},
           {"Name": "Chris", "Wins": 33}]
        `)
		defer cleanDatabase()

		store := v1.FileSystemPlayerStore{Database: database}

		got := store.GetPlayerScore("Chris")

		assertScoreEquals(t, got, 33)
	})
}

// ---------------------------------
// Helper Functions

func assertLeagueMatching(t *testing.T, got []serializers.Player, want []serializers.Player) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tempFile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tempFile.Write([]byte(initialData))

	removeFile := func() {
		tempFile.Close()
		os.Remove(tempFile.Name())
	}

	tempFile.Seek(0, 0)

	return tempFile, removeFile
}

func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}
