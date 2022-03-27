package v1

import (
	"encoding/json"
	"github.com/pmatsinopoulos/players/v1/serializers"
	"io"
	"os"
)

type FileSystemPlayerStore struct {
	Database io.Writer
	league   League
}

func NewFileSystemPlayerStore(database *os.File) *FileSystemPlayerStore {
	database.Seek(0, 0)
	league, _ := newLeague(database)

	return &FileSystemPlayerStore{
		Database: &Tape{File: database},
		league:   league,
	}
}

func newLeague(rdr io.Reader) (League, error) {
	var result League

	err := json.NewDecoder(rdr).Decode(&result)
	if err == io.EOF {
		return League{}, nil
	}

	return result, err
}

func (fsps FileSystemPlayerStore) GetLeague() League {
	return fsps.league
}

func (fsps FileSystemPlayerStore) GetPlayerScore(playerName string) int {
	player := fsps.league.Find(playerName)
	if player != nil {
		return player.Wins
	}

	return -1
}

func (fsps *FileSystemPlayerStore) RecordWin(playerName string) {
	player := fsps.league.Find(playerName)

	if player == nil {
		fsps.league = append(fsps.league, serializers.Player{Name: playerName, Wins: 1})
	} else {
		player.Wins++
	}

	json.NewEncoder(fsps.Database).Encode(fsps.league)
}
