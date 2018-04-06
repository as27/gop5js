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
	}
	t := template.Must(template.New("globals").Parse(templateGlobals))
	err := t.Execute(w, globalParams)

	ErrorContainer.add(err)
}
