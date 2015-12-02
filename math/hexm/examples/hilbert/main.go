package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/hilbert"

	"github.com/gitchander/go-lang/cairo"
	"github.com/gitchander/go-lang/cairo/color"
)

type Size struct {
	Width, Height int
}

func HilbertCurve(c *cairo.Canvas, n int, size Size) error {

	s, err := hilbert.New(n)
	if err != nil {
		return err
	}

	var (
		x, y   int
		fX, fY float64
	)

	var (
		dX = float64(size.Width) / float64(s.N)
		dY = float64(size.Height) / float64(s.N)
	)

	c.SetLineWidth(0.2 * ((dX + dY) / 2))
	c.SetLineCap(cairo.LINE_CAP_ROUND)
	c.SetLineJoin(cairo.LINE_JOIN_ROUND)

	m := cairo.NewMatrix()
	m.InitIdendity()
	m.Scale(dX, dY)
	m.Translate(0.5, 0.5)

	if nn := s.N * s.N; nn > 0 {

		x, y, _ = s.Map(0)
		fX, fY = m.TransformPoint(float64(x), float64(y))
		c.MoveTo(fX, fY)

		for i := 1; i < nn; i++ {

			x, y, _ = s.Map(i)
			fX, fY = m.TransformPoint(float64(x), float64(y))
			c.LineTo(fX, fY)
		}
	}

	c.Stroke()

	return nil
}

func drawCurve(c *cairo.Canvas, n int, size Size) error {

	c.SetSourceRGB(0, 0, 0)
	if err := HilbertCurve(c, n, size); err != nil {
		return err
	}

	return nil
}

func drawDoubleCurve(c *cairo.Canvas, n int, size Size) error {

	c.SetSourceRGB(0.2, 0, 0)
	if err := HilbertCurve(c, n, size); err != nil {
		return err
	}

	c.SetSourceRGB(0.8, 0, 0)
	if err := HilbertCurve(c, n*2, size); err != nil {
		return err
	}

	return nil
}

func makeHilbertPNG(fileName string, n int, size Size) error {

	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, size.Width, size.Height)
	defer surface.Destroy()

	canvas, err := cairo.NewCanvas(surface)
	if err != nil {
		return err
	}
	defer canvas.Destroy()

	color.CanvasFillRGB(canvas, color.NewRGB(1.0, 1.0, 1.0))

	err = drawCurve(canvas, n, size)
	//err = drawDoubleCurve(canvas, n, size)
	if err != nil {
		return err
	}

	if err = surface.WriteToPNG(fileName); err != nil {
		return err
	}

	return nil
}

func makeDir(dir string) error {

	fi, err := os.Stat(dir)
	if err != nil {
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return err
		}
	} else {
		if !fi.IsDir() {
			return errors.New("file is not dir")
		}
	}
	return nil
}

func makeFiles() error {

	dir := "./curves"
	size := Size{512, 512}

	if err := makeDir(dir); err != nil {
		return err
	}

	p := 2
	for i := 0; i < 9; i++ {
		fileName := filepath.Join(dir, fmt.Sprintf("hilbert_curve_%d.png", p))
		if err := makeHilbertPNG(fileName, p, size); err != nil {
			return err
		}
		p *= 2
	}

	return nil
}

func main() {
	if err := makeFiles(); err != nil {
		fmt.Println(err.Error())
	}
}
