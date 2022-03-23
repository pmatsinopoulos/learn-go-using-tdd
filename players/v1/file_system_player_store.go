package v1

import (
	"encoding/json"
	"github.com/pmatsinopoulos/players/v1/serializers"
	"io"
	"log"
)

type FileSystemPlayerStore struct {
	Database io.ReadWriteSeeker
}

func NewLeague(rdr io.Reader) ([]serializers.Player, error) {
	var result []serializers.Player

	err := json.NewDecoder(rdr).Decode(&result)

	return result, err
}

func (fsps FileSystemPlayerStore) GetLeague() []serializers.Player {
	fsps.Database.Seek(0, 0)

	var result, err = NewLeague(fsps.Database)

	if err != nil {
		log.Panicf("Error decoding database content: %v", err)
	}

	return result
}

func (fsps FileSystemPlayerStore) GetPlayerScore(playerName string) int {
	var playerScores, _ = NewLeague(fsps.Database)

	for _, player := range playerScores {
		if player.Name == playerName {
			return player.Wins
		}
	}
	return -1
}

func (fsps FileSystemPlayerStore) RecordWin(playerName string) {
	playerStats := fsps.GetLeague()

	for index, player := range playerStats {
		if player.Name == playerName {
			playerStats[index].Wins++
		}
	}

	fsps.Database.Seek(0, 0)
	json.NewEncoder(fsps.Database).Encode(playerStats)
	fsps.Database.Seek(0, 0)
}
