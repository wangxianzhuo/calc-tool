package circle

import "testing"

func Test_Point(t *testing.T) {
	var ps PointSet
	t.Logf("%v", ps)

	ps = append(ps, Point{1, 2, 3})
	t.Logf("%v", ps)
}
