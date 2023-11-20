package alns

import (

	"math"
	"sync"
	"test-alns/algorithms/sa"
	"test-alns/common"
	"test-alns/common/constant"
)

var DestroyMap map[string]Destroy
var RepairMap map[string]Repair
var OperatorUsageTimes sync.Map

type ALNS struct {
	DestroyScoreBoard map[string]float64 // 计分板
	RepairScoreBoard  map[string]float64 // 计分板
	HistoricallyBest  float64            // 历史最有解
	BestPath          []int              // 最优路径
	Nodes             []common.Node      // 城市节点
	SettingIteration  int
	NowPath           []int              // 当前解的路径
	NowValue          float64            // 当前解大小
	Cut               int                // 破坏个数
	Sa                *sa.SA             // 模拟退火接受准则
	DestroyWeights    map[string]float64 // 算子权重
	RepairWeights     map[string]float64 // 算子权重
	OperatorUsageTimes map[string]int
}

func NewALNS(cut, iteration int, sa *sa.SA) *ALNS {
	alns := &ALNS{
		RepairScoreBoard:  make(map[string]float64),
		DestroyScoreBoard: make(map[string]float64),
		HistoricallyBest:  math.MaxFloat64,
		Cut:               cut,
		SettingIteration:  iteration,
		Sa:                sa,
		DestroyWeights:    make(map[string]float64),
		RepairWeights:     make(map[string]float64),
		OperatorUsageTimes: make(map[string]int),
	}
	
	for k := range DestroyMap {
		alns.DestroyWeights[k] = 1.0
	}
	for k := range RepairMap {
		alns.RepairWeights[k] = 1.0
	}
	return alns
}

func (alns *ALNS) Run() {

	// init path
	alns.InitPath(constant.InitSolutionByGreedy)
	for i := 1; i <= alns.SettingIteration; i++ {
		// destroy and repair
		destroyFuncName := common.RouletteSelect(alns.DestroyWeights)
		repairFuncName := common.RouletteSelect(alns.RepairWeights)
		destroy, repair := DestroyMap[destroyFuncName], RepairMap[repairFuncName]
		tmpPath, destroyed := destroy.Func(alns.NowPath, alns.Cut)
		newPath := repair.Func(tmpPath, destroyed)
		newValue := common.CalcTSP(newPath)
		alns.OperatorUsageTimes[destroyFuncName] ++
		alns.OperatorUsageTimes[repairFuncName] ++

		// update score
		if newValue <= alns.NowValue {
			alns.NowPath = newPath
			alns.NowValue = newValue
			if newValue <= alns.HistoricallyBest {
				alns.HistoricallyBest = newValue
				alns.BestPath = newPath
				alns.DestroyScoreBoard[destroyFuncName] = 1.5
				alns.RepairScoreBoard[repairFuncName] = 1.5
			} else {
				alns.DestroyScoreBoard[destroyFuncName] = 1.2
				alns.RepairScoreBoard[repairFuncName] = 1.2
			}
		} else {
			if alns.Sa.Accept(newValue - alns.NowValue) {
				alns.NowPath = newPath
				alns.NowValue = newValue
				alns.DestroyScoreBoard[destroyFuncName] = 0.8
				alns.RepairScoreBoard[repairFuncName] = 0.8
			} else {
				alns.DestroyScoreBoard[destroyFuncName] = 0.5
				alns.RepairScoreBoard[repairFuncName] = 0.5
			}
		}

		// update weight
		alns.CalcWeight()

		// update temperature
		alns.Sa.NowTemperature *= alns.Sa.CoolingRate
		// // print middle data
		// fmt.Println("当前迭代次数:", i, " 历史最优解:", alns.HistoricallyBest, " 当前解:", alns.NowValue)
	}
	// fmt.Println(accept)
}

// 计算权重：ρ = λρ + (1 - λ)Ψ
func (alns *ALNS) CalcWeight() {
	lambda := constant.WeightCoefficient
	// 计算破坏算子权重
	for k := range DestroyMap {
		alns.DestroyWeights[k] = lambda*alns.DestroyWeights[k] + (1-lambda)*alns.DestroyScoreBoard[k]
	}
	// 计算修复算子权重
	for k := range RepairMap {
		alns.RepairWeights[k] = lambda*alns.RepairWeights[k] + (1-lambda)*alns.RepairScoreBoard[k]
	}
}

// 随机排列
func (alns *ALNS) InitPath(method string) {
	path := make([]int, 0)
	for i := 0; i < len(common.DistMatrix); i ++ {
		path = append(path, i)
	}
	switch method {
	case constant.InitSolutionByRandom:
		shuffle(path)
	case constant.InitSolutionByGreedy:
		greedy(path)
	}
	alns.NowPath = path
	alns.NowValue = common.CalcTSP(path)
}

func shuffle(path []int) {
	for i := 0; i < len(path); i ++ {
		j := common.RandInt(0, len(path) - 1)
		path[i], path[j] = path[j], path[i]
	}
}

func greedy(path []int) {
	new := GreedyRepair2(make([]int, 0), path)
	copy(path, new)
}

// func greedy(path []int) {
// 	bestGreedy := math.MaxFloat64
// 	for i := 0; i < len(path); i ++ {
// 		newPath := make([]int, 0)
// 		newPath = append(newPath, path[i])
// 		for len(newPath) != len(path) {
// 			minDist := math.MaxFloat64
// 			var next int
// 			for _, node := range path {
// 				dist := common.DistMatrix[node][newPath[len(newPath) - 1]]
// 				if dist < minDist && common.NotIn(newPath, node) {
// 					next = node
// 					minDist = dist
// 				}
// 			}
// 			newPath = append(newPath, next)
// 		}
// 		if common.CalcTSP(newPath) < bestGreedy {
// 			bestGreedy = common.CalcTSP(newPath)
// 			for i := 0; i < len(path); i ++ {
// 				path[i] = newPath[i]
// 			}
// 		}
// 	}
// }

