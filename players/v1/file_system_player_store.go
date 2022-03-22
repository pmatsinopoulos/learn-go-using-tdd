package v1

import (
	"encoding/json"
	"github.com/pmatsinopoulos/players/v1/serializers"
	"io"
	"log"
)

type FileSystemPlayerStore struct {
	Database io.Reader
}

func NewLeague(rdr io.Reader) ([]serializers.Player, error) {
	var result []serializers.Player

	err := json.NewDecoder(rdr).Decode(&result)

	return result, err
}

func (fsps FileSystemPlayerStore) GetLeague() []serializers.Player {
	var result, err = NewLeague(fsps.Database)

	if err != nil {
		log.Panicf("Error decoding database content: %v", err)
	}

	return result
}
