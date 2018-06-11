package hydropower

import (
	"fmt"
	"math"
)

// WaterHead 计算水头
//
// waterLevel		水位
// tailWaterLevel	尾水位
// waterHeadLoss	水头损失
func WaterHead(waterLevel, tailWaterLevel, waterHeadLoss float64) (float64, error) {
	if waterLevel < tailWaterLevel {
		return 0, fmt.Errorf("water level [%v] is lower than tail water level [%v], can't calculate water head", waterLevel, tailWaterLevel)
	}
	return waterLevel - tailWaterLevel - waterHeadLoss, nil
}

// WaterHeadLoss 计算水头损失
//
// waterLevel				水位
// tailWaterLevel			尾水位
// activePower				机组负荷（有功功率）
// lossCoefficiency			损失系数
// HydroturbineEfficiency 	水轮机效率
func WaterHeadLoss(waterLevel, tailWaterLevel, activePower, lossCoefficiency, HydroturbineEfficiency float64) (float64, error) {
	if math.Abs(waterLevel-tailWaterLevel) < FLOAT_COMPARE_PREC {
		return 0, fmt.Errorf("water level [%v] lower than tail water level [%v], can't calculate water head loss", waterLevel, tailWaterLevel)
	}

	loss := activePower / (G * (waterLevel - tailWaterLevel) * HydroturbineEfficiency)
	loss = math.Pow(loss, 2) * lossCoefficiency
	return loss, nil
}
