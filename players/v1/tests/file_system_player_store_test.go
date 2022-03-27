package tests

import (
	v1 "github.com/pmatsinopoulos/players/v1"
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

		store := v1.NewFileSystemPlayerStore(database)

		got := store.GetLeague()

		want := v1.League{
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

		store := v1.NewFileSystemPlayerStore(database)

		got := store.GetPlayerScore("Chris")

		assertScoreEquals(t, got, 33)

		got = store.GetPlayerScore("Cleo")

		assertScoreEquals(t, got, 10)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `
          [
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}
          ]
        `)
		defer cleanDatabase()

		store := v1.NewFileSystemPlayerStore(database)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")

		assertScoreEquals(t, got, 34)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `
           [
             {"Name": "Chris", "Wins": 10},
             {"Name": "Peter", "Wins": 31}
           ]
        `)
		defer cleanDatabase()

		store := v1.NewFileSystemPlayerStore(database)

		// fire
		store.RecordWin("Panos")

		got := store.GetPlayerScore("Panos")

		assertScoreEquals(t, got, 1)

		league := store.GetLeague()

		found := league.Find("Panos")

		assertScoreEquals(t, found.Wins, 1)
	})
}

// ---------------------------------
// Helper Functions

func assertLeagueMatching(t *testing.T, got v1.League, want v1.League) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
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
