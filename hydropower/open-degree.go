package hydropower

import "math"

// StraightSluiceOpenDistance 计算直板闸门开度
//
// 支持两种方案：
// 1. basic < measuredValue（measuredValue随着闸门开度增大而增大）
// 2. basic > measuredValue（measuredValue随着闸门开度增大而减小）
//
// basic 基准值，即测量零点的测量值
// measuredValue 测量值
func StraightSluiceOpenDistance(basic, measuredValue float64) float64 {
	if basic < measuredValue {
		return measuredValue - basic
	}
	return basic - measuredValue
}

// CurveSluiceOpenDistance 计算弧形门开度
//
// a 				基准值，即测量零点的测量值
// b				弧形门转轴到测点的距离
// c				弧形门半径
// measuredValue	测量值
func CurveSluiceOpenDistance(a, b, c, measuredValue float64) float64 {
	e := c * math.Atan((a-measuredValue)/b)
	if e < 0 {
		return 0
	}
	return e
}
