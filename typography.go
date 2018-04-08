package gop5js

import "fmt"

// Text draws the text onto the canvas.
// The JS Syntax is:
// text(str,x,y,[x2],[y2])
// x	Number: x-coordinate of text
// y	Number: y-coordinate of text
// x2	Number: by default, the width of the text box, see rectMode() for more info
// y2	Number: by default, the height of the text box, see rectMode() for more info
func Text(str string, xyWidth ...float64) {
	v := joinFloat64(xyWidth...)
	_, err := sketchDraw.WriteString(
		fmt.Sprintf("text('%s',%s);", str, v),
	)
	ErrorContainer.add(err)
}

// TextSize sets the current font size. This size will be used
// in all subsequent calls to the text() function. Font size is
// measured in pixels.
func TextSize(size int) {
	_, err := sketchDraw.WriteString(
		fmt.Sprintf("textSize(%d);", size),
	)
	ErrorContainer.add(err)
}
