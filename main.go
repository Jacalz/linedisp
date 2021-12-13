package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/Jacalz/linalg/matrix"
	"github.com/Jacalz/linalg/rn"
)

func main() {
	a := app.New()
	w := a.NewWindow("Opera ❤️  Linear Algebra")

	// Outside points:
	p1 := rn.VecN{0.0, 6.5, 0}
	p2 := rn.VecN{3.0, 3.0, 0}
	p3 := rn.VecN{7.0, 0.0, 0}
	p4 := rn.VecN{10.0, 3.0, 0}
	p5 := rn.VecN{14.0, 6.5, 0}
	p6 := rn.VecN{10.0, 10.0, 0}
	p7 := rn.VecN{7.0, 13.0, 0}
	p8 := rn.VecN{3.0, 10.0, 0}

	// Inside points:
	p9 := rn.VecN{6.0, 5.0, 0}
	p10 := rn.VecN{8.0, 5.0, 0}
	p11 := rn.VecN{5.0, 6.0, 0}
	p12 := rn.VecN{9.0, 6.0, 0}
	p13 := rn.VecN{5.0, 7.0, 0}
	p14 := rn.VecN{9.0, 7.0, 0}
	p15 := rn.VecN{6.0, 8.0, 0}
	p16 := rn.VecN{8.0, 8.0, 0}

	M := matrix.NewFromVec(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14, p15, p16)

	// Scale by a factor s in all directions.
	s := float64(40)
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
