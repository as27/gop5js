package gop5js

import "fmt"

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
