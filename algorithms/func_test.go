package algorithms

import (
	"fmt"
	"test-alns/algorithms/sa"
	al "test-alns/algorithms/alns"
	"test-alns/common"
	"testing"
)

func TestConti(t *testing.T) {
	nodes := common.ReadTsp("../benchmark/xqf131.tsp")
	fmt.Println("Test2")
	// nodes := common.ReadTsp("./benchmark/china34.tsp")
	common.GenerateDistMatrix(nodes, "Euc")
	Register()
	alns := al.NewALNS(13, 10, sa.NewSA(100, 0.98))
	alns.InitPath("greedy")
	i, i2 := al.DestroyMap["continuous_destroy"].Func(alns.NowPath, 13)
	fmt.Println(i)
	fmt.Println(i2)
}