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

	switch renderTarget {
	case target.Image:
		w := 1000.0
		h := 1000.0
		render.Image(w, h, "out/out.png", renderFrame, 0.0)
		render.ViewImage("out/out.png")
		break

	case target.Montage:
		w := 200.0
		h := 200.0
		numFrames := 49
		render.Frames(w, h, numFrames, "out/frames", renderFrame)
		render.MakeMontage("out/frames", "out/out.png")
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

func renderFrame(context *cairo.Context, width, height, percent float64) {

	// 1. Create the complex plane. ==========================
	cp := complexplane.FromCenterAndSize(-0.0, 0, 4, 4)

	// 2. Set params. ========================================
	// algo
	cr := 0.40
	ci := -0.20
	// color
	r := 1.0
	g := 0.4
	b := 0.4
	x := 2.0
	// warp
	wavelength := 100.0
	offset := 50.0
	phase := 0.0
	// iter
	iterations := 20

	// 3. Create the iterator. ===============================
	it := iterator.New(context, cp)

	// 4. Set the algorithm. =================================
	it.FractalFunc = algos.Duck(cr, ci)

	// 5. Set the color function. (optional) =================
	it.ColorFunc = colorizors.Duck(r, g, b, x)

	// 6. Set the warp function. (optional) ==================
	it.WarpFunc = warpers.Ripple(width/2, height/2, wavelength, offset, phase)

	// 7. Iterate. ===========================================
	it.Iterate(iterations)

	// 8. Print parametes on image (optional) ================
	it.PrintParams("cr", cr, "ci", ci, "iter", iterations)
}
