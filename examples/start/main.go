package main

import (
	"github.com/as27/gop5js"
)

func main() {
	gop5js.CanvasHeight = 400
	gop5js.CanvasWidth = 400
	gop5js.Draw = draw
	gop5js.Serve()
}

func draw() {
	gop5js.Background("127")
	gop5js.Ellipse(20, 20, 30, 30)
	gop5js.Rect(100, 100, 20, 20)
	gop5js.Text("some text here", 20, 200)

}
