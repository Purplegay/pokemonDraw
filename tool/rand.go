package tool

import (
	"math/rand"
	"time"
)

var randSeed = int64(time.Now().Unix()) // 自定义的随机数种子
// TODO
//
//	func Rand[T uint32 | int16 | int32 | int64](items [][]T) {
//		rander := rand.New(rand.NewSource(randSeed))
//	}
func RandUInt32(items [][]uint32) uint32 {
	rander := rand.New(rand.NewSource(randSeed))
	var (
		randItems   []uint32
		totalWeight uint32
	)
	for _, item := range items {
		totalWeight += item[1]
		randItems = append(randItems, item...)
	}

	randWeight := rander.Int63n(int64(totalWeight))
	for i, sz := 0, len(randItems); i < sz; i += 2 {
		if n := randItems[i+1]; randWeight < int64(n) {
			return uint32(randItems[i])
		}
	}

	return 0
}
