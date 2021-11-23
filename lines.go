package main

import (
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/Jacalz/linalg/matrix"
)

type LineDrawer struct {
	widget.BaseWidget
	lines  []fyne.CanvasObject
	matrix matrix.Matrix
}

func NewLineDrawer(matrix matrix.Matrix) *LineDrawer {
	return &LineDrawer{
		lines:  LinesFromMatrix(matrix),
		matrix: matrix,
	}
}

func LinesFromMatrix(M matrix.Matrix) []fyne.CanvasObject {
	return []fyne.CanvasObject{
		NewLineBetween(M[0][0], M[0][1], M[1][0], M[1][1]),
		NewLineBetween(M[0][1], M[0][2], M[1][1], M[1][2]),
		NewLineBetween(M[0][2], M[0][3], M[1][2], M[1][3]),
		NewLineBetween(M[0][3], M[0][4], M[1][3], M[1][4]),
		NewLineBetween(M[0][4], M[0][5], M[1][4], M[1][5]),
		NewLineBetween(M[0][5], M[0][0], M[1][5], M[1][0]),
		NewLineBetween(M[0][6], M[0][7], M[1][6], M[1][7]),
		NewLineBetween(M[0][7], M[0][8], M[1][7], M[1][8]),
		NewLineBetween(M[0][8], M[0][9], M[1][8], M[1][9]),
		NewLineBetween(M[0][9], M[0][10], M[1][9], M[1][10]),
		NewLineBetween(M[0][10], M[0][11], M[1][10], M[1][11]),
		NewLineBetween(M[0][11], M[0][6], M[1][11], M[1][6]),
		NewLineBetween(M[0][0], M[0][6], M[1][0], M[1][6]),
		NewLineBetween(M[0][1], M[0][7], M[1][1], M[1][7]),
		NewLineBetween(M[0][2], M[0][8], M[1][2], M[1][8]),
		NewLineBetween(M[0][3], M[0][9], M[1][3], M[1][9]),
		NewLineBetween(M[0][4], M[0][10], M[1][4], M[1][10]),
		NewLineBetween(M[0][5], M[0][11], M[1][5], M[1][11]),
	}
}

func (l *LineDrawer) Scrolled(s *fyne.ScrollEvent) {
	a := float64(s.Scrolled.DY) / 8 // One scroll step seems to be 10.
	if a < 0 {
		a += 2 // Get it back into the positive range.
	}

	T := matrix.Matrix{
		{a, 0, 0},
		{0, a, 0},
		{0, 0, a},
	}

	l.matrix, _ = matrix.Mult(T, l.matrix)
	l.lines = LinesFromMatrix(l.matrix)
	l.Refresh()
}

func (l *LineDrawer) Dragged(d *fyne.DragEvent) {
	a := float64(d.Dragged.DY) * 0.007
	X := matrix.Matrix{
		{1, 0, 0},
		{0, math.Cos(a), -math.Sin(a)},
		{0, math.Sin(a), math.Cos(a)},
	}

	b := float64(d.Dragged.DX) * -0.007
	Y := matrix.Matrix{
		{math.Cos(b), 0, math.Sin(b)},
		{0, 1, 0},
		{-math.Sin(b), 0, math.Cos(b)},
	}

	l.matrix, _ = matrix.Mult(X, l.matrix)
	l.matrix, _ = matrix.Mult(Y, l.matrix)
	l.lines = LinesFromMatrix(l.matrix)
	l.Refresh()
}

func (l *LineDrawer) DragEnd() {

}

func (l *LineDrawer) CreateRenderer() fyne.WidgetRenderer {
	l.ExtendBaseWidget(l)
	return &lineRenderer{lineDrawer: l}
}

type lineRenderer struct {
	lineDrawer *LineDrawer
}

func (lr *lineRenderer) Destroy() {
}

func (lr *lineRenderer) Layout(s fyne.Size) {
}

func (lr *lineRenderer) MinSize() fyne.Size {
	return fyne.NewSize(theme.IconInlineSize(), theme.IconInlineSize())
}

func (lr *lineRenderer) Objects() []fyne.CanvasObject {
	return lr.lineDrawer.lines
}

func (lr *lineRenderer) Refresh() {
	canvas.Refresh(lr.lineDrawer)
}

func NewLineBetween(x1, y1, x2, y2 float64) *canvas.Line {
	return &canvas.Line{
		Position1:   fyne.NewPos(float32(x1)+300, float32(x2)+300),
		Position2:   fyne.NewPos(float32(y1)+300, float32(y2)+300),
		StrokeColor: theme.PrimaryColor(),
		StrokeWidth: 3,
	}
}
