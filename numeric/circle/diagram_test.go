package circle

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"testing"
	"time"
)

func Test_Diagram(t *testing.T) {
	f, err := os.OpenFile("out", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	ps, err := LoadPointSet(f)
	if err != nil {
		t.Fatal(err)
	}

	dia := Diagram{
		Points: ps,
		Optimal: Point{
			X: 0,
			Y: 0,
			Z: 11,
		},
	}

	z, err := GetZByX(dia, Point{X: 1, Y: 2.5})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("z: %v", z)
	t.Logf("radius: %v", math.Sqrt(1+2.5*2.5))
}

func Test_Circle(t *testing.T) {
	// t.Log(generateCircle(1, 0, 0.1))
	// t.Log(generateCircleWithRandom(1, 0, 0.1))

	out, err := os.Create("out")
	if err != nil {
		t.Fatal(err)
	}
	defer out.Close()

	// printProjection(generateProjection(1, 10, 0.1), out)
	// printProjection(generateProjection(2, 9, 0.3), out)
	// printProjection(generateProjection(3, 8, 0.2), out)
	// printProjection(generateProjection(4, 7, 0.15), out)
	printProjection(generateProjection(1, 10, 0.1), out)
	printProjection(generateProjection(2, 9, 0.3), out)
	printProjection(generateProjection(3, 8, 0.2), out)
	printProjection(generateProjection(4, 7, 0.15), out)
}

func generateProjectionWithRandom(radius, z, step float64) Projection {
	ps, _ := generateCircleWithRandom(radius, z, step)
	dia := Diagram{Points: ps, Optimal: Point{0, 0, 0}}
	proj := Get2DProjectionFrom3D(dia, SortDefault)
	proj[z] = append(proj[z], proj[z][0])
	return proj
}

func generateProjection(radius, z, step float64) Projection {
	ps, _ := generateCircle(radius, z, step)
	dia := Diagram{Points: ps, Optimal: Point{0, 0, 0}}
	proj := Get2DProjectionFrom3D(dia, SortDefault)
	proj[z] = append(proj[z], proj[z][0])
	return proj
}

func printProjection(proj Projection, w io.Writer) {
	var buf bytes.Buffer
	for _, points := range proj {
		for _, p := range points {
			buf.WriteString(fmt.Sprintf("%v %v %v\n", p.Z, p.X, p.Y))
		}
	}
	fmt.Fprint(w, buf.String())
}

func generateCircle(radius, z, step float64) (PointSet, int) {
	circle := make(PointSet, 0)
	for x := step; x <= radius; x += step {
		y := math.Sqrt(radius*radius - float64(x*x))
		circle = append(circle, Point{
			X: float64(x),
			Y: y,
			Z: z,
		})
		circle = append(circle, Point{
			X: float64(x),
			Y: -y,
			Z: z,
		})
		circle = append(circle, Point{
			X: -float64(x),
			Y: y,
			Z: z,
		})
		circle = append(circle, Point{
			X: -float64(x),
			Y: -y,
			Z: z,
		})
	}
	return circle, len(circle)
}

func generateCircleWithRandom(radius, z, step float64) (PointSet, int) {
	circle := make(PointSet, 0)
	for x := step; x <= radius; x += step {
		y := math.Sqrt(radius*radius - float64(x*x))
		rand.Seed(time.Now().UnixNano())
		a1 := rand.Float64()
		a2 := rand.Float64()
		a3 := rand.Float64()
		a4 := rand.Float64()
		circle = append(circle, Point{
			X: float64(x),
			Y: y + a1,
			Z: z,
		})
		circle = append(circle, Point{
			X: float64(x),
			Y: -(y + a2),
			Z: z,
		})
		circle = append(circle, Point{
			X: -float64(x),
			Y: y + a3,
			Z: z,
		})
		circle = append(circle, Point{
			X: -float64(x),
			Y: -y - a4,
			Z: z,
		})
	}
	return circle, len(circle)
}
