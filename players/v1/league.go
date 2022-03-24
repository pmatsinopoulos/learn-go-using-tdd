package v1

import "github.com/pmatsinopoulos/players/v1/serializers"

type League []serializers.Player

func (l League) Find(name string) *serializers.Player {
	for index, player := range l {
		if player.Name == name {
			return &l[index]
		}
	}
	return nil
}
