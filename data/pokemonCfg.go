package data

type PokemonCfg struct {
	Pokemon map[uint32]*PokemonInfo
	Moves   map[uint32]*MoveInfo
}

type PokemonInfo struct {
	Id    uint32
	Gen   uint32 //代
	Level uint32
	// IV         [][]uint16 //个体值 含最高和最低
	ExpType uint32   //经验类型
	Ability []uint32 //特性
	// IsShiny    []bool
	Gender []uint32 //性别权重

	// Nature []uint16 //性格
	Ball  [][]uint16 //可用球种:权重
	Moves []uint32
}

type MoveInfo struct {
	Id    uint32
	Gen   uint32
	Name  string
	Attr  uint32 //属性
	Type  uint32 //类型
	Power uint32
	Acc   uint32 //命中率
	PP    uint32
}

func (p *PokemonCfg) GetPokemonInfo(id uint32) *PokemonInfo {
	return p.Pokemon[id]
}

func (p *PokemonCfg) GetMoveInfo(id uint32) *MoveInfo {
	return p.Moves[id]
}

func (p *PokemonCfg) GetPokemonInfoAll() map[uint32]*PokemonInfo {
	return p.Pokemon
}

func (p *PokemonCfg) GetMoveInfoAll() map[uint32]*MoveInfo {
	return p.Moves
}

func (p *PokemonCfg) GenNewPokemon()
