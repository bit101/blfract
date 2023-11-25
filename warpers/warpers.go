// Package warpers defines functions that pre-warp the complex plane
package warpers

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/noise"
)

// Warper is the interface for a warper function.
type Warper func(x, y float64) (float64, float64)

// Default returns a function that does not actually warp.
func Default() Warper {
	return func(x, y float64) (float64, float64) {
		return x, y
	}
}

// Simplex returns a warper that warps a complex plane via simplex noise.
func Simplex(z, scale, offset float64) Warper {
	return func(x, y float64) (float64, float64) {
		a := noise.Simplex3(x*scale, y*scale, z) * blmath.Tau
		return x + math.Cos(a)*offset, y + math.Sin(a)*offset
	}
}

// Fisheye returns a warper that warps a complex plane with a fisheye lens effect.
func Fisheye(cx, cy, radius float64) Warper {
	return func(x, y float64) (float64, float64) {
		x -= cx
		y -= cy
		r := radius / math.Hypot(x, y)
		d := 2 / (r + 1)
		return cx + x*d, cy + y*d
	}
}

// Swirl returns a warper that warps a complex plane with a fisheye lens effect.
func Swirl(cx, cy, radius float64) Warper {
	return func(x, y float64) (float64, float64) {
		x -= cx
		y -= cy
		r := math.Hypot(x, y) / radius
		r = r * r
		return cx + x*math.Cos(r) - y*math.Sin(r), cy + y*math.Cos(r) + x*math.Sin(r)
	}
}
