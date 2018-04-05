package gop5js

import "fmt"

func Background(color string) {
	_, err := sketchDraw.WriteString(
		fmt.Sprintf("background(%s);", color),
	)
	ErrorContainer.add(err)
}

// Clear clears the pixels within a buffer. This function only
//  works on p5.Canvas objects created with the createCanvas()
// function; it won't work with the main display window. Unlike
// the main graphics context, pixels in additional graphics areas
// created with createGraphics() can be entirely or partially
// transparent. This function clears everything to make all of
// the pixels 100% transparent.
func Clear() {
	_, err := sketchDraw.WriteString("clear();")
	ErrorContainer.add(err)
}

func Fill(color string) {
	_, err := sketchDraw.WriteString(
		fmt.Sprintf("fill(%s);", color),
	)
	ErrorContainer.add(err)
}

func NoFill() {
	_, err := sketchDraw.WriteString("noFill();")
	ErrorContainer.add(err)
}

func NoStroke() {
	_, err := sketchDraw.WriteString("noStroke();")
	ErrorContainer.add(err)
}

func Stroke(color string) {
	_, err := sketchDraw.WriteString(
		fmt.Sprintf("stroke(%s);", color),
	)
	ErrorContainer.add(err)
}
