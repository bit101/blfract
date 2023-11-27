// Package colorizors holds colorizing functions.
package colorizors

import (
	"math"

	"github.com/bit101/bitlib/blcolor"
	"github.com/bit101/bitlib/blmath"
)

// Colorizor is a function that determines what color a point will be rendered with.
type Colorizor func(m float64) blcolor.Color

// Binary returns a ...
func Binary(colA, colB blcolor.Color, iter int) Colorizor {
	return func(m float64) blcolor.Color {
		if int(m*float64(iter))%2 == 0 {
			return colA
		}
		return colB
	}
}

// BinaryZero returns a ...
func BinaryZero(colA, colB, zero blcolor.Color, iter int) Colorizor {
	return func(m float64) blcolor.Color {
		if m == 0 {
			return zero
		}
		if int(m*float64(iter))%2 == 0 {
			return colA
		}
		return colB
	}
}

// Duck returns a duck fractal colorizor.
// x should be in the range of 2 up to ... anything. Higher numbers give smoother gradients.
func Duck(r, g, b, x float64) Colorizor {
	return func(m float64) blcolor.Color {
		co := 1.0 - math.Log2(0.5*math.Log2(m*x))
		red := 0.5 + 0.5*math.Cos(blmath.TwoPi*co+r)
		green := 0.5 + 0.5*math.Cos(blmath.TwoPi*co+g)
		blue := 0.5 + 0.5*math.Cos(blmath.TwoPi*co+b)
		return blcolor.RGB(red, green, blue)
	}
}

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

// HSV is a simple greyscale colorizor
func HSV(minHue, maxHue, minSat, maxSat, minVal, maxVal float64) Colorizor {
	return func(m float64) blcolor.Color {
		h := blmath.Lerp(m, minHue, maxHue)
		s := blmath.Lerp(m, minSat, maxSat)
		v := blmath.Lerp(m, minVal, maxVal)
		return blcolor.HSV(h, s, v)
	}
}
