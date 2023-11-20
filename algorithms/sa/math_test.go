package sa

import (
	"fmt"
	"math"
	"testing"
)

func TestMath(t *testing.T) {
	T := 50000.0
	rate := 0.99
	for i := 0; i < 10000; i ++ {
		f := math.Exp(-100/T)
		fmt.Println(f)
		T *= rate
	}
}

