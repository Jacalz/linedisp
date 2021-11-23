package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/Jacalz/linalg/matrix"
	"github.com/Jacalz/linalg/rn"
)

func main() {
	a := app.New()
	w := a.NewWindow("Linear Algebra Graphics")

	// These are actually points, but use vectors for simplicity.
	p1 := rn.VecN{1.000, -0.800, 0.000}
	p2 := rn.VecN{0.500, -0.800, -0.866}
	p3 := rn.VecN{-0.500, -0.800, -0.866}
	p4 := rn.VecN{-1.000, -0.800, 0.000}
	p5 := rn.VecN{-0.500, -0.800, 0.866}
	p6 := rn.VecN{0.500, -0.800, 0.866}
	p7 := rn.VecN{0.840, -0.400, 0.000}
	p8 := rn.VecN{0.315, 0.125, -0.546}
	p9 := rn.VecN{-0.210, 0.650, -0.364}
	p10 := rn.VecN{-0.360, 0.800, 0.000}
	p11 := rn.VecN{-0.210, 0.650, 0.364}
	p12 := rn.VecN{0.315, 0.125, 0.546}

	M := matrix.NewFromVec(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12)

	// Transpose by u in specific directions.
	//u := rn.VecN{2, 1, 0}
	//M, _ = matrix.AddVec(M, u)

	// Scale by a factor s in all directions.
	s := float64(200)
	T := matrix.Matrix{
		{s, 0, 0},
		{0, s, 0},
		{0, 0, s},
	}
	M, _ = matrix.Mult(T, M)

	w.SetContent(NewLineDrawer(M))
	w.Resize(fyne.NewSize(600, 600))
	w.ShowAndRun()
}
