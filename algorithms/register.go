package algorithms

import "test-alns/algorithms/alns"

func Register() {
	// 注册算子
	alns.RegisterDestroy()
	alns.RegisterRepair()
}