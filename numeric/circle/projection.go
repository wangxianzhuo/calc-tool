package circle

import (
	"sort"

	"github.com/wangxianzhuo/calc-tool/util"
)

// Projection is the projection of the Diagram point set
// to 2d(x,y) coordinate
type Projection map[float64][]Point

// Get2DProjectionFrom3D get the 2d projection.
// It can be sorted by sort function, default sort function is raw order.
func Get2DProjectionFrom3D(dia Diagram, sort func(Projection, Point) Projection) Projection {
	projection := make(Projection)

	for _, p := range dia.Points {
		if projection[p.Z] == nil {
			projection[p.Z] = make([]Point, 0)
		}
		projection[p.Z] = append(projection[p.Z], p)
	}

	if sort == nil {
		return projection
	}
	return sort(projection, dia.Optimal)
}

// SortDefault is order by x
func SortDefault(input Projection, optiaml Point) (output Projection) {
	output = make(Projection)

	for k, values := range input {
		positiveList := make(PointSet, 0, len(values))
		negativeList := make(PointSet, 0, len(values))

		for _, v := range values {
			switch util.CompareFloat64(v.Y, optiaml.Y) {
			case -1:
				negativeList = append(negativeList, v)
			default:
				positiveList = append(positiveList, v)
			}
		}

		sort.Sort(positiveList)
		sort.Sort(sort.Reverse(negativeList))

		tmp := make(PointSet, len(values))
		copy(tmp, positiveList)
		copy(tmp[len(positiveList):], negativeList)
		output[k] = tmp
	}

	return output
}

// GetList get list order by porjection's key
func (p Projection) GetList() []float64 {
	list := make([]float64, 0)
	for k := range p {
		list = append(list, k)
	}

	sort.Sort(sort.Reverse(sort.Float64Slice(list)))
	return list
}
