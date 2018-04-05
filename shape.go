package gop5js

import (
	"fmt"
	"strings"
)

// Arc Draw an arc to the screen. If called with only
// x, y, w, h, start, and stop, the arc will be drawn and filled as
// an open pie segment. If a mode parameter is provided, the arc will
// be filled like an open semi-circle (OPEN) , a closed semi-circle (CHORD),
// or as a closed pie segment (PIE).
func Arc(x, y, w, h, start, stop float64, mode string) {
	v := joinFloat64(x, y, w, h, start, stop)
	if mode != "" {
		v = fmt.Sprintf("%s,%s", v, mode)
	}
	_, err := sketchDraw.WriteString(
		fmt.Sprintf("arc(%s);", v),
	)
	ErrorContainer.add(err)
}

// Ellipse Draws an ellipse (oval) to the screen. An ellipse with
// equal width and height is a circle. By default, the first two
// parameters set the location, and the third and fourth parameters
// set the shape's width and height.
func Ellipse(x, y, w, h float64) {
	v := joinFloat64(x, y, w, h)
	_, err := sketchDraw.WriteString(
		fmt.Sprintf("ellipse(%s);", v),
	)
	ErrorContainer.add(err)
}

// Line Draws a line (a direct path between two points) to the screen. The version of line()
// with four parameters draws the line in 2D. To color a line, use the stroke() function. A
// line cannot be filled, therefore the fill() function will not affect the color of a line.
// 2D lines are drawn with a width of one pixel by default, but this can be changed with the
// strokeWeight() function.
func Line(x1, y1, x2, y2 float64) {
	v := joinFloat64(x1, y1, x2, y2)
	_, err := sketchDraw.WriteString(
		fmt.Sprintf("line(%s);", v),
	)
	ErrorContainer.add(err)
}

// Point Draws a point, a coordinate in space at the dimension of one pixel. The first
// parameter is the horizontal value for the point, the second value is the vertical value
// for the point. The color of the point is determined by the current stroke.
func Point(x, y float64) {
	v := joinFloat64(x, y)
	_, err := sketchDraw.WriteString(
		fmt.Sprintf("point(%s);", v),
	)
	ErrorContainer.add(err)
}

// Quad Draw a quad. A quad is a quadrilateral, a four sided polygon. It is similar to a
// rectangle, but the angles between its edges are not constrained to ninety degrees. The
// first pair of parameters (x1,y1) sets the first vertex and the subsequent pairs should
// proceed clockwise or counter-clockwise around the defined shape.
func Quad(x1, y1, x2, y2, x3, y3, x4, y4 float64) {
	v := joinFloat64(x1, y1, x2, y2, x3, y3, x4, y4)
	_, err := sketchDraw.WriteString(
		fmt.Sprintf("quad(%s);", v),
	)
	ErrorContainer.add(err)
}

// Rect Draws a rectangle to the screen. A rectangle is a four-sided shape
// with every angle at ninety degrees. By default, the first two parameters
// set the location of the upper-left corner, the third sets the width, and
// the fourth sets the height. The way these parameters are interpreted, however,
// may be changed with the rectMode() function.
//
// The fifth, sixth, seventh and eighth parameters, if specified, determine corner
// radius for the top-right, top-left, lower-right and lower-left corners,
// respectively. An omitted corner radius parameter is set to the value of the previously
// specified radius value in the parameter list.
// The JavaScript documentation is:
// rect(x,y,w,h,[tl],[tr],[br],[bl])
// tl	Number: optional radius of top-left corner.
// tr	Number: optional radius of top-right corner.
// br	Number: optional radius of bottom-right corner.
// bl	Number: optional radius of bottom-left corner.
func Rect(vals ...float64) {
	v := joinFloat64(vals...)
	_, err := sketchDraw.WriteString(
		fmt.Sprintf("rect(%s);", v),
	)
	ErrorContainer.add(err)
}

// Triangle draws a triangle. A triangle is a plane created by connecting three points.
// The first two arguments specify the first point, the middle two arguments specify the
// second point, and the last two arguments specify the third point.
func Triangle(x1, y1, x2, y2, x3, y3 float64) {
	v := joinFloat64(x1, y1, x2, y2, x3, y3)
	_, err := sketchDraw.WriteString(
		fmt.Sprintf("triangle(%s);", v),
	)
	ErrorContainer.add(err)
}

func joinFloat64(fs ...float64) string {
	var ss []string
	for _, f := range fs {
		ss = append(ss, fmt.Sprintf("%f", f))
	}
	return strings.Join(ss, ",")
}
