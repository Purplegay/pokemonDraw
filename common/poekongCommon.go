package common

type PokemonArgs struct {
	Uid     uint32
	Level   uint32
	IV      []uint16
	Ability uint8
	IsShiny bool
	Gender  uint8
	Nature  uint16
	Ball    uint16
	Moves   []*Move
}
