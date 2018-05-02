# gop5js

Call p5.js from Go

To program something with p5.js is real fun, because it shows what happens. gop5js gives you the possibility to use p5.js together with go. You write your code in Go and you can see the results in your browser. 

## Installation

To install just use `go get`

    go get github.com/as27/gop5js

## How to use the package

* Create function of the type `func()`
* let gop5js.Draw point to the function
* Start the gop5js server: `gop5js.Serve()`

## Example 
```go
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
```

### How to include images

That p5.js can access 
```go
func main() {
	gop5js.Draw = draw
	// set the FilesPath variable. That folder can be accessed from p5.js
	gop5js.FilesPath = "files/"
	// before you can use the image inside draw() the image needs to be
	// loaded
	gop5js.LoadImage("star", "star.png")
	gop5js.Serve()
}

func draw() {
	gop5js.Background("127")
	x = x + 1.5
	// access the loaded image
	gop5js.Image("star", 50, 50, 150, 150)
}
```