package main

import (
	"fmt"

	"github.com/as27/gop5js"
)

func main() {
	gop5js.Draw = draw
	gop5js.Serve()
}

var x float64 = 0

func draw() {
	gop5js.Ellipse(x, 20, 30, 30)
	x = x + 0.5
	gop5js.Ellipse(200, 200, 30, 30)
	fmt.Println("zwei Ellipsen")
}
