package hydropower

import (
	"encoding/binary"
	"math"
	"strconv"
)

// CubicSplineCurve 三次样条曲线查值
func CubicSplineCurve() (float64, error) {
	return 0, nil
}

// Float64Compare 比较浮点数大小
func Float64Compare(a, b float64) int {
	switch {
	case math.Abs(a-b) < FLOAT_COMPARE_PREC || a-b == 0:
		return 0
	case a-b > 0:
		return 1
	default:
		return -1
	}
}

// JoinStringMaps 合并两个map，返回一个新的map，其中map的key是string类型，value是interface{}类型
func JoinStringMaps(a, b map[string]interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	for k, v := range a {
		switch value := v.(type) {
		case float64:
			m[k] = strconv.FormatFloat(value, 'f', 5, 64)
		case int64:
			m[k] = strconv.FormatInt(value, 10)
		case bool:
			m[k] = strconv.FormatBool(value)
		case string:
			m[k] = value
		default:
			continue
		}
	}
	for k, v := range b {
		switch value := v.(type) {
		case float64:
			m[k] = strconv.FormatFloat(value, 'f', 5, 64)
		case int64:
			m[k] = strconv.FormatInt(value, 10)
		case bool:
			m[k] = strconv.FormatBool(value)
		case string:
			m[k] = value
		default:
			continue
		}
	}
	return m
}

// BytesToFloat32 将bytes转换成IEEE754标准的四字节浮点数
//
// byteOrder 是binary.BigEndian（大端法）或者binary.LittleEndian（小端法）
// order 是字节顺序，例如order = {3,2,1,0}表示根据[]byte{input[3],input[2],input[1],input[0]}来得到f
// encode 是编码模式：默认0（ieee754浮点数）；1（浮点数）
func BytesToFloat32(input []byte, byteOrder binary.ByteOrder, order []byte, encode uint8) (f float64, err error) {
	b := []byte{input[order[0]], input[order[1]], input[order[2]], input[order[3]]}
	bits := byteOrder.Uint32(b)
	switch encode {
	case 1:
		f = float64(bits)
	default:
		f = float64(math.Float32frombits(bits))
	}
	return
}
