package gop5js

import (
	"html/template"
	"net/http"
)

const templateIndex = `<!DOCTYPE html>
<html>
  <head>
    <meta name="viewport" width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0>
    <style> body {padding: 0; margin: 0;} </style>
    <script src="/lib/p5.js"></script>
    <script src="globals.js"></script>
    <script src="sketch.js"></script>
  </head>
  <body>
  </body>
</html>
`

const templateGlobals = `
/*
* All the global JS vars starts with a capital letter!
*/
{{ range .StringVars }}
var {{ .Name }} = "{{ .Value }}";{{ end }}
{{ range .IntVars }}
var {{ .Name }} = {{ .IntValue }};{{ end }}
var images = [];
function preload(){
{{ range $key, $value := .Images}}
images['{{ $key }}'] = loadImage('/files/{{ $value }}');
{{ end }}
}
`

func globalsHandler(w http.ResponseWriter, r *http.Request) {
	globalParams := struct {
		StringVars []struct {
			Name  string
			Value string
		}
		IntVars []struct {
			Name     string
			IntValue int
		}
		Images map[string]string
	}{
		[]struct{ Name, Value string }{
			{"ServerName", serverName},
			{"ServerPort", ServerPort},
		},
		[]struct {
			Name     string
			IntValue int
		}{
			{"CanvasWidth", CanvasWidth},
			{"CanvasHeight", CanvasHeight},
		},
		images,
	}
	t := template.Must(template.New("globals").Parse(templateGlobals))
	err := t.Execute(w, globalParams)

	ErrorContainer.add(err)
}

const templateSketch = `
var socket = new WebSocket("ws://" + ServerName + ServerPort + "/ws");

var p5Data = {
  mouseX: 0,
  mouseY: 0
};

var images = {}

function setup() {
  createCanvas(CanvasWidth, CanvasHeight);
  noLoop();
}

var sketch_draw = "";

function draw() {
  clear();
  eval(sketch_draw);
  socket.send(JSON.stringify(getParams()));
}

socket.onmessage = function (event) {
  newData = JSON.parse(event.data);
  sketch_draw = newData.sketch_draw;
  draw();
}

function getParams() {
  return {
    mouse_x: mouseX,
    mouse_y: mouseY,
    p_mouse_x: pmouseX,
    p_mouse_y: pmouseY,
    win_mouse_x: winMouseX,
    win_mouse_y: winMouseY,
    p_win_mouse_x: winMouseX,
    p_win_mouse_y: winMouseY,
    mouse_button: (mouseButton === 0) ? "" : mouseButton,
    mouse_is_pressed: mouseIsPressed
  }
}
`
