package gop5js

import (
	"fmt"
	"strings"
)

// javaScripter is a internal interface which is used to
// generate the JavaScript definition for each p5.js
// element.
type javaScripter interface {
	js() string
}

// EllipseShape defines the type for sending to the client
type EllipseShape struct {
	X, Y, W, H float64
}

func (es EllipseShape) js() string {
	return fmt.Sprintf("ellipse(%f,%f,%f,%f);", es.X, es.Y, es.W, es.H)
}

// Ellipse Draws an ellipse (oval) to the screen. An ellipse with
// equal width and height is a circle. By default, the first two
// parameters set the location, and the third and fourth parameters
// set the shape's width and height.
func Ellipse(x, y, w, h float64) {
	e := EllipseShape{x, y, w, h}
	_, err := sketchDraw.WriteString(e.js())
	ErrorContainer.add(err)
}

func Rect(vals ...float64) {
	v := joinFloat64(vals...)
	_, err := sketchDraw.WriteString(
		fmt.Sprintf("rect(%s);", v),
	)
	ErrorContainer.add(err)
}

type LineShape struct {
	X1, Y1, X2, Y2 float64
}

func (ls LineShape) js() string {
	return fmt.Sprintf("line(%f,%f,%f,%f);", ls.X1, ls.Y1, ls.X2, ls.Y2)
}

func Line(x1, y1, x2, y2 float64) {
	l := LineShape{x1, y1, x2, y2}
	_, err := sketchDraw.WriteString(l.js())
	ErrorContainer.add(err)
}

func joinFloat64(fs ...float64) string {
	var ss []string
	for _, f := range fs {
		ss = append(ss, fmt.Sprintf("%f", f))
	}
	return strings.Join(ss, ",")
}
