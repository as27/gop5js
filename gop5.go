package gop5js

import (
	"bytes"
	"os"
	"time"
)

//var Setup = func() {}

// sketchDraw
var sketchDraw bytes.Buffer

// Draw is called every frame
var Draw = func() {}

// SleepPerFrame makes the execution a litle slower, beccause the Go implementation
// is so extrem fast.
var SleepPerFrame = time.Millisecond * time.Duration(100)

// ErrorContainer collects all errors
var ErrorContainer errorContainer

type Data struct {
	FrameCount int    `json:"frame_count,omitempty"`
	SketchDraw string `json:"sketch_draw,omitempty"`
}

func DrawCmd(cmd string) {
	_, err := sketchDraw.WriteString(cmd + ";")
	ErrorContainer.add(err)
}

var data Data

func nextFrame() {
	data.FrameCount++
	sketchDraw = bytes.Buffer{}
	Draw()
	data.SketchDraw = sketchDraw.String()
	time.Sleep(SleepPerFrame)
	wshub.writeJSON(data)
	if data.FrameCount == 500 {
		os.Exit(0)
	}
}
