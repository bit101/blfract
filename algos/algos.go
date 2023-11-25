// Package algos contains fractal algorithms
package algos

import (
	"math"
	"math/cmplx"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/blfract/complexplane"
)

// Algo represents a specific fractal algorithm called on each rendered point of a complex plane.
// The functions in this package all return other functions that are used in the iterators to process each pixel.
type Algo func(re, im, iter float64) float64

// Grid just generates a grid of res many lines. Not actually a fractal at all, but useful for debugging warpers.
func Grid(cp *complexplane.ComplexPlane, res float64) Algo {
	return func(r, i, iter float64) float64 {
		x := blmath.Map(r, cp.RealMin, cp.RealMax, 0, res)
		y := blmath.Map(i, cp.ImagMin, cp.ImagMax, 0, res)
		if math.Abs(math.Mod(x, 1)) < 0.05 || math.Abs(math.Mod(y, 1)) < 0.05 {
			return 1.0
		}
		return 0.0
	}
}

// Mandel returns a basic Mandelbrot algorithm.
// There are no customizable parameters for this algorithm.
// This is because c is formed from the r and i params of the inner function.
// And z is always 0 + 0i.
func Mandel() Algo {
	return func(r, i, iter float64) float64 {
		c := complex(r, i)
		z := complex(0, 0)
		for n := 0.0; n < iter; n++ {
			z = z*z + c
			if blmath.ComplexMagnitude(z) > 2 {
				return n / iter
			}
		}
		return 0.0
	}
}

// Duck returns a fractal duck algorithm.
// The cr and ci params define the complex number c used in the algorithm.
// z if formed from the r and i params of the inner function.
func Duck(cr, ci float64) Algo {
	// real and imag are flipped here to rotate the plane for aesthetic reasons
	c := complex(cr, ci)
	return func(r, i, iter float64) float64 {
		z := complex(i, r)
		m := 0.0
		for i := 0.0; i < iter; i++ {
			z = cmplx.Log10(blmath.ComplexImagAbs(z) + c)
			m += blmath.ComplexMagnitude(z)
		}
		m /= iter
		return math.Max(m, 0.5)
	}
}

// Nova returns a nova fractal function.
func Nova() Algo {
	return func(r, i, iter float64) float64 {
		c := complex(r, i)
		z := complex(0, 0)
		for i := 0.0; i < iter; i++ {
			z = (z - cmplx.Pow(z-1, 3) + c) / (3 * z * z)
			if blmath.ComplexMagnitude(z) > 2 {
				return i / iter
			}
		}
		return 0.0
	}
}

// Kali returns a nova fractal function.
func Kali(cr, ci, bailOut float64) Algo {
	c := complex(cr, ci)
	return func(r, i, iter float64) float64 {
		z := complex(r, i)
		for n := 0.0; n < iter; n++ {
			r := math.Abs(real(z))
			i := math.Abs(imag(z))
			m := r*r + i*i
			r = r/m + real(c)
			i = i/m + imag(c)
			z = complex(r, i)
			if blmath.ComplexMagnitude(z) > bailOut {
				return n / iter
			}
		}
		return 0.0
	}
}
