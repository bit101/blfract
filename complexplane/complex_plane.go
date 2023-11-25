// Package complexplane defines a complex plane section.
package complexplane

// ComplexPlane represents the area of the complex plane that will be rendered.
type ComplexPlane struct {
	RealMin, RealMax, ImagMin, ImagMax float64
}

// FromCenterAndSize creates a complex plane based on the center point, width and height
func FromCenterAndSize(cr, ci, width, height float64) *ComplexPlane {
	return &ComplexPlane{
		RealMin: cr - width/2,
		RealMax: cr + width/2,
		ImagMin: ci - height/2,
		ImagMax: ci + height/2,
	}
}

// FromRect creates a complex plane based on the top left point, width and height
func FromRect(r, i, width, height float64) *ComplexPlane {
	return &ComplexPlane{
		RealMin: r,
		RealMax: r + width,
		ImagMin: i,
		ImagMax: i + height,
	}
}
