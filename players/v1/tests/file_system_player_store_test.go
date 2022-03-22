package tests

import (
	v1 "github.com/pmatsinopoulos/players/v1"
	"github.com/pmatsinopoulos/players/v1/serializers"
	"reflect"
	"strings"
	"testing"
)

func TestFileSystemPlayerStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database := strings.NewReader(`[
          {"Name": "Cleo", "Wins": 10},
          {"Name": "Chris", "Wins": 33}
        ]`)

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
}

func assertLeagueMatching(t *testing.T, got []serializers.Player, want []serializers.Player) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v, got %v", want, got)
	}
}
