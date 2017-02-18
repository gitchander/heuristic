package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/gitchander/cairo"
	"github.com/gitchander/cairo/imutil"
	"github.com/gitchander/heuristic/math/hexm"
)

const or = hexm.Angled

func main() {

	ims := []InfoMakePNG{
		InfoMakePNG{"hexboard1.png", drawHexBoard1},
		InfoMakePNG{"hexboard2.png", drawHexBoard2},
		InfoMakePNG{"hexboard3.png", drawHexBoard3},
		InfoMakePNG{"orientation-flat.png", drawOrientationFlat},
		InfoMakePNG{"orientation-angled.png", drawOrientationAngled},
	}

	for _, im := range ims {
		if err := makeImages(im); err != nil {
			log.Fatal(err.Error())
		}
	}
}

type InfoMakePNG struct {
	Filename string
	Draw     DrawFunc
}

type DrawFunc func(*cairo.Canvas) error

func makeImages(im InfoMakePNG) error {

	surface, err := cairo.NewSurface(cairo.FORMAT_ARGB32, 512, 512)
	if err != nil {
		return err
	}
	defer surface.Destroy()

	canvas, err := cairo.NewCanvas(surface)
	if err != nil {
		return err
	}
	defer canvas.Destroy()

	if err = im.Draw(canvas); err != nil {
		return err
	}

	if err = surface.WriteToPNG(im.Filename); err != nil {
		return err
	}

	return nil
}

func drawHexBoard1(c *cairo.Canvas) error {

	hm := hexm.NewMatrix(hexm.Coord{5, 5, 5})

	var (
		surface = c.GetTarget()
		nX      = surface.GetWidth()
		nY      = surface.GetHeight()
	)

	imutil.CanvasFillColor(c, color.White)

	size := hm.Size()
	radius := calcRadius(nX, nY, size)

	m := cairo.NewMatrix()
	m.InitIdendity()
	m.InitTranslate(float64(nX)*0.5, float64(nY)*0.5)
	m.Scale(radius, radius)
	c.SetMatrix(m)

	poly := hexm.HexVertexes(or)

	cs := palettes[4]

	var extents cairo.TextExtents

	c.SetLineWidth(0.03)

	index := 0
	for I := hexm.NewIterator(hm); I.HasValue(); I.Next() {

		coord := I.Coord()
		v := hexm.CoordToVector(or, coord)

		p := poly[len(poly)-1]
		p = v.Add(p)
		c.MoveTo(p.X, p.Y)

		for _, p := range poly {
			p = v.Add(p)
			c.LineTo(p.X, p.Y)
		}

		var cl RGB

		switch {
		case (coord.X == 0) && (coord.Z > 0):
			cl = cs[0]

		case (coord.Y == 0) && (coord.X > 0):
			cl = cs[1]

		case (coord.Z == 0) && (coord.Y > 0):
			cl = cs[2]

		default:
			cl = RGB{255, 255, 255}
		}

		setSourceColor(c, cl)
		//c.SetSourceRGB(0, 0, 1)
		c.FillPreserve()

		c.SetSourceRGB(0, 0, 0)
		c.Stroke()

		// draw text
		{
			text := fmt.Sprintf("(%d,%d,%d)", coord.X, coord.Y, coord.Z)
			//text := fmt.Sprintf("%d", index)

			c.SetFontSize(0.35)

			c.TextExtents(text, &extents)
			x := v.X - (extents.Width/2 + extents.BearingX)
			y := v.Y - (extents.Height/2 + extents.BearingY)

			c.SelectFontFace("Sans", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_NORMAL)
			c.MoveTo(x, y)
			c.ShowText(text)
		}

		index++
	}

	return nil
}

func drawHexBoard2(c *cairo.Canvas) error {

	hm := hexm.NewMatrix(hexm.Coord{5, 5, 5})

	var (
		surface = c.GetTarget()
		nX      = surface.GetWidth()
		nY      = surface.GetHeight()
	)

	imutil.CanvasFillColor(c, color.White)

	radius := calcRadius(nX, nY, hm.Size())

	m := cairo.NewMatrix()
	m.InitIdendity()
	m.InitTranslate(float64(nX)*0.5, float64(nY)*0.5)
	m.Scale(radius, radius)
	c.SetMatrix(m)

	poly := hexm.HexVertexes(or)

	c.SetLineWidth(0.05)
	c.SetSourceRGB(0.7, 0.7, 0.7)

	for I := hexm.NewIterator(hm); I.HasValue(); I.Next() {

		coord := I.Coord()
		v := hexm.CoordToVector(or, coord)

		p := poly[len(poly)-1]
		p = v.Add(p)
		c.MoveTo(p.X, p.Y)

		for _, p := range poly {
			p = v.Add(p)
			c.LineTo(p.X, p.Y)
		}

		//c.SetSourceRGB(1, 1, 1)
		//c.FillPreserve()

		c.SetSourceRGB(0.7, 0.7, 0.7)
		c.Stroke()
	}

	var (
		cs [3]hexm.Coord
		vs [3]hexm.Vector
	)

	c.SetLineWidth(0.06)
	c.SetSourceRGB(0.5, 0, 0)

	size := hm.Size()

	for x := 0; x < size.X; x++ {

		cs[0] = hexm.Coord{x, size.Y - 1, 0}
		cs[1] = hexm.Coord{x, 0, 0}
		cs[2] = hexm.Coord{x, 0, size.Z - 1}

		vs[0] = hexm.CoordToVector(or, cs[0])
		vs[1] = hexm.CoordToVector(or, cs[1])
		vs[2] = hexm.CoordToVector(or, cs[2])

		c.MoveTo(vs[0].X, vs[0].Y)
		c.LineTo(vs[1].X, vs[1].Y)
		c.LineTo(vs[2].X, vs[2].Y)

		c.Stroke()
	}

	for y := 0; y < size.Y; y++ {

		cs[0] = hexm.Coord{0, y, size.Z - 1}
		cs[1] = hexm.Coord{0, y, 0}
		cs[2] = hexm.Coord{size.X - 1, y, 0}

		vs[0] = hexm.CoordToVector(or, cs[0])
		vs[1] = hexm.CoordToVector(or, cs[1])
		vs[2] = hexm.CoordToVector(or, cs[2])

		c.MoveTo(vs[0].X, vs[0].Y)
		c.LineTo(vs[1].X, vs[1].Y)
		c.LineTo(vs[2].X, vs[2].Y)

		c.Stroke()
	}

	for z := 0; z < size.Z; z++ {

		cs[0] = hexm.Coord{size.X - 1, 0, z}
		cs[1] = hexm.Coord{0, 0, z}
		cs[2] = hexm.Coord{0, size.Y - 1, z}

		vs[0] = hexm.CoordToVector(or, cs[0])
		vs[1] = hexm.CoordToVector(or, cs[1])
		vs[2] = hexm.CoordToVector(or, cs[2])

		c.MoveTo(vs[0].X, vs[0].Y)
		c.LineTo(vs[1].X, vs[1].Y)
		c.LineTo(vs[2].X, vs[2].Y)

		c.Stroke()
	}

	return nil
}

func drawHexBoard3(c *cairo.Canvas) error {

	hm := hexm.NewMatrix(hexm.Coord{5, 5, 5})

	var (
		surface = c.GetTarget()
		nX      = surface.GetWidth()
		nY      = surface.GetHeight()
	)

	imutil.CanvasFillColor(c, color.White)

	radius := calcRadius(nX, nY, hm.Size())

	m := cairo.NewMatrix()
	m.InitIdendity()
	m.InitTranslate(float64(nX)*0.5, float64(nY)*0.5)
	m.Scale(radius, radius)
	c.SetMatrix(m)

	c.SetLineWidth(0.03)

	poly := hexm.HexVertexes(or)

	cs := palettes[3]

	for I := hexm.NewIterator(hm); I.HasValue(); I.Next() {

		coord := I.Coord()
		v := hexm.CoordToVector(or, coord)

		p := poly[len(poly)-1]
		p = v.Add(p)
		c.MoveTo(p.X, p.Y)

		for _, p := range poly {
			p = v.Add(p)
			c.LineTo(p.X, p.Y)
		}

		colorIndex := (coord.X + coord.Y + coord.Z) % len(cs)
		setSourceColor(c, cs[colorIndex])
		//c.SetSourceRGB(0, 0, 1)
		c.FillPreserve()

		c.SetSourceRGB(0, 0, 0)
		c.Stroke()
	}

	return nil
}

func calcRadius(nX, nY int, size hexm.Coord) float64 {
	return float64(min(nX, nY)) / (float64(max(size.X, size.Y, size.Z)) * 3.5)
}

func drawOrientationFlat(c *cairo.Canvas) error {
	return drawOrientation(c, hexm.Flat)
}

func drawOrientationAngled(c *cairo.Canvas) error {
	return drawOrientation(c, hexm.Angled)
}

func drawOrientation(c *cairo.Canvas, or hexm.Orientation) error {
	var (
		surface = c.GetTarget()
		nX      = surface.GetWidth()
		nY      = surface.GetHeight()
	)

	vs := hexm.HexVertexes(or)

	radius := 0.8 * (float64(max(nX, nY)) / 2)

	imutil.CanvasFillColor(c, color.White)

	m := cairo.NewMatrix()
	m.InitIdendity()
	m.InitTranslate(float64(nX)*0.5, float64(nY)*0.5)
	m.Scale(radius, radius)
	c.SetMatrix(m)

	c.SetLineWidth(0.03)
	c.SetLineJoin(cairo.LINE_JOIN_ROUND)
	c.SetLineCap(cairo.LINE_CAP_ROUND)

	if len(vs) > 0 {

		v := vs[len(vs)-1]
		c.MoveTo(v.X, v.Y)

		for _, v := range vs {
			c.LineTo(v.X, v.Y)
		}

		c.SetSourceRGB(0, 0, 0)
		c.Stroke()
	}

	return nil
}
