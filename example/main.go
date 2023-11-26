// Package main renders an image, gif or video
package main

import (
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

	renderTarget := target.Image
	// renderFrame := renderFrameDuck
	// renderFrame := renderFrameGrid
	// renderFrame := renderFrameKali
	renderFrame := renderFrameJulia
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
		render.Frames(400, 400, seconds*fps, "out/frames", renderFrame)
		render.ConvertToVideo("out/frames", "out/out.mp4", 400, 400, fps, seconds)
		render.PlayVideo("out/out.mp4")
		break
	}
}

func renderFrameDuck(context *cairo.Context, width, height, percent float64) {
	cr := 0.4
	ci := -0.2
	cp := complexplane.FromCenterAndSize(-0.0, 0, 4, 4)

	it := iterator.New(context, cp)
	it.FractalFunc = algos.Duck(cr, ci)
	it.ColorFunc = colorizors.Duck(1, 0.4, 0.4, 2)
	it.Iterate(40)
}

func renderFrameGrid(context *cairo.Context, width, height, percent float64) {
	size := 10.0
	cp := complexplane.FromCenterAndSize(0.0, 0.0, size, size)

	it := iterator.New(context, cp)
	it.FractalFunc = algos.Grid(cp, 30)
	it.ColorFunc = colorizors.GreyScale(1, 0)
	it.WarpFunc = warpers.Fisheye(width/2, height/2, width/2)
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
