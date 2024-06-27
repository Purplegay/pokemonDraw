package model

import "sync"

type PlayerUnit struct {
	lock     *sync.RWMutex
	Pokemons map[uint32]*BasePokemon
}
