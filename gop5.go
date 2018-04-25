package gop5js

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

//var Setup = func() {}

// sketchDraw
var sketchDraw bytes.Buffer

// images needs to be preloaded
var images = make(map[string]string)

// Draw is called every frame
var Draw = func() {}

// CanvasWidth defines the canvas which is generated inside the browser
var CanvasWidth = 700

// CanvasHeight defines the canvas which is generated inside the browser
var CanvasHeight = 700

var Event P5Event

// SleepPerFrame makes the execution a litle slower, beccause the Go implementation
// is so extrem fast.
var SleepPerFrame = time.Millisecond * time.Duration(50)

// ErrorContainer collects all errors
var ErrorContainer errorContainer

type Data struct {
	FrameCount    int    `json:"frame_count,omitempty"`
	SketchDraw    string `json:"sketch_draw,omitempty"`
	SketchPreload string `json:"sketch_preload,omitempty"`
}

type P5Event struct {
	MouseX         int    `json:"mouse_x,omitempty"`
	MouseY         int    `json:"mouse_y,omitempty"`
	PMouseX        int    `json:"p_mouse_x,omitempty"`
	PMouseY        int    `json:"p_mouse_y,omitempty"`
	WinMouseX      int    `json:"win_mouse_x,omitempty"`
	WinMouseY      int    `json:"win_mouse_y,omitempty"`
	PWinMouseX     int    `json:"p_win_mouse_x,omitempty"`
	PWinMouseY     int    `json:"p_win_mouse_y,omitempty"`
	MouseButton    string `json:"mouse_button,omitempty"`
	MouseIsPressed bool   `json:"mouse_is_pressed,omitempty"`
}

// DrawCmd takes a JavaScript command, which is added to draw()
func DrawCmd(cmd string) {
	_, err := sketchDraw.WriteString(cmd + ";")
	ErrorContainer.add(err)
}

var data Data

func nextFrame(message []byte) {

	err := json.Unmarshal(message, &Event)
	if err != nil {
		fmt.Println(err, message, string(message))
	}
	ErrorContainer.add(err)

	data.FrameCount++
	sketchDraw = bytes.Buffer{}
	Draw()
	data.SketchDraw = sketchDraw.String()
	time.Sleep(SleepPerFrame)
	wshub.writeJSON(data)

}
