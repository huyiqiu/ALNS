package common

import (
	"math"
	"test-alns/common/constant"
)


type Node struct {
	Id int
	X  float64
	Y  float64
}

var DistMatrix [][]float64

func GenerateDistMatrix(nodes []Node, edgeType string) {
	num := len(nodes)
    DistMatrix = make([][]float64, num)
    for i := range DistMatrix {
        DistMatrix[i] = make([]float64, num)
    }
	for i := 0; i < num; i ++ {
		for j := i + 1; j < num; j ++ {
			switch edgeType {
			case constant.EUC:
				DistMatrix[i][j] = EUCDistance(nodes[i].X, nodes[i].Y, nodes[j].X, nodes[j].Y)
				DistMatrix[j][i] = DistMatrix[i][j]
			case constant.WGS:
				DistMatrix[i][j] = WGSDistance(nodes[i].Y, nodes[i].X, nodes[j].Y, nodes[j].X)
				DistMatrix[j][i] = DistMatrix[i][j]
			}
		}
	}
}

const earthRadiusKm = 6378.14

// degreeToRadian 将角度转换为弧度
func degreeToRadian(deg float64) float64 {
	return deg * (math.Pi / 180.0)
}

func WGSDistance(lat1, lon1, lat2, lon2 float64) float64 {
	// 将经纬度转换为弧度
	lat1Rad := degreeToRadian(lat1)
	lon1Rad := degreeToRadian(lon1)
	lat2Rad := degreeToRadian(lat2)
	lon2Rad := degreeToRadian(lon2)

	// // 计算差值
	dLat := lat2Rad - lat1Rad
	dLon := lon2Rad - lon1Rad

	// 使用 Haversine 公式计算距离
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	

	// 距离为地球半径乘以角度
	distance := earthRadiusKm * c

	return distance
}

func EUCDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x1 - x2, 2) + math.Pow(y1- y2, 2))
}


func CalcTSP(path []int) float64 {
	v := 0.0
	for i := 1; i < len(path); i++ {
		v += DistMatrix[path[i-1]][path[i]]
	}
	return v + DistMatrix[path[0]][path[len(path)-1]]
}
