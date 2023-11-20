package sa

import (
	"math"
	"test-alns/common"
)

type SA struct {
	CoolingRate        float64
	InitialTemperature float64
	NowTemperature     float64

}

func NewSA(temperature, coolingRate float64) *SA {
	return &SA{
		CoolingRate: coolingRate,
		InitialTemperature: temperature,
		NowTemperature: temperature,
	}
}

func (sa SA) Accept(diff float64) bool {
	probability := math.Exp(-diff / sa.NowTemperature) 
	return common.RandDecimal() < probability
}
