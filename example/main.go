// Package main renders an image, gif or video
package main

import (
	"github.com/bit101/bitlib/blmath"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/blfract/algos"
	"github.com/bit101/blfract/colorizors"
	"github.com/bit101/blfract/complexplane"
	"github.com/bit101/blfract/iterator"
	"github.com/bit101/blfract/warpers"
)

// revive:disable:unused-parameter

func main() {

	renderTarget := target.Video
	// renderFrame := renderFrameDuck
	renderFrame := renderFrameGrid
	// renderFrame := renderFrameKali
	// renderFrame := renderFrameJulia
	// renderFrame := renderFrameMandel
	// renderFrame := renderFrameNovaBase
	// renderFrame := renderFrameNovaRelaxed
	// renderFrame := renderFrameNovaZ

	switch renderTarget {
	case target.Image:
		w := 1000.0
		h := 1000.0
		render.Image(w, h, "out/out.png", renderFrame, 0.0)
		render.ViewImage("out/out.png")
		break

	case target.Video:
		seconds := 2
		fps := 30
		w := 400.0
		h := 400.0
		render.Frames(w, h, seconds*fps, "out/frames", renderFrame)
		render.ConvertToVideo("out/frames", "out/out.mp4", 400, 400, fps, seconds)
		render.PlayVideo("out/out.mp4")
		break
	}
}

func renderFrameDuck(context *cairo.Context, width, height, percent float64) {
	// cr := 0.20 + math.Cos(percent*blmath.Tau)*0.02
	// ci := -0.40 + math.Sin(percent*blmath.Tau)*0.02
	cr := 0.38 + percent*0.04
	ci := -0.2
	cp := complexplane.FromCenterAndSize(-0.0, 0, 4, 4)

	it := iterator.New(context, cp)
	it.FractalFunc = algos.Duck(cr, ci)
	it.ColorFunc = colorizors.Duck(1, 0.4, 0.4, 4)
	// it.WarpFunc = warpers.Ripple(width/2, blmath.LerpSin(percent, 0, height), 200, 12, 0) // -percent*blmath.Tau)
	it.Iterate(40)
}

func renderFrameGrid(context *cairo.Context, width, height, percent float64) {
	size := 10.0
	cp := complexplane.FromCenterAndSize(0.0, 0.0, size, size)

	it := iterator.New(context, cp)
	// it.FractalFunc = algos.Checker(cp, 20)
	it.FractalFunc = algos.Rings(cp, 40)
	it.ColorFunc = colorizors.GreyScale(1, 0)
	it.WarpFunc = warpers.Swirl(width/4, height/4, blmath.LerpSin(percent, width/4, width))
	// it.WarpFunc = warpers.Simplex(blmath.LerpSin(percent, 0, 1), 0.0015, 20)
	// it.WarpFunc = warpers.Fisheye(blmath.LerpSin(percent, 0, width), blmath.LerpSin(percent, 0, height), width)
	// it.WarpFunc = warpers.Ripple(width/2, height/2, 50, 5, percent*blmath.Tau)
	it.Iterate(80)
}

func renderFrameJulia(context *cairo.Context, width, height, percent float64) {
	cp := complexplane.FromCenterAndSize(0, 0, 3.2, 3.2)

	it := iterator.New(context, cp)
	it.FractalFunc = algos.Julia(-0.65, -0.37)
	it.ColorFunc = colorizors.GreyScale(0, 1)
	it.Iterate(30)
}

func renderFrameKali(context *cairo.Context, width, height, percent float64) {
	cr := -0.2
	ci := -1.0
	size := 2.6
	cp := complexplane.FromCenterAndSize(0.0, 0.0, size, size)

	it := iterator.New(context, cp)
	it.FractalFunc = algos.Kali(cr, ci, 2)
	it.ColorFunc = colorizors.GreyScale(1, 0)
	it.Iterate(80)
}

func renderFrameMandel(context *cairo.Context, width, height, percent float64) {
	cp := complexplane.FromCenterAndSize(-0.5, -0.0, 3, 3)

	it := iterator.New(context, cp)
	it.FractalFunc = algos.Mandel()
	it.ColorFunc = colorizors.GreyScale(0, 1)
	it.Iterate(40)
}

func renderFrameNovaBase(context *cairo.Context, width, height, percent float64) {
	size := 1.4
	cp := complexplane.FromCenterAndSize(0, 0, size, size)

	it := iterator.New(context, cp)
	it.FractalFunc = algos.NovaBase(4, 0.001)
	it.ColorFunc = colorizors.HSV(20, 60, 0, 1, 1.5, 0)
	it.Iterate(60)
}

func renderFrameNovaRelaxed(context *cairo.Context, width, height, percent float64) {
	size := 1.4
	cp := complexplane.FromCenterAndSize(-0.25, 0, size, size)

	it := iterator.New(context, cp)
	it.FractalFunc = algos.NovaRelaxed(1, 1, 1, 0, 4, 0.001)
	it.ColorFunc = colorizors.HSV(300, 60, 0, 1, 1.5, 0)
	it.Iterate(50)
}

func renderFrameNovaZ(context *cairo.Context, width, height, percent float64) {
	size := 1.4
	cp := complexplane.FromCenterAndSize(-0.3, 0, size, size)

	it := iterator.New(context, cp)
	it.FractalFunc = algos.NovaZ(1, 0, 4, 0.001)
	it.ColorFunc = colorizors.HSV(20, 60, 0, 1, 1.5, 0)
	it.Iterate(60)
}
