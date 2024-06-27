package model

import "pokemonDraw/common"

type BasePokemon struct {
	Id         uint32
	Level      uint32
	BaseStatus []uint16   //努力值
	IV         [][]uint16 //个体值 含最高和最低
	Exp        uint32
	Ability    []bool //特性
	IsShiny    []bool
	Gender     []bool

	Nature []uint16 //性格
	Ball   []uint16
	Moves  []*common.Move

	//TODO
	Item []uint32

	//IsForm     bool
	//IsAbility  bool
	//IsStatus   bool
	//IsEgg      bool
	//IsAlolan   bool
	//IsMega     uint8
	//IsZ        bool
	//IsDynamax  bool
	//IsGmax     bool
	//IsBattleBond bool
	//IsUltraBurst bool
	//IsPrimal bool
	//IsMegaX bool
	//IsMegaY bool
	//IsGigantamax bool
	//IsAshGreninja bool
}

func newPokemon(uid uint32) *BasePokemon {
	pokemon := &BasePokemon{
		Id:         uid,
		Level:      1,
		BaseStatus: make([]uint16, 6),
		IV:         make([][]uint16, 6),
		Ability:    make([]bool, 3),
		IsShiny:    make([]bool, 2),
		Gender:     make([]bool, 2),
		Nature:     make([]uint16, 0),
		Ball:       make([]uint16, 0),
		Moves:      make([]*common.Move, 0),
	}

	for index := range pokemon.IV {
		pokemon.IV[index] = make([]uint16, 2)
	}

	return pokemon
}

func (u *PlayerUnit) AddNewPokemon(args common.PokemonArgs) {
	u.lock.Lock()
	defer u.lock.Unlock()

	pokemon := u.Pokemons[args.Uid]
	if pokemon == nil {
		pokemon = newPokemon(args.Uid)
		u.Pokemons[args.Uid] = pokemon
	}

	if pokemon.Level < args.Level {
		pokemon.Level = args.Level
		//TODO 经验值变化
	}

	//保存个体值最大值和最小值
	for index, value := range args.IV {
		if value >= pokemon.IV[index][1] {
			pokemon.IV[index][1] = value
		}
		if value <= pokemon.IV[index][0] {
			pokemon.IV[index][0] = value
		}
	}

	//保存特性
	pokemon.Ability[args.Ability] = true

	//保存闪光
	if args.IsShiny {
		pokemon.IsShiny[1] = true
	} else {
		pokemon.IsShiny[0] = true
	}

	//保存性别
	pokemon.Gender[args.Gender] = true

	//保存性格
	addNewNature := true
	for _, nature := range pokemon.Nature {
		if nature == args.Nature {
			addNewNature = false
			break
		}
	}

	if addNewNature {
		pokemon.Nature = append(pokemon.Nature, args.Nature)
	}

	//保存精灵球
	addNewBall := true
	for _, ball := range pokemon.Ball {
		if ball == args.Ball {
			addNewBall = false
			break
		}
	}

	if addNewBall {
		pokemon.Ball = append(pokemon.Ball, args.Ball)
	}

	//保存招式
	for _, move := range args.Moves {
		addNewMove := true
		for _, m := range pokemon.Moves {
			if m == move {
				if m.PPplus+1 <= common.MovePulsMax {
					m.PPplus++
				}
				addNewMove = false
				break
			}
		}

		if addNewMove {
			pokemon.Moves = append(pokemon.Moves, move)
		}
	}

	//TODO 保存道具

	//TODO 保存形态

}
