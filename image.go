package gop5js

import "fmt"

// LoadImage preloads the image. When you want to use a image you first
// have to load it. Then you are able to call Image() with the defined
// name. You should not call this function inside the draw() function.
// The imgFile is the file path inside your file folder. To define that
// folder you have to set FilesPath
func LoadImage(imgName, imgFile string) {
	images[imgName] = imgFile
}

// Image draws a image to the canvas. Before you can use this function
// you need to call LoadImage(), because all imges needs to be preloaded.
//   func init() {
//   	gop5js.FilesPath = "files"
//   	gop5js.LoadImage("star", "star.png")
//   }
//   func draw() {
//   	gop5js.Image("star", 10, 10, 30, 30)
//   }
func Image(imgName string, x, y, w, h float64) {
	v := joinFloat64(x, y, w, h)

	_, err := sketchDraw.WriteString(
		fmt.Sprintf("image(images['%s'],%s);", imgName, v),
	)
	ErrorContainer.add(err)
}
