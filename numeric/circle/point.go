package circle

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Point 3d coordinate point
type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

func (p Point) String() string {
	return fmt.Sprintf("(%f, %f, %f)", p.X, p.Y, p.Z)
}

// PointSet 3d coordinate point set
type PointSet []Point

func (ps PointSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, point := range ps {
		buf.WriteString(point.String())
		if i < len(ps)-1 {
			buf.WriteString(",")
		}
	}
	buf.WriteString("]")
	return buf.String()
}

// LoadPointSet load point set from byte stream
func LoadPointSet(r io.Reader) (PointSet, error) {
	var pointSet PointSet
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		elements := strings.Fields(line)
		if len(elements) < 3 {
			return nil, fmt.Errorf("point: %s do not have enough data", line)
		}

		var p Point
		x, err := strconv.ParseFloat(elements[1], 64)
		if err != nil {
			return nil, err
		}
		y, err := strconv.ParseFloat(elements[2], 64)
		if err != nil {
			return nil, err
		}
		z, err := strconv.ParseFloat(elements[0], 64)
		if err != nil {
			return nil, err
		}
		p.X = x
		p.Y = y
		p.Z = z

		pointSet = append(pointSet, p)
	}

	return pointSet, nil
}

func (ps PointSet) Len() int {
	return len(ps)
}

func (ps PointSet) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ps PointSet) Less(i, j int) bool {
	return ps[i].X < ps[j].X
}
