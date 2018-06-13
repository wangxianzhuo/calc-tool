package circle

import (
	"io"
	"math"

	"github.com/wangxianzhuo/calc-tool/util"
)

// Diagram ...
type Diagram struct {
	Points  PointSet `json:"point_set"`
	Optimal Point    `json:"optimal"`
}

// GetZByXAndY get z value by x & y value
func GetZByXAndY(dia Diagram, point Point) (z float64, err error) {
	isStop := false

	var y, x float64

	var x1, x2 float64
	var y1, y2 float64
	var xx, yy float64
	var l1, l2 float64 // line1, line2
	var f1, f2 float64
	var fft float64

	count := 1

	y = point.Y - dia.Optimal.Y
	x = point.X - dia.Optimal.X

	projection := Get2DProjectionFrom3D(dia, SortDefault)
	projectionList := projection.GetList()

	for _, z := range projectionList {
		if isStop {
			break
		}
		values := projection[z]
		currentZ := z

		for _, value := range values {
			x2 = x1
			y2 = y1

			x1 = value.Y - dia.Optimal.Y
			y1 = value.X - dia.Optimal.X

			if util.CompareFloat64(x1, y1) == 0 {
				continue
			}
			if util.CompareFloat64(x1, y) == 0 && util.CompareFloat64(y1, x) == 0 {
				return currentZ, nil
			}
			if util.CompareFloat64(x/y, (y2-y1)/(x2-x1)) == 0 {
				continue
			}

			if (util.CompareFloat64(y, 0)) == 0 {
				yy = (y2-y1)*x1/(x2-x1) + y1
				xx = (yy - (y1 - (y2-y1)*x1/(x2-x1))) / ((y2 - y1) / (x2 - x1))
			} else {
				xx = (y1 - (y2-y1)*x1/(x2-x1)) / (x/y - (y2-y1)/(x2-x1)) // 两点直线方程计算
				yy = (x / y) * xx                                        // point点的位置
			}
			if ((y >= 0 && xx >= 0) || (y <= 0 && xx <= 0)) && ((x >= 0 && yy >= 0) || (x <= 0 && yy <= 0)) {
				if (xx != 0) || (yy != 0) {
					if (math.Sqrt(y*y+x*x) <= math.Sqrt(xx*xx+yy*yy)) && ((xx >= x1 && xx <= x2) || (xx >= x2 && xx <= x1)) {
						l1 = math.Sqrt((y-xx)*(y-xx) + (x-yy)*(x-yy)) // z小的圈
						f1 = currentZ
						count = 0
						isStop = true
						break
					}
					if (xx >= x1 && xx <= x2) || (xx >= x2 && xx <= x1) {
						l2 = math.Sqrt((y-xx)*(y-xx) + (x-yy)*(x-yy)) // z大的圈
						f2 = currentZ
					}
				} else if (xx != 0) || (yy != 0) && ((util.CompareFloat64(xx, x1) == 0 && util.CompareFloat64(yy, y1) == 0) || (util.CompareFloat64(xx, x2) == 0 && util.CompareFloat64(yy, y2) == 0)) {
					if util.CompareFloat64(y, 0) > 0 && util.CompareFloat64(y, xx) < 0 || util.CompareFloat64(y, 0) < 0 && util.CompareFloat64(y, xx) > 0 {
						l2 = math.Sqrt((y-xx)*(y-xx) + (x-yy)*(x-yy))
						f2 = currentZ
					} else if util.CompareFloat64(y, 0) > 0 && util.CompareFloat64(y, xx) > 0 || util.CompareFloat64(y, 0) < 0 && util.CompareFloat64(y, xx) < 0 {
						l1 = math.Sqrt((y-xx)*(y-xx) + (x-yy)*(x-yy))
						f1 = currentZ
						count = 0
						isStop = true
						break
					}
				}
			}
		} // for values end

		x1 = 0
		x2 = 0
		y1 = 0
		y2 = 0
	} // for projectionList end

	if count == 0 && l2 == 0 {
		fft = f1 + (dia.Optimal.Z-f1)*l1/math.Sqrt(xx*xx+yy*yy)
	} else {
		if l1 == 0 && f1 == 0 {
			return 0, nil
		} else if l1+l2 != 0 {
			fft = f1 + (f2-f1)*l1/(l1+l2)
		}
	}
	return fft, nil
}

// GetZViaStream ...
func GetZViaStream(r io.Reader, optimal, wantedPoint Point) (z float64, err error) {
	ps, err := LoadPointSet(r)
	if err != nil {
		return 0, err
	}

	dia := Diagram{
		Points:  ps,
		Optimal: optimal,
	}

	return GetZByXAndY(dia, wantedPoint)
}
