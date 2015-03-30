package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/gitchander/go-lang/cairo"
	"github.com/gitchander/heuristic/math/cubspl"
)

func main() {

	var (
		width  = 800
		height = 400
	)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ps := makePoints(r, 10, width, height)

	spline, err := cubspl.NewCubicSpline(ps)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, width, height)
	defer surface.Destroy()

	canvas, _ := cairo.NewCanvas(surface)
	defer canvas.Destroy()

	canvas.SetSourceRGB(1, 1, 1)
	canvas.Paint()

	canvas.SetSourceRGB(0.7, 0, 0)

	minX := ps[0].X
	maxX := ps[len(ps)-1].X
	dX := (maxX - minX) / 200

	x := minX
	y := spline.Interpolate(x)
	canvas.MoveTo(x, y)
	x += dX

	for x < maxX {
		y = spline.Interpolate(x)
		canvas.LineTo(x, y)
		x += dX
	}

	canvas.Stroke()

	for _, p := range ps {

		canvas.Arc(p.X, p.Y, 4, 0, 2*math.Pi)

		canvas.SetSourceRGB(1, 1, 1)
		canvas.FillPreserve()

		canvas.SetSourceRGB(0, 0.5, 0)
		canvas.Stroke()
	}

	surface.WriteToPNG("test.png")
}

func makePoints(r *rand.Rand, n int, width, height int) []cubspl.Point {

	ps := make([]cubspl.Point, n)

	var (
		minX = float64(width) * 0.1
		maxX = float64(width) * 0.9
	)
	var dX = (maxX - minX) / float64(n-1)

	var (
		minY = float64(height) * 0.2
		maxY = float64(height) * 0.8
	)

	for i, _ := range ps {

		x := minX + dX*float64(i)
		y := minY + r.Float64()*(maxY-minY)

		ps[i].Set(x, y)
	}

	return ps
}
