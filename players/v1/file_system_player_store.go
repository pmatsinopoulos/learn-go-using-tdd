package v1

import (
	"encoding/json"
	"fmt"
	"github.com/pmatsinopoulos/players/v1/serializers"
	"io"
	"os"
	"sort"
)

type FileSystemPlayerStore struct {
	Database *json.Encoder
	league   League
}

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	file.Seek(0, 0)
	league, err := newLeague(file)

	if err != nil {
		return nil, fmt.Errorf("problem loadking player store from file %s, error: %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{
		Database: json.NewEncoder(&Tape{File: file}),
		league:   league,
	}, nil
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
	sort.Slice(fsps.league, func(i, j int) bool {
		return fsps.league[i].Wins > fsps.league[j].Wins
	})
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

	fsps.Database.Encode(fsps.league)
}
