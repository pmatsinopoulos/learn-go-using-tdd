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

func (fsps FileSystemPlayerStore) GetLeague() []serializers.Player {
	var result []serializers.Player

	error := json.NewDecoder(fsps.Database).Decode(&result)

	if error != nil {
		log.Panicf("Error decoding database content: %v", error)
	}

	return result
}
