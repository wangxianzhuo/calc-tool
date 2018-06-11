package hydropower

import (
	"fmt"
	"math"
)

// G 重力加速度
// HYDROTURBINE_M 水轮机效率近似值
// FLOAT_COMPARE_PREC 浮点数比较精度
const (
	G                  = 9.81
	FLOAT_COMPARE_PREC = 0.00001
)

// UnitUsedFlow 计算发电流量
//
// motorEfficiency 	电机效率
// waterHead		水头
// efficiency		机组效率
// activePower		机组负荷（有功功率）
func UnitUsedFlow(motorEfficiency, waterHead, efficiency, activePower float64) (float64, error) {
	if efficiency > 1 {
		efficiency /= 100
	}

	if waterHead*efficiency == 0 {
		return 0, fmt.Errorf("water head [%v] or efficiency [%v] is 0, can't calculate unit uesed flow", waterHead, efficiency)
	}
	return (activePower / motorEfficiency) / (G * (waterHead * efficiency)), nil
}

// UnitSpecialWorkPointInflow 计算特定功率下的流量
//
// waterHead		水头
// motorEfficiency	电机效率
// activePower		机组负荷（有功功率）
// openDegree 		导叶开度 (单位 %)
// k 				系数
func UnitSpecialWorkPointInflow(waterHead, motorEfficiency, activePower, openDegree, k float64) (float64, error) {
	if motorEfficiency > 1 {
		motorEfficiency /= 100
	}

	if waterHead*motorEfficiency*k == 0 {
		return 0, fmt.Errorf("water head [%v] or efficiency [%v] or k [%v] is 0, can't calculate unit uesed flow", waterHead, motorEfficiency, k)
	}
	return (activePower / (G * waterHead * motorEfficiency * k)) * (openDegree / 50), nil
}

// KValue 计算导叶开度相关流量的K值
func KValue(p1, p2, p3, p4, p5, waterHead float64) (k float64) {
	k = (p1 - ((waterHead-p2)/p3)*p4) / p5
	return
}

// SluiceFlow 计算闸口出流
//
// k1 	闸口系数
// k2 	闸口系数
// e 	闸口开度
// l 	闸门宽度
// h 	水头
func SluiceFlow(k1, k2, e, l, h float64) float64 {
	if e > h {
		return k2 * math.Pow(h, 1.5)
	}
	return k1 * e * math.Pow(2*G*h, 0.5)
}

// SluiceOverflow 计算闸门溢流流量
//
// waterLevel		液位
// sluiceHeight		闸门高度
// sluiceOpenDegree	闸门开度
// k				溢流系数，一般为0.67
func SluiceOverflow(waterLevel, sluiceHeight, sluiceOpenDegree, k float64) float64 {
	if waterLevel > (sluiceHeight + sluiceOpenDegree) {
		return k * math.Pow((waterLevel-(sluiceHeight+sluiceOpenDegree)), 1.5)
	}
	return 0
}

// DamOverflow 计算大坝溢流流量
//
// waterLevel	液位
// damHeight	大坝高度
// k			溢流系数，一般为0.67
func DamOverflow(waterLevel, damHeight, k float64) float64 {
	if waterLevel > damHeight {
		return k * math.Pow((waterLevel-damHeight), 1.5)
	}
	return 0
}
