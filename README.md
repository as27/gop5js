# gop5js

Call p5.js from Go

To program something with p5.js is real fun, because it shows what happens. gop5js gives you the possibility to use p5.js together with go. You write your code in Go and you can see the results in your browser. 

## Installation

To install just use `go get`

    go get github.com/as27/gop5js

## How to use the package

* Create function of the type `func()`
* Set the gop5js.Draw to the function
* Start the gop5js server: `gop5js.Serve()`

## Example 

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