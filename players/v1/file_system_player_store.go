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

func NewLeague(rdr io.Reader) (League, error) {
	var result League

	err := json.NewDecoder(rdr).Decode(&result)
	if err == io.EOF {
		return League{}, nil
	}

	return result, err
}

func (fsps FileSystemPlayerStore) GetLeague() League {
	fsps.Database.Seek(0, 0)

	var result, err = NewLeague(fsps.Database)

	if err != nil {
		log.Panicf("Error decoding database content: %v", err)
	}

	return result
}

func (fsps FileSystemPlayerStore) GetPlayerScore(playerName string) int {
	var playerScores, _ = NewLeague(fsps.Database)

	player := playerScores.Find(playerName)
	if player != nil {
		return player.Wins
	}

	return -1
}

func (fsps FileSystemPlayerStore) RecordWin(playerName string) {
	playerStats := fsps.GetLeague()

	player := playerStats.Find(playerName)

	if player == nil {
		playerStats = append(playerStats, serializers.Player{Name: playerName, Wins: 1})
	} else {
		player.Wins++
	}

	fsps.Database.Seek(0, 0)
	json.NewEncoder(fsps.Database).Encode(playerStats)
	fsps.Database.Seek(0, 0)
}
