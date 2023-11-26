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

// Duck returns a fractal duck algorithm.
// The cr and ci params define the complex number c used in the algorithm.
// z if formed from the r and i params of the inner function.
func Duck(cr, ci float64) Algo {
	c := complex(cr, ci)
	return func(r, i, iter float64) float64 {
		// real and imag are flipped here to rotate the plane for aesthetic reasons.
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

// Julia returns a basic Julia algorithm.
// cr and ci are the complex components of c.
func Julia(cr, ci float64) Algo {
	return func(r, i, iter float64) float64 {
		c := complex(cr, ci)
		z := complex(r, i)
		for n := 0.0; n < iter; n++ {
			z = z*z + c
			if blmath.ComplexMagnitude(z) > 2 {
				return n / iter
			}
		}
		return 0.0
	}
}

// Kali returns a kali fractal function.
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

// NovaBase returns a nova fractal function.
// This version has perfect radial symmetry.
// power generally determines how many arms are created.
// smaller bailOut usually makes for a more complex image. You might go down to 0.00001 or lower
func NovaBase(power, bailOut float64) Algo {
	return func(r, i, iter float64) float64 {
		z := complex(r, i)
		c := complex(r, i)
		p := complex(power, 0)
		n := 0.0
		for ; n < iter; n++ {
			z1 := z - (cmplx.Pow(z, p)-1)/(p*cmplx.Pow(z, p-1)) + c
			// quit when we've converged close enough, based on the bailOut.
			if cmplx.Abs(z-z1) < bailOut {
				break
			}
			z = z1
		}
		return n / iter
	}
}

// RelaxedNova returns a nova fractal function with a relaxation param.
// power generally determines how many arms are created.
// zr and zi are the components of z.
// rxr and rxi are the components of the complex relaxation.
// smaller bailOut usually makes for a more complex image. You might go down to 0.00001 or lower
func NovaRelaxed(zr, zi, rxr, rxi, power, bailOut float64) Algo {
	return func(r, i, iter float64) float64 {
		z := complex(zr, zi)
		c := complex(r, i)
		p := complex(power, 0)
		rx := complex(rxr, rxi)
		n := 0.0
		for ; n < iter; n++ {
			z1 := z - rx*(cmplx.Pow(z, p)-1)/(p*cmplx.Pow(z, p-1)) + c
			// quit when we've converged close enough, based on the bailOut.
			if cmplx.Abs(z-z1) < bailOut {
				break
			}
			z = z1
		}
		return n / iter
	}
}

// NovaZ returns a nova fractal function.
// This version lets you set the inital values of z.
// power generally determines how many arms are created.
// smaller bailOut usually makes for a more complex image. You might go down to 0.00001 or lower
func NovaZ(zr, zi, power, bailOut float64) Algo {
	return func(r, i, iter float64) float64 {
		z := complex(zr, zi)
		c := complex(r, i)
		p := complex(power, 0)
		n := 0.0
		for ; n < iter; n++ {
			z1 := z - (cmplx.Pow(z, p)-1)/(p*cmplx.Pow(z, p-1)) + c
			// quit when we've converged close enough, based on the bailOut.
			if cmplx.Abs(z-z1) < bailOut {
				break
			}
			z = z1
		}
		return n / iter
	}
}
