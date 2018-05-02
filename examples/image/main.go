package main

import (
	"github.com/as27/gop5js"
)

func main() {
	gop5js.Draw = draw
	// set the FilesPath variable. That folder can be accessed from p5.js
	gop5js.FilesPath = "files/"
	// before you can use the image inside draw() the image needs to be
	// loaded
	gop5js.LoadImage("star", "star.png")
	gop5js.Serve()
}

var x float64 = 0

func draw() {
	gop5js.Background("127")

	x = x + 1.5
	// access the loaded image
	gop5js.Image("star", 50+x*0.5, 50, 150, 150)

}
