package testblas

import (
	"testing"

	"github.com/gonum/blas"
)

type Dtrmver interface {
	Dtrmv(o blas.Order, ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int,
		a []float64, lda int, x []float64, incX int)
}

func DtrmvTest(t *testing.T, blasser Dtrmver) {
	x1 := []float64{2, 3, 4}
	x2 := []float64{2, 1, 3, 1, 4}
	incX2 := 2

	//ul == blas.Upper
	tr := []float64{1, 0, 0, 2, 3, 0, 0, 4, 5}

	//d == blas.NonUnit
	solNoTrans := []float64{8, 25, 20}
	solTrans := []float64{2, 13, 32}

	in := make([]float64, len(x1))
	copy(in, x1)
	blasser.Dtrmv(blas.ColMajor, blas.Upper, blas.NoTrans, blas.NonUnit, 3, tr, 3, in, 1)

	if !dSliceTolEqual(in, solNoTrans) {
		t.Error("Wrong Dtrmv result for: ColMajor, Upper, NoTrans, NonUnit, IncX==1")
	}

	in = make([]float64, len(x1))
	copy(in, x1)
	blasser.Dtrmv(blas.ColMajor, blas.Upper, blas.Trans, blas.NonUnit, 3, tr, 3, in, 1)

	if !dSliceTolEqual(in, solTrans) {
		t.Error("Wrong Dtrmv result for: ColMajor, Upper, Trans, NonUnit, IncX==1")
	}

	in = make([]float64, len(x2))
	copy(in, x2)
	blasser.Dtrmv(blas.ColMajor, blas.Upper, blas.NoTrans, blas.NonUnit, 3, tr, 3, in, 2)

	if !dStridedSliceTolEqual(3, in, incX2, solNoTrans, 1) {
		t.Error("Wrong Dtrmv result for: ColMajor, Upper, NoTrans, NonUnit, IncX==2")
	}

	in = make([]float64, len(x2))
	copy(in, x2)
	blasser.Dtrmv(blas.ColMajor, blas.Upper, blas.Trans, blas.NonUnit, 3, tr, 3, in, 2)

	if !dStridedSliceTolEqual(3, in, incX2, solTrans, 1) {
		t.Error("Wrong Dtrmv result for: ColMajor, Upper, Trans, NonUnit, IncX==2")
	}

	//d == blas.Unit
	solNoTrans = []float64{8, 19, 4}
	solTrans = []float64{2, 7, 16}

	in = make([]float64, len(x1))
	copy(in, x1)
	blasser.Dtrmv(blas.ColMajor, blas.Upper, blas.NoTrans, blas.Unit, 3, tr, 3, in, 1)

	if !dSliceTolEqual(in, solNoTrans) {
		t.Error("Wrong Dtrmv result for: ColMajor, Upper, NoTrans, Unit, IncX==1")
	}

	in = make([]float64, len(x1))
	copy(in, x1)
	blasser.Dtrmv(blas.ColMajor, blas.Upper, blas.Trans, blas.Unit, 3, tr, 3, in, 1)

	if !dSliceTolEqual(in, solTrans) {
		t.Error("Wrong Dtrmv result for: ColMajor, Upper, Trans, Unit, IncX==1")
	}

	in = make([]float64, len(x2))
	copy(in, x2)
	blasser.Dtrmv(blas.ColMajor, blas.Upper, blas.NoTrans, blas.Unit, 3, tr, 3, in, 2)

	if !dStridedSliceTolEqual(3, in, incX2, solNoTrans, 1) {
		t.Error("Wrong Dtrmv result for: ColMajor, Upper, NoTrans, Unit, IncX==2")
	}

	in = make([]float64, len(x2))
	copy(in, x2)
	blasser.Dtrmv(blas.ColMajor, blas.Upper, blas.Trans, blas.Unit, 3, tr, 3, in, 2)

	if !dStridedSliceTolEqual(3, in, incX2, solTrans, 1) {
		t.Error("Wrong Dtrmv result for: ColMajor, Upper, Trans, Unit, IncX==2")
	}

	//ul == blas.Lower
	tr = []float64{1, 2, 0, 0, 3, 4, 0, 0, 5}

	//d == blas.NonUnit
	solNoTrans = []float64{2, 13, 32}
	solTrans = []float64{8, 25, 20}

	in = make([]float64, len(x1))
	copy(in, x1)
	blasser.Dtrmv(blas.ColMajor, blas.Lower, blas.NoTrans, blas.NonUnit, 3, tr, 3, in, 1)

	if !dSliceTolEqual(in, solNoTrans) {
		t.Error("Wrong Dtrmv result for: ColMajor, Lower, NoTrans, NonUnit, IncX==1")
	}

	in = make([]float64, len(x1))
	copy(in, x1)
	blasser.Dtrmv(blas.ColMajor, blas.Lower, blas.Trans, blas.NonUnit, 3, tr, 3, in, 1)

	if !dSliceTolEqual(in, solTrans) {
		t.Error("Wrong Dtrmv result for: ColMajor, Lower, Trans, NonUnit, IncX==1")
	}

	in = make([]float64, len(x2))
	copy(in, x2)
	blasser.Dtrmv(blas.ColMajor, blas.Lower, blas.NoTrans, blas.NonUnit, 3, tr, 3, in, 2)

	if !dStridedSliceTolEqual(3, in, incX2, solNoTrans, 1) {
		t.Error("Wrong Dtrmv result for: ColMajor, Lower, NoTrans, NonUnit, IncX==2")
	}

	in = make([]float64, len(x2))
	copy(in, x2)
	blasser.Dtrmv(blas.ColMajor, blas.Lower, blas.Trans, blas.NonUnit, 3, tr, 3, in, 2)

	if !dStridedSliceTolEqual(3, in, incX2, solTrans, 1) {
		t.Error("Wrong Dtrmv result for: ColMajor, Lower, Trans, NonUnit, IncX==2")
	}

	//d == blas.Unit
	solNoTrans = []float64{2, 7, 16}
	solTrans = []float64{8, 19, 4}

	in = make([]float64, len(x1))
	copy(in, x1)
	blasser.Dtrmv(blas.ColMajor, blas.Lower, blas.NoTrans, blas.Unit, 3, tr, 3, in, 1)

	if !dSliceTolEqual(in, solNoTrans) {
		t.Error("Wrong Dtrmv result for: ColMajor, Lower, NoTrans, Unit, IncX==1")
	}

	in = make([]float64, len(x1))
	copy(in, x1)
	blasser.Dtrmv(blas.ColMajor, blas.Lower, blas.Trans, blas.Unit, 3, tr, 3, in, 1)

	if !dSliceTolEqual(in, solTrans) {
		t.Error("Wrong Dtrmv result for: ColMajor, Lower, Trans, Unit, IncX==1")
	}

	in = make([]float64, len(x2))
	copy(in, x2)
	blasser.Dtrmv(blas.ColMajor, blas.Lower, blas.NoTrans, blas.Unit, 3, tr, 3, in, 2)

	if !dStridedSliceTolEqual(3, in, incX2, solNoTrans, 1) {
		t.Error("Wrong Dtrmv result for: ColMajor, Lower, NoTrans, Unit, IncX==2")
	}

	in = make([]float64, len(x2))
	copy(in, x2)
	blasser.Dtrmv(blas.ColMajor, blas.Lower, blas.Trans, blas.Unit, 3, tr, 3, in, 2)

	if !dStridedSliceTolEqual(3, in, incX2, solTrans, 1) {
		t.Error("Wrong Dtrmv result for: ColMajor, Lower, Trans, Unit, IncX==2")
	}
}