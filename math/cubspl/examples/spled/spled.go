package main

import (
	"math"

	"github.com/gitchander/go-lang/cairo"
	"github.com/gitchander/heuristic/math/cubspl"
)

type SplineEditor struct {
	rect   Rectangle
	ps     []cubspl.Point
	index  int
	spline *cubspl.CubicSpline
}

func NewSplineEditor(r Rectangle) (*SplineEditor, error) {

	ps := []cubspl.Point{
		cubspl.Point{r.Min.X, r.Min.Y},
		cubspl.Point{r.Max.X, r.Max.Y},
	}

	cs, err := cubspl.NewCubicSpline(ps)
	if err != nil {
		return nil, err
	}

	se := &SplineEditor{
		rect:   r,
		ps:     ps,
		index:  -1,
		spline: cs,
	}

	return se, nil
}

func (this *SplineEditor) Close() error {

	return nil
}

func (this *SplineEditor) Resize(width, height int) {

}

func (this *SplineEditor) KeyPress(key uint) bool {

	return false
}

func (this *SplineEditor) ButtonPress(x, y int) bool {

	q := cubspl.Point{float64(x), float64(y)}

	this.index = -1

	for i, p := range this.ps {
		if vectorsNear(p, q) {
			this.index = i
			break
		}
	}

	if this.index == -1 {

		p := cubspl.Point{
			X: q.X,
			Y: this.spline.Interpolate(q.X),
		}

		p = this.rect.NormPoint(p)

		if vectorsNear(q, p) {
			if j := this.appendPoint(p); j != -1 {
				this.index = j
				return true
			}
		}
	}

	return false
}

func (this *SplineEditor) appendPoint(p cubspl.Point) (index int) {

	ps := this.ps
	n := len(ps)
	j := -1

	for i := 1; i < n; i++ {
		if (ps[i-1].X < p.X) && (p.X < ps[i].X) {
			j = i
			break
		}
	}

	if j != -1 {

		this.ps = insertByIndex(this.ps, j, p)
		spline, err := cubspl.NewCubicSpline(this.ps)
		if err != nil {
			return -1
		}

		this.spline = spline

		return j
	}

	return -1
}

func (this *SplineEditor) ButtonRelease() bool {

	this.index = -1

	return false
}

func (this *SplineEditor) ButtonMove(x, y int) bool {

	if this.index != -1 {

		p := cubspl.Point{float64(x), float64(y)}
		p = this.rect.NormPoint(p)

		// norm X
		{
			if this.index == 0 {
				p.X = this.rect.Min.X
			}

			if this.index == len(this.ps)-1 {
				p.X = this.rect.Max.X
			}
		}

		// remove for nearest
		{
			fRemove := false

			if this.index > 0 {
				prev, _ := this.spline.GetPointByIndex(this.index - 1)
				if vectorsNear(p, prev) {
					fRemove = true
				}
			}

			if this.index < len(this.ps) {
				next, _ := this.spline.GetPointByIndex(this.index + 1)
				if vectorsNear(p, next) {
					fRemove = true
				}
			}

			if fRemove {
				_, this.ps = removeByIndex(this.ps, this.index)
				this.spline, _ = cubspl.NewCubicSpline(this.ps)
				this.spline.Recalculate()
				this.index = -1
				return true
			}
		}

		err := this.spline.SetPointByIndex(this.index, p)
		if err != nil {
			return false
		}

		this.spline.Recalculate()
		this.ps[this.index] = p

		return true
	}

	return false
}

func (this *SplineEditor) Draw(c *cairo.Canvas) {

	drawBorder(c, &(this.rect))

	drawSpline(c, &(this.rect), this.spline)
	//drawSplineFill(c, &(this.rect), this.spline)

	drawPoints(c, this.ps)
}

func (this *SplineEditor) DrawCairoNative(ptr uintptr) {

	c, _ := cairo.NewCanvasNative(ptr)
	defer c.Destroy()

	this.Draw(c)
}

func (this *SplineEditor) Name() string {
	return "SplineEditor"
}

//-------------------------------------------------------------------
func drawBorder(c *cairo.Canvas, r *Rectangle) {

	c.SetSourceRGBA(0, 0, 0, 0.3)
	c.SetLineWidth(1)
	c.Rectangle(r.Min.X, r.Min.Y, r.Width(), r.Height())
	c.Stroke()
}

func drawSpline(c *cairo.Canvas, r *Rectangle, spline *cubspl.CubicSpline) {

	c.SetSourceRGB(0.6, 0, 0)
	c.SetLineWidth(3)

	dX := 2.0
	var p cubspl.Point

	p.X = r.Min.X
	p.Y = spline.Interpolate(p.X)
	p = r.NormPoint(p)
	c.MoveTo(p.X, p.Y)
	p.X += dX

	for p.X < r.Max.X {

		p.Y = spline.Interpolate(p.X)
		p = r.NormPoint(p)
		c.LineTo(p.X, p.Y)
		p.X += dX
	}

	p.X = r.Max.X
	p.Y = spline.Interpolate(p.X)
	p = r.NormPoint(p)
	c.LineTo(p.X, p.Y)

	c.Stroke()
}

func drawSplineFill(c *cairo.Canvas, r *Rectangle, spline *cubspl.CubicSpline) {

	dX := 2.0
	var p cubspl.Point

	p.X = r.Min.X
	p.Y = spline.Interpolate(p.X)
	p = r.NormPoint(p)
	c.MoveTo(p.X, p.Y)
	p.X += dX

	for p.X < r.Max.X {

		p.Y = spline.Interpolate(p.X)
		p = r.NormPoint(p)
		c.LineTo(p.X, p.Y)
		p.X += dX
	}

	p.X = r.Max.X
	p.Y = spline.Interpolate(p.X)
	p = r.NormPoint(p)
	c.LineTo(p.X, p.Y)

	c.LineTo(r.Max.X, r.Max.Y)
	c.LineTo(r.Min.X, r.Max.Y)

	c.SetSourceRGB(0.6, 0, 0)
	c.Fill()
}

func drawPoints(c *cairo.Canvas, ps []cubspl.Point) {

	c.SetLineWidth(2)

	for _, p := range ps {

		c.Arc(p.X, p.Y, 4, 0, 2*math.Pi)

		c.SetSourceRGB(1, 1, 1)
		c.FillPreserve()

		c.SetSourceRGB(0, 0.5, 0)
		c.Stroke()
	}
}
