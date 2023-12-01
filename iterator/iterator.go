// Package iterator holds the data to iterate over a complex plane
package iterator

import (
	"fmt"

	"github.com/bit101/bitlib/blmath"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blfract/algos"
	"github.com/bit101/blfract/colorizors"
	"github.com/bit101/blfract/complexplane"
	"github.com/bit101/blfract/warpers"
)

// Iterator holds the data to iterate over a complex plane
type Iterator struct {
	Context     *cairo.Context
	Plane       *complexplane.ComplexPlane
	FractalFunc algos.Algo
	ColorFunc   colorizors.Colorizor
	WarpFunc    warpers.Warper
}

// New creates a new Iterator
func New(context *cairo.Context, plane *complexplane.ComplexPlane) *Iterator {
	return &Iterator{
		context,
		plane,
		algos.Mandel(),
		colorizors.GreyScale(0, 1),
		warpers.Default(),
	}
}

// Iterate iterates over a complex plane to render it with a given algorithm and colorizor.
func (i *Iterator) Iterate(iter int) {
	width, height := i.Context.Width, i.Context.Height
	fIter := float64(iter)
	for y := 0.0; y < height; y++ {
		for x := 0.0; x < width; x++ {
			wx, wy := i.WarpFunc(x, y)
			re := blmath.Map(wx, 0, width, i.Plane.RealMin, i.Plane.RealMax)
			im := blmath.Map(wy, 0, height, i.Plane.ImagMin, i.Plane.ImagMax)
			m := i.FractalFunc(re, im, fIter)
			col := i.ColorFunc(m)
			i.Context.SetSourceColor(col)
			i.Context.FillRectangle(x, y, 1, 1)
		}
	}
}

func (i *Iterator) PrintParams(params ...any) {
	i.Context.Save()
	i.Context.SetFontSize(14)
	x := 11.0
	y := 16.0
	for j := 0; j < len(params); j += 2 {
		s := fmt.Sprintf("%v: %v", params[j], params[j+1])
		i.Context.SetSourceBlack()
		i.Context.FillText(fmt.Sprintf(s), x+1, y+1)
		i.Context.SetSourceWhite()
		i.Context.FillText(fmt.Sprintf(s), x, y)
		y += 16.0
	}
	i.Context.Restore()
}
