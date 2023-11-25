// Package main renders an image, gif or video
package main

import (
	"math"

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

func main() {

	renderTarget := target.Video
	renderFrame := renderFrameDuck

	switch renderTarget {
	case target.Image:
		render.Image(600, 600, "out/out.png", renderFrame, 0.0)
		render.ViewImage("out/out.png")
		break

	case target.Video:
		seconds := 2
		fps := 30
		render.Frames(400, 400, seconds*fps, "out/frames", renderFrame)
		render.ConvertToVideo("out/frames", "out/out.mp4", 400, 400, fps)
		render.PlayVideo("out/out.mp4")
		break
	}
}

//revive:disable-next-line:unused-parameter
func renderFrameMandel(context *cairo.Context, width, height, percent float64) {
	cp := complexplane.FromCenterAndSize(-0.5, -0.0, 3, 3)

	it := iterator.New(context, cp)
	it.FractalFunc = algos.Mandel()
	it.ColorFunc = colorizors.GreyScale(0, 1)
	it.Iterate(40)
}

//revive:disable-next-line:unused-parameter
func renderFrameDuck(context *cairo.Context, width, height, percent float64) {
	cr := 0.41 + math.Cos(percent*blmath.Tau)*0.02
	ci := -0.21 + math.Sin(percent*blmath.Tau)*0.01
	// cr := 0.4
	// ci := -0.2
	cp := complexplane.FromCenterAndSize(-0.0, 0, 4, 4)

	it := iterator.New(context, cp)
	it.FractalFunc = algos.Duck(cr, ci)
	it.ColorFunc = colorizors.Duck(1, 0.4, 0.4)
	// it.WarpFunc = warpers.Simplex(blmath.LerpSin(percent, 0, 0.2), 0.005, blmath.LerpSin(percent, 5, 30))
	// it.WarpFunc = warpers.Fisheye(
	// 	width/2+math.Cos(percent*blmath.Tau)*width/2,
	// 	height/2+math.Sin(percent*blmath.Tau)*height/2,
	// 	width,
	// )
	it.WarpFunc = warpers.Swirl(width/2, height/2, blmath.LerpSin(percent, width, width/4))
	it.Iterate(40)
}

//revive:disable-next-line:unused-parameter
func renderFrameKali(context *cairo.Context, width, height, percent float64) {
	cr := -0.2 + math.Sin(percent*blmath.Tau)*0.1
	ci := -1.0 + math.Cos(percent*blmath.Tau)*0.1
	size := 2.6
	cp := complexplane.FromCenterAndSize(0.0, 0.0, size, size)

	it := iterator.New(context, cp)
	it.FractalFunc = algos.Kali(cr, ci, 2)
	it.ColorFunc = colorizors.GreyScale(1, 0)
	it.Iterate(80)
}

//revive:disable-next-line:unused-parameter
func renderFrameGrid(context *cairo.Context, width, height, percent float64) {
	size := 10.0
	cp := complexplane.FromCenterAndSize(0.0, 0.0, size, size)

	it := iterator.New(context, cp)
	it.FractalFunc = algos.Grid(cp, 30)
	it.ColorFunc = colorizors.GreyScale(1, 0)
	it.WarpFunc = warpers.Fisheye(width/2, height/2, width/2)
	it.Iterate(80)
}
