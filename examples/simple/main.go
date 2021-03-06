package main

import (
	"github.com/as27/gop5js"
)

func main() {
	gop5js.Draw = draw
	gop5js.FilesPath = "files/"

	gop5js.Serve()
}

var x float64 = 0

func draw() {
	gop5js.Background("127")

	gop5js.DrawCmd("fill('blue')")
	gop5js.Ellipse(x, 20, 30, 30)
	x = x + 1.5

	gop5js.Ellipse(200, 200, 30, 200-x*0.01)
	gop5js.DrawCmd("fill('red')")
	if gop5js.Event.MouseIsPressed {
		gop5js.Line(200, 200, 350, 350)
		gop5js.Rect(150, 150, 210, 210)
	}

}
