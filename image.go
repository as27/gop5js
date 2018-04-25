package gop5js

import "fmt"

// LoadImage preloads the image. When you want to use a image you first
// have to load it. Then you are able to call Image() with the defined
// name
func LoadImage(imgName, imgFile string) {
	images[imgName] = imgFile
}

func Image(imgName string, x, y, w, h float64) {
	v := joinFloat64(x, y, w, h)

	_, err := sketchDraw.WriteString(
		fmt.Sprintf("image(images['%s'],%s);", imgName, v),
	)
	ErrorContainer.add(err)
}
