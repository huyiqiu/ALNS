package alns

import (
	"math"
	"test-alns/common"
)

type RepairFunc func(path []int, destroyed []int) []int

type Repair struct {
	Name string
	Func RepairFunc
}

func RandomRepair(path []int, destroyed []int) []int {
	afterRepair := path
	for _, node := range destroyed {
		afterRepair = RandomInsert(node, afterRepair)
	}
	return afterRepair
}

func GreedyRepair(path []int, destroyed []int) []int {
	afterRepair := path
	for _, node := range destroyed {
		afterRepair = GreedyInsert(node, afterRepair)
	}
	return afterRepair
}

func GreedyRepair2(x []int, removedCities []int) []int {
	dis := math.Inf(1)
	insertIndex := -1

	for i := 0; i < len(removedCities); i++ {
		// 寻找插入后的最小总距离
		for j := 0; j <= len(x); j++ {
			newX := make([]int, len(x)+1)
			copy(newX, x)
			newX = append(newX[:j], append([]int{removedCities[i]}, newX[j:]...)...)
			if common.CalcTSP(newX) < dis {
				dis = common.CalcTSP(newX)
				insertIndex = j
			}
		}

		// 最小位置处插入
		x = append(x[:insertIndex], append([]int{removedCities[i]}, x[insertIndex:]...)...)
		dis = math.Inf(1)
	}

	return x
}

func RandomInsert(node int, path []int) []int {
	position := common.RandInt(0, len(path))
	newPath := IndexInsert(node, position, path)
	return newPath
}

func IndexInsert(node int, position int, path []int) []int {
	newPath := make([]int, 0)
	newPath = append(newPath, path[:position]...)
	newPath = append(newPath, node)
	if position < len(path) {
		newPath = append(newPath, path[position:]...)
	}
	return newPath
}

func GreedyInsert(node int, path []int) []int {
	increaseMap := make(map[int]float64)
	beforeInsert := common.CalcTSP(path)
	for p := 0; p <= len(path); p ++ {
		afterInsert := common.CalcTSP(IndexInsert(node, p, path))
		increaseMap[p] = afterInsert - beforeInsert
	}
	minimalIncrese := math.MaxFloat64
	var bestPosition int
	for k, v := range increaseMap {
		if v < minimalIncrese {
			minimalIncrese = v
			bestPosition = k
		}
	}
	return IndexInsert(node, bestPosition, path)
}

func RegisterRepair() {
	RepairMap = make(map[string]Repair)
	// RepairMap["random_repair"] = Repair{Name: "random_repair", Func: RandomRepair}
	RepairMap["greedy_repair"] = Repair{Name: "greedy_repair", Func: GreedyRepair}
	// RepairMap["greedy_repair2"] = Repair{Name: "greedy_repair2", Func: GreedyRepair2}
}