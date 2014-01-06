// Uses the netlib standard. Other implementations may differ. Difference
// is that the code panics for n < 0 and incx == 0 rather than returning zero.
// (Documentation says incx must not be zero)
//
// TODO: Improve documentation
package naivegoblas

import (
	"github.com/gonum/blas"
	"math"
)

type Blas struct{}

var negativeN = "blas: negative number of elements"
var zeroInc = "blas: zero value of increment"
var negInc = "blas: negative value of increment"

// Ddot computes the dot product of the two vectors \sum_i x[i]*y[i]
func (Blas) Ddot(n int, x []float64, incX int, y []float64, incY int) float64 {
	if n < 0 {
		panic(negativeN)
	}
	if incX == 0 || incY == 0 {
		panic(zeroInc)
	}
	var ix, iy int
	var sum float64
	if incX < 0 {
		ix = (-n + 1) * incX
	}
	if incY < 0 {
		iy = (-n + 1) * incY
	}
	for i := 0; i < n; i++ {
		sum += y[iy] * x[ix]
		ix += incX
		iy += incY
	}
	return sum
}

// Dnrm2 computes the euclidean norm of a vector via the function
// name so that
//       dnrm2 = sqrt(x'x)
// This function also does not allow negative increments, see:
// http://www.netlib.org/blas/dnrm2.f
func (Blas) Dnrm2(n int, x []float64, incX int) float64 {
	if incX < 1 {
		if incX == 0 {
			panic(zeroInc)
		}
		panic(negInc)
	}
	if n < 2 {
		if n == 1 {
			return math.Abs(x[0])
		}
		if n == 0 {
			return 0
		}
		if n < 1 {
			panic(negativeN)
		}
	}
	scale := 0.0
	sumSquares := 1.0
	for ix := 0; ix < n*incX; ix += incX {
		val := x[ix]
		if val == 0 {
			continue
		}
		absxi := math.Abs(x[ix])
		if scale < absxi {
			sumSquares = 1 + sumSquares*(scale/absxi)*(scale/absxi)
			scale = absxi
		} else {
			sumSquares = sumSquares + (absxi/scale)*(absxi/scale)
		}
	}
	return scale * math.Sqrt(sumSquares)
}

// Dasum computes the sum of the absolute values of the elements of x
// Dasum returns for negative increment in the netlib package (seems
// to differ from behavior of other routines) and so it panics here
func (Blas) Dasum(n int, x []float64, incX int) float64 {
	var sum float64
	if n < 0 {
		panic(negativeN)
	}
	if incX == 0 {
		panic(zeroInc)
	}
	if incX < 0 {
		panic(negInc)
	}
	for i := 0; i < n; i++ {
		sum += math.Abs(x[i*incX])
	}
	return sum
}

// Idamax returns the index of the largest element of x. If there are multiple
// such indices it returns the earliest
func (Blas) Idamax(n int, x []float64, incX int) int {
	if incX < 1 {
		if incX == 0 {
			panic(zeroInc)
		}
		panic(negInc)
	}
	if n < 2 {
		if n == 1 {
			return 0
		}
		if n == 0 {
			return 0 // Netlib returns first index when n == 0
		}
		if n < 1 {
			panic(negativeN)
		}
	}
	idx := 0
	max := x[0]

	for i := 1; i < n; i++ {
		v := x[i*incX]
		if v > max {
			max = v
			idx = i * incX
		}
	}
	return idx
}

// Dswap interchanges two vectors
func (Blas) Dswap(n int, x []float64, incX int, y []float64, incY int) {
	if n < 1 {
		if n == 0 {
			return
		}
		panic(negativeN)
	}
	if incX == 0 || incY == 0 {
		panic(zeroInc)
	}
	var ix, iy int
	if incX < 0 {
		ix = (-n + 1) * incX
	}
	if incY < 0 {
		iy = (-n + 1) * incY
	}
	for i := 0; i < n; i++ {
		x[ix], y[iy] = y[iy], x[ix]
		ix += incX
		iy += incY
	}
}

func (Blas) Dcopy(n int, x []float64, incX int, y []float64, incY int) {
	if n < 1 {
		if n == 0 {
			return
		}
		panic(negativeN)
	}
	if incX == 0 || incY == 0 {
		panic(zeroInc)
	}
	var ix, iy int
	if incX < 0 {
		ix = (-n + 1) * incX
	}
	if incY < 0 {
		iy = (-n + 1) * incY
	}
	for i := 0; i < n; i++ {
		y[iy] = x[ix]
		ix += incX
		iy += incY
	}
}

// Daxpy computes y <- α x + y
func (Blas) Daxpy(n int, alpha float64, x []float64, incX int, y []float64, incY int) {
	if n < 1 {
		if n == 0 {
			return
		}
		panic(negativeN)
	}
	if incX == 0 || incY == 0 {
		panic(zeroInc)
	}
	if alpha == 0 {
		return
	}
	var ix, iy int
	if incX < 0 {
		ix = (-n + 1) * incX
	}
	if incY < 0 {
		iy = (-n + 1) * incY
	}
	for i := 0; i < n; i++ {
		y[iy] += alpha * x[ix]
		ix += incX
		iy += incY
	}
}

// DrotG gives plane rotation
//
// _      _    _   _     _   _
// | c  s |    | a |     | r |
// | -s c |  * | b |   = | 0 |
// _      _    _   _     _   _
//
// r = ±(a^2 + b^2)
// c = a/r, the cosine of the plane rotation
// s = b/r, the sine of the plane rotation
//
// NOTE: Netlib library seems to give a different
// sign for r when a or b is zero than other implementations
// and different than BLAS technical manual
func (Blas) Drotg(a, b float64) (c, s, r, z float64) {
	if b == 0 && a == 0 {
		return 1, 0, a, 0
	}
	/*
		if a == 0 {
			if b > 0 {
				return 0, 1, b, 1
			}
			return 0, -1, -b, 1
		}
	*/
	absA := math.Abs(a)
	absB := math.Abs(b)
	aGTb := absA > absB
	r = math.Hypot(a, b)
	if aGTb {
		r = math.Copysign(r, a)
	} else {
		r = math.Copysign(r, b)
	}
	c = a / r
	s = b / r
	if aGTb {
		z = s
	} else if c != 0 { // r == 0 case handled above
		z = 1 / c
	} else {
		z = 1
	}
	return
}

// Drotmg computes the modified Givens rotation. See
// http://www.netlib.org/lapack/explore-html/df/deb/drotmg_8f.html
// for more details
func (Blas) Drotmg(d1, d2, x1, y1 float64) (p blas.DrotmParams, rd1, rd2, rx1 float64) {
	var p1, p2, q1, q2, u float64

	gam := 4096.0
	gamsq := 16777216.0
	rgamsq := 5.9604645e-8
	p = blas.DrotmParams{}
	if d1 < 0 {
		p.Flag = -1
		return
	}

	p2 = d2 * y1
	if p2 == 0 {
		p.Flag = -2
		rd1 = d1
		rd2 = d2
		rx1 = x1
		return
	}
	p1 = d1 * x1
	q2 = p2 * y1
	q1 = p1 * x1

	absQ1 := math.Abs(q1)
	absQ2 := math.Abs(q2)

	if absQ1 < absQ2 && q2 < 0 {
		p.Flag = -1
		return
	}

	if d1 == 0 {
		p.Flag = 1
		p.H[0] = p1 / p2
		p.H[3] = x1 / y1
		u = 1 + p.H[0]*p.H[3]
		rd1, rd2 = d2/u, d1/u
		rx1 = y1 / u
		return
	}

	// Now we know that d1 != 0, and d2 != 0. If d2 == 0, it would be caught
	// when p2 == 0, and if d1 == 0, then it is caught above
	//fmt.Println("absq1", absQ1)
	//fmt.Println("absq2", absQ2)

	if math.Abs(q1) > math.Abs(q2) {
		p.H[1] = -y1 / x1
		p.H[2] = p2 / p1
		u = 1 - p.H[2]*p.H[1]
		rd1 = d1
		rd2 = d2
		rx1 = x1
		p.Flag = 0
		// u must be greater than zero because |q1| > |q2|, so check from netlib
		// is unnecessary
		// This is left in for ease of comparison with complex routines
		//if u > 0 {
		rd1 /= u
		rd2 /= u
		rx1 *= u
		//}
	} else {
		p.Flag = 1
		p.H[0] = p1 / p2
		p.H[3] = x1 / y1
		u = 1 + p.H[0]*p.H[3]
		rd1 = d2 / u
		rd2 = d1 / u
		rx1 = y1 * u

	}
	//fmt.Println("Flag = ", p.Flag)
	//fmt.Println("rd1 = ", rd1)
	for rd1 <= rgamsq || rd1 >= gamsq {
		if p.Flag == 0 {
			p.H[0] = 1
			p.H[3] = 1
			p.Flag = -1
		} else {
			p.H[1] = -1
			p.H[2] = 1
			p.Flag = -1
		}
		if rd1 <= rgamsq {
			rd1 *= gam * gam
			rx1 /= gam
			p.H[0] /= gam
			p.H[2] /= gam
		} else {
			rd1 /= gam * gam
			rx1 *= gam
			p.H[0] *= gam
			p.H[2] *= gam
		}
	}
	//fmt.Println("rd2 = ", rd2)
	for math.Abs(rd2) <= rgamsq || math.Abs(rd2) >= gamsq {
		if p.Flag == 0 {
			p.H[0] = 1
			p.H[3] = 1
			p.Flag = -1
		} else {
			p.H[1] = -1
			p.H[2] = 1
			p.Flag = -1
		}
		if math.Abs(rd2) <= rgamsq {
			rd2 *= gam * gam
			p.H[1] /= gam
			p.H[3] /= gam
		} else {
			rd2 /= gam * gam
			p.H[1] *= gam
			p.H[3] *= gam
		}
	}
	return
}

// Drot applies a plane transformation
func (Blas) Drot(n int, x []float64, incX int, y []float64, incY int, c float64, s float64) {
	if n < 1 {
		if n == 0 {
			return
		}
		panic(negativeN)
	}
	if incX == 0 || incY == 0 {
		panic(zeroInc)
	}
	var ix, iy int
	if incX < 0 {
		ix = (-n + 1) * incX
	}
	if incY < 0 {
		iy = (-n + 1) * incY
	}
	for i := 0; i < n; i++ {
		x[ix], y[iy] = c*x[ix]+s*y[iy], c*y[iy]-s*x[ix]
		ix += incX
		iy += incY
	}
}

// Drotm applies the modified Givens rotation to the 2 x N matrix
func (Blas) Drotm(n int, x []float64, incX int, y []float64, incY int, p blas.DrotmParams) {
	// Some weirdness between this and drotmg. Maybe drotmg should panic on flag of -2?
	// Need to check if it's used elsewhere
	//
	// Same with some terms being zero
	//
	// Odd that this doesn't panic with incX == 0 or incY == 0 like all the others

	if n <= 0 {
		if n == 0 {
			return
		}
		panic(negativeN)
	}
	if incX == 0 || incY == 0 {
		panic(zeroInc)
	}
	flag := p.Flag
	if flag == -2 {
		panic("flag is negative 2")
	}
	if incX == incY && incX > 0 {
		nsteps := n * incX
		if flag < 0 {
			h11 := p.H[0]
			h12 := p.H[2]
			h21 := p.H[1]
			h22 := p.H[3]
			for i := 0; i < nsteps; i += incX {
				w := x[i]
				z := y[i]
				x[i] = w*h11 + z*h12
				y[i] = w*h21 + z*h22
			}
			return
		}
		if flag == 0 {
			h12 := p.H[2]
			h21 := p.H[1]
			for i := 0; i < nsteps; i += incX {
				w := x[i]
				z := y[i]
				x[i] = w + z*h12
				y[i] = w*h21 + z
			}
			return
		}
		h11 := p.H[0]
		h22 := p.H[3]
		for i := 0; i < nsteps; i += incX {
			w := x[i]
			z := y[i]
			x[i] = w*h11 + z
			y[i] = -w + h22*z
		}
		return
	}
	ix := 0
	iy := 0
	if incX < 0 {
		ix = (-n + 1) * incX
	}
	if incY < 0 {
		iy = (-n + 1) * incY
	}
	if flag < 0 {
		h11 := p.H[0]
		h12 := p.H[2]
		h21 := p.H[1]
		h22 := p.H[3]
		for i := 0; i < n; i++ {
			w := x[ix]
			z := y[iy]
			x[ix] = w*h11 + z*h12
			y[iy] = w*h21 + z*h22
			ix += incX
			iy += incY
		}
		return
	}
	if flag == 0 {
		h12 := p.H[2]
		h21 := p.H[1]
		for i := 0; i < n; i++ {
			w := x[ix]
			z := y[iy]
			x[ix] = w + z*h12
			y[iy] = w*h21 + z
			ix += incX
			iy += incY
		}
		return
	}
	h11 := p.H[0]
	h22 := p.H[3]
	for i := 0; i < n; i++ {
		w := x[ix]
		z := y[iy]
		x[ix] = w*h11 + z
		y[iy] = -w + z*h22
		ix += incX
		iy += incY
	}
	return
}

func (Blas) Dscal(n int, alpha float64, x []float64, incX int) {
	if incX < 1 {
		if incX == 0 {
			panic(zeroInc)
		}
		panic(negInc)
	}
	if n < 1 {
		if n == 0 {
			return
		}
		if n < 1 {
			panic(negativeN)
		}
	}
	for ix := 0; ix < n*incX; ix += incX {
		x[ix] *= alpha
	}
	return
}