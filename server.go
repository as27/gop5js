package gop5js

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gobuffalo/packr"
	"github.com/skratchdot/open-golang/open"

	"github.com/as27/golib/js/p5js"
	"github.com/gorilla/mux"
)

//go:generate packr

// PathPrefix can be used to give all the urls a prefix. The prefix
// is used for all routes of that package.
// If you set a prefix ensure that the prefix starts with a backslash "/"
// For example "/p5js" or "/myprefix"
var PathPrefix = ""

// ServerPort defines the port for the communication to the client
var ServerPort = ":2700"

// newRouter returns a gorillatoolkit router with all routes needed for
// the gop5js package
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc(PathPrefix+"/lib/p5.js", p5js.Handler)
	r.HandleFunc(PathPrefix+"/ws", wsHandleFunc)
	box := packr.NewBox("./templates")
	r.HandleFunc(PathPrefix+"/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(templateIndex))
	})
	r.HandleFunc(PathPrefix+"/sketch.js", func(w http.ResponseWriter, r *http.Request) {
		w.Write(box.Bytes("sketch.js"))
	})

	return r
}

// Serve starts the server for the communication with the browser
func Serve() error {
	r := newRouter()
	log.Println("Serving on port: ", ServerPort)
	baseURL := fmt.Sprintf("http://localhost%s%s", ServerPort, PathPrefix)
	log.Println("Open your browser and go to ", baseURL)
	open.Run(baseURL)
	return http.ListenAndServe(ServerPort, r)
}
