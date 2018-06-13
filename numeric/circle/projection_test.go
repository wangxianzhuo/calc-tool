package circle

import (
	"testing"
)

func Test_Projection(t *testing.T) {
	ps := PointSet{
		Point{1, 2, 0},
		Point{2, 2, 0},
		Point{2, 1, 0},
		Point{1, 1, 0},
	}

	dia := Diagram{
		Points:  ps,
		Optimal: Point{1.5, 1.5, -1},
	}

	proj := Get2DProjectionFrom3D(dia, SortDefault)
	t.Log(proj)
}
