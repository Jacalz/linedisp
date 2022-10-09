package main

import (
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/Jacalz/linalg/matrix"
)

var _ fyne.Widget = (*LineDrawer)(nil)
var _ fyne.Draggable = (*LineDrawer)(nil)
var _ fyne.Scrollable = (*LineDrawer)(nil)

// LineDrawer draws lines from a matrix of position vectors.
type LineDrawer struct {
	widget.BaseWidget
	lines  []fyne.CanvasObject
	matrix matrix.Matrix
}

// NewLineDrawer creates a new LineDrawer with the given matrix.
func NewLineDrawer(matrix matrix.Matrix) *LineDrawer {
	draw := &LineDrawer{
		lines:  LinesFromMatrix(matrix),
		matrix: matrix,
	}
	draw.ExtendBaseWidget(draw)
	return draw
}

// LinesFromMatrix creates new canvas.Line from the matrix.
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

// Scrolled handles the zooming of the view.
func (l *LineDrawer) Scrolled(s *fyne.ScrollEvent) {
	scale := math.Abs(float64(s.Scrolled.DY)) / 8

	// Zooming out uses scale factor  0 < scale < 1, not negative numbers.
	if s.Scrolled.DY < 0 {
		scale = 1 / scale
	}

	T := matrix.Matrix{
		{scale, 0, 0},
		{0, scale, 0},
		{0, 0, scale},
	}

	l.matrix, _ = matrix.Mult(T, l.matrix)
	l.lines = LinesFromMatrix(l.matrix)
	l.Refresh()
}

// Dragged handles the rotation of the view.
func (l *LineDrawer) Dragged(d *fyne.DragEvent) {
	dy := float64(d.Dragged.DY) / 100
	dx := -float64(d.Dragged.DX) / 100

	cosDx := math.Cos(dx)
	sinDx := math.Sin(dx)

	cosDy := math.Cos(dy)
	sinDy := math.Sin(dy)

	// Combined matrix for dragging in both x and y directions.
	R := matrix.Matrix{
		{cosDx, 0, sinDx},
		{sinDy * sinDx, cosDy, -sinDy * cosDx},
		{-cosDy * sinDx, sinDy, cosDy * cosDx},
	}

	l.matrix, _ = matrix.Mult(R, l.matrix)
	l.lines = LinesFromMatrix(l.matrix)
	l.Refresh()
}

// DragEnd is not currently needed other than to satisfy fyne.Draggable.
func (l *LineDrawer) DragEnd() {

}

// CreateRenderer is a method that creates a renderer for the widget.
func (l *LineDrawer) CreateRenderer() fyne.WidgetRenderer {
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
		Position1:   fyne.NewPos(float32(x1)+300, float32(x2)+300),
		Position2:   fyne.NewPos(float32(y1)+300, float32(y2)+300),
		StrokeColor: theme.PrimaryColor(),
		StrokeWidth: 3,
	}
}
