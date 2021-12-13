package main

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/Jacalz/linalg/matrix"
)

// LineDrawer draws lines from a matrix of position vectors.
type LineDrawer struct {
	widget.BaseWidget
	lines  []fyne.CanvasObject
	matrix matrix.Matrix
}

// NewLineDrawer creates a new LineDrawer with the given matrix.
func NewLineDrawer(matrix matrix.Matrix) *LineDrawer {
	return &LineDrawer{
		lines:  LinesFromMatrix(matrix),
		matrix: matrix,
	}
}

// LinesFromMatrix creates new canvas.Line from the matrix.
func LinesFromMatrix(M matrix.Matrix) []fyne.CanvasObject {
	return []fyne.CanvasObject{
		NewLineBetween(M[0][0], M[0][12], M[1][0], M[1][12]),
		NewLineBetween(M[0][0], M[0][10], M[1][0], M[1][10]),

		NewLineBetween(M[0][1], M[0][10], M[1][1], M[1][10]),
		NewLineBetween(M[0][1], M[0][8], M[1][1], M[1][8]),

		NewLineBetween(M[0][2], M[0][9], M[1][2], M[1][9]),
		NewLineBetween(M[0][2], M[0][8], M[1][2], M[1][8]),

		NewLineBetween(M[0][3], M[0][11], M[1][3], M[1][11]),
		NewLineBetween(M[0][3], M[0][9], M[1][3], M[1][9]),

		NewLineBetween(M[0][4], M[0][13], M[1][4], M[1][13]),
		NewLineBetween(M[0][4], M[0][11], M[1][4], M[1][11]),

		NewLineBetween(M[0][5], M[0][15], M[1][5], M[1][15]),
		NewLineBetween(M[0][5], M[0][13], M[1][5], M[1][13]),

		NewLineBetween(M[0][6], M[0][14], M[1][6], M[1][14]),
		NewLineBetween(M[0][6], M[0][15], M[1][6], M[1][15]),

		NewLineBetween(M[0][7], M[0][12], M[1][7], M[1][12]),
		NewLineBetween(M[0][7], M[0][14], M[1][7], M[1][14]),
	}
}

// Scrolled handles the zooming of the view.
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

// Dragged handles the rotation of the view.
func (l *LineDrawer) Dragged(d *fyne.DragEvent) {
	a := float64(d.Dragged.DY) * 0.007
	b := float64(d.Dragged.DX) * -0.007

	R := matrix.Matrix{
		{math.Cos(b), 0, math.Sin(b)},
		{math.Sin(a) * math.Sin(b), math.Cos(a), -math.Sin(a) * math.Cos(b)},
		{-math.Cos(a) * math.Sin(b), math.Sin(a), math.Cos(a) * math.Cos(b)},
	}

	l.matrix, _ = matrix.Mult(R, l.matrix)
	l.lines = LinesFromMatrix(l.matrix)
	l.Refresh()
}

// DragEnd is not currently needed other than  to satisfy fyne.Draggable.
func (l *LineDrawer) DragEnd() {

}

// CreateRenderer is a method that creates a renderer for the widget.
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

// NewLineBetween creates a new line between the given coordinates.
func NewLineBetween(x1, y1, x2, y2 float64) *canvas.Line {
	return &canvas.Line{
		Position1:   fyne.NewPos(float32(x1)+20, float32(x2)+20),
		Position2:   fyne.NewPos(float32(y1)+20, float32(y2)+20),
		StrokeColor: color.NRGBA{R: 0xff, G: 0xeb, B: 0x3b, A: 0xaf},
		StrokeWidth: 7,
	}
}
