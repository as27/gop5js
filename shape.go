package gop5js

var objects = struct{
	Ellipses []Ellipse
}

// Ellipse Draws an ellipse (oval) to the screen. An ellipse with
// equal width and height is a circle. By default, the first two
// parameters set the location, and the third and fourth parameters
// set the shape's width and height.
type Ellipse struct {
	X, Y, W, H float64 `json:"x,omitempty"`
	Detail     int     `json:"detail,omitempty"`
}
