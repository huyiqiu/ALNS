package main

import (
	"flag"
	"fmt"
	"math"
	"sync"
	"test-alns/algorithms"
	"test-alns/algorithms/alns"
	"test-alns/algorithms/sa"
	"test-alns/common"
	"test-alns/common/constant"
	"time"
)

func main() {
	var filepath string
	var coor string
	flag.StringVar(&filepath, "filepath", "./benchmark/xqf131.tsp", "指定tsp文件路径")
	flag.StringVar(&coor, "coor", constant.EUC, "指定坐标系")
	flag.Parse()

	nodes := common.ReadTsp(filepath)
	common.GenerateDistMatrix(nodes, coor)
	algorithms.Register()
	algorithmsRunTimes := 5
	run(algorithmsRunTimes)
}

func run(times int) {
	var wg sync.WaitGroup
	var result sync.Map
	// 算法开始时间
	startTime := time.Now()
	for i := 0; i < times; i ++ {
		wg.Add(1)
		go func(i int){
			defer wg.Done()
			randDestroyNum := 0.3 * float64(len(common.DistMatrix))
			alns := alns.NewALNS(int(randDestroyNum), 1000, sa.NewSA(50000.0, 0.98))
			alns.Run()
			result.Store(i, alns.BestPath)
			for k, v := range alns.OperatorUsageTimes {
				fmt.Println(k, ":", v)
			}
			fmt.Println("线程:", i, "最优解:", alns.HistoricallyBest, "最优路径:", alns.BestPath)
		}(i)
	}
	wg.Wait()
	bestvalue := math.MaxFloat64
	var bestpath []int
	result.Range(func(key, value interface{}) bool {
		if common.CalcTSP(value.([]int)) < bestvalue {
			bestpath = value.([]int)
			bestvalue = common.CalcTSP(value.([]int))
		}
		return true
	})
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Println("best:", bestvalue)
	fmt.Println("best path:", bestpath)
	fmt.Println("算法执行时间:", elapsedTime)
}