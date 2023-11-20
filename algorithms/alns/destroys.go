package alns

import (
	"test-alns/common"
)

type DesTroyFunc func(path []int, cnt int) ([]int, []int)

type Destroy struct {
	Name string
	Func DesTroyFunc
}

func RandomDestroy(path []int, cnt int) ([]int, []int) {
	tmpPath := make([]int, 0)
	destroyed := make([]int, 0)
	// 生成随机不重复的cnt个点
	for i := 0; i < cnt; i ++ {
		for {
			removeNode := common.RandInt(0, len(path) - 1)
			if common.NotIn(destroyed, removeNode) {
				destroyed = append(destroyed, removeNode)
				break
			}
		}	
	}
	for _, node := range path {
		if common.NotIn(destroyed, node) {
			tmpPath = append(tmpPath, node)
		}
	}
	return tmpPath, destroyed
}

func GreedyDestroy(path []int, cnt int) ([]int, []int) {
	tmpPath := make([]int, 0)
	destroyed := make([]int, 0)
	savingMap := make(map[int]float64)
	// 计算每个节点的节约值
	for _, node := range path {
		pathWithoutNode := PathWithoutNode(node, path)
		saving := common.CalcTSP(path) - common.CalcTSP(pathWithoutNode)
		savingMap[node] = saving
	}
	// 计算节约值最大的cnt个节点
	for i := 0; i < cnt; i ++ {
		maxSaving := 0.0
		var tobeDestroyed int
		for node, saving := range savingMap {
			if saving > maxSaving {
				maxSaving = saving
				tobeDestroyed = node
			}
		}
		destroyed = append(destroyed, tobeDestroyed)
		delete(savingMap, tobeDestroyed)
	}
	// 得到剩余的路径
	for _, node := range path {
		if common.NotIn(destroyed, node) {
			tmpPath = append(tmpPath, node)
		}
	}
	return tmpPath, destroyed
}

func MaxDestroy(path []int, cnt int) ([]int, []int) {
	tmpPath := make([]int, 0)
	destroyed := make([]int, 0)
	distMap := make(map[int]float64)
	for i := 0; i < len(path) - 1; i ++ {
		distMap[i] = common.DistMatrix[path[i]][path[i + 1]]
	}
	distMap[len(path) - 1] = common.DistMatrix[path[0]][path[len(path) - 1]]
	// 计算距离最大的cnt个节点
	for i := 0; i < cnt; i ++ {
		maxDist := 0.0
		var tobeDestroyed int
		for node, saving := range distMap {
			if saving > maxDist {
				maxDist = saving
				tobeDestroyed = node
			}
		}
		destroyed = append(destroyed, tobeDestroyed)
		delete(distMap, tobeDestroyed)
	}
	// 得到剩余的路径
	for _, node := range path {
		if common.NotIn(destroyed, node) {
			tmpPath = append(tmpPath, node)
		}
	}
	return tmpPath, destroyed
}

func ContinuousDestroy(path []int, cnt int) ([]int, []int) {
	tmpPath := make([]int, 0)
	destroyed := make([]int, 0)
	start := common.RandInt(0, len(path)-cnt) 
	destroyed = append(destroyed, path[start:start+cnt]...)
	tmpPath = append(tmpPath, path[:start]...) 
	tmpPath = append(tmpPath, path[start+cnt:]...)
	return tmpPath, destroyed
}

func PathWithoutNode(node int, path []int) []int {
	pathWithoutNode := make([]int, 0)
	for _, this := range path {
		if this != node {
			pathWithoutNode = append(pathWithoutNode, this)
		}
	}
	return pathWithoutNode
}

func RegisterDestroy() {
	DestroyMap = make(map[string]Destroy)
	DestroyMap["random_destroy"] = Destroy{Name: "random_destroy", Func: RandomDestroy}
	DestroyMap["greedy_destroy"] = Destroy{Name: "greedy_destroy", Func: GreedyDestroy}
	DestroyMap["continuous_destroy"] = Destroy{Name: "continuous_destroy", Func: ContinuousDestroy}
	// DestroyMap["max_destroy"] = Destroy{Name: "max_destroy", Func: MaxDestroy}
}