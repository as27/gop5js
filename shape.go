package gop5js

var objects Shapes

type Shapes struct {
	Ellipses []EllipseShape
}

// EllipseShape defines the type for sending to the client
type EllipseShape struct {
	X float64 `json:"x,omitempty"`
	Y float64 `json:"y,omitempty"`
	W float64 `json:"w,omitempty"`
	H float64 `json:"h,omitempty"`
}

// Ellipse Draws an ellipse (oval) to the screen. An ellipse with
// equal width and height is a circle. By default, the first two
// parameters set the location, and the third and fourth parameters
// set the shape's width and height.
func Ellipse(x, y, w, h float64) {
	objects.Ellipses = append(objects.Ellipses, EllipseShape{x, y, w, h})
}
