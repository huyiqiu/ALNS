package main

import (
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
	nodes := common.ReadTsp("./benchmark/xqf131.tsp")
	// nodes := common.ReadTsp("./benchmark/china34.tsp")
	common.GenerateDistMatrix(nodes, constant.EUC)
	algorithms.Register()
	algorithmsRunTimes := 1
	run(algorithmsRunTimes)
}

func run(times int) float64 {
	var wg sync.WaitGroup
	resultChan := make(chan float64)
	// 算法开始时间
	startTime := time.Now()
	for i := 0; i < times; i ++ {
		wg.Add(1)
		go func(i int){
			defer wg.Done()
			randDestroyNum := 0.3 * float64(len(common.DistMatrix))
			alns := alns.NewALNS(int(randDestroyNum), 1000, sa.NewSA(50000.0, 0.98))
			alns.Run()
			resultChan <- alns.HistoricallyBest
			for k, v := range alns.OperatorUsageTimes {
				fmt.Println(k, ":", v)
			}
			fmt.Println("线程:", i, "最优解:", alns.HistoricallyBest, "最优路径:", alns.BestPath)
		}(i)
	}
	go func() {
		wg.Wait()
		close(resultChan)
	}()
	best := math.MaxFloat64
	for result := range resultChan {
		if result < best {
			best = result
		}
	}
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Println("best:", best)
	fmt.Println("算法执行时间:", elapsedTime)
	return best
}