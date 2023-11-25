// Package colorizors holds colorizing functions.
package colorizors

import (
	"math"

	"github.com/bit101/bitlib/blcolor"
	"github.com/bit101/bitlib/blmath"
)

// Colorizor is a function that determines what color a point will be rendered with.
type Colorizor func(m float64) blcolor.Color

// GreyScale is a simple greyscale colorizor
func GreyScale(min, max float64) Colorizor {
	return func(m float64) blcolor.Color {
		m = blmath.Lerp(m, min, max)
		return blcolor.Grey(m)
	}
}

// Hue is a simple greyscale colorizor
func Hue(min, max float64) Colorizor {
	return func(m float64) blcolor.Color {
		m = blmath.Lerp(m, min, max)
		return blcolor.HSV(m, 1, 1)
	}
}

// Duck returns a duck fractal colorizor.
func Duck(r, g, b float64) Colorizor {
	return func(m float64) blcolor.Color {
		co := 1.0 - math.Log2(0.5*math.Log2(m*2.0))
		red := 0.5 + 0.5*math.Cos(blmath.TwoPi*co+r)
		green := 0.5 + 0.5*math.Cos(blmath.TwoPi*co+g)
		blue := 0.5 + 0.5*math.Cos(blmath.TwoPi*co+b)
		return blcolor.RGB(red, green, blue)
	}
}
