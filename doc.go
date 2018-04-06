/*
	Package gop5js let you use the p5.js library with Go programs. Part of the
	package is a webserver, which uses websockets to connect with the browser.
	All the changes inside the Go program will be pushed to the browser.

	Mouse events are also send back to the server, so that the Go program can
	handle these events.

	How to use p5.js can be found at https://p5js.org/

	Not all functions inside Go are completely the same like in JavaScript,
	because Go is a typed language and JS not.

	To get this package work you need to define a simple function of type
	func(). This function should be set to `gpp5js.Draw`.

	The second thing is to call gop5js.Serve() at the end of the main function.

	A simple example which draws a circle would look like this:
		func draw(){
			gop5js.Ellipse(100, 100, 50, 50)
		}
		func main() {
			gop5js.Draw = draw // maps the draw Function
			gop5js.Serve()
		}

	The draw() function is the area, were you define all the elements, which
	will show up inside the web browser. That function is called every frame
	so you are also able to get some dynamics into the canvas.

	If you want to have a moving circle, the example above would look something
	like this:
		var x = float64(0)
		func draw(){
			gop5js.Ellipse(100, 100, x, 50)
			x++
		}
		func main() {
			gop5js.Draw = draw // maps the draw Function
			gop5js.Serve()
		}

	If that package had not implemented some p5.js features you can always
	use DrawCmd(). That function send direct the JavaScript to the browser.
	The ellipse function in JavaScript could be also called with 5 parameters
	  ellipse(x,y,w,h,detail)
	So if you want to set also the detail to the function you can write:
		func draw(){
			gop5js.DrawCmd("ellipse(100, 100, 50, 50, 4)")
		}
*/
package gop5js
