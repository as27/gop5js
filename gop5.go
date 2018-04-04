package gop5js

import (
	"os"
	"time"
)

//var Setup = func() {}

// Draw is called every frame
var Draw = func() {}

var SleepDuration = time.Second

type Data struct {
	FrameCount int    `json:"frame_count,omitempty"`
	Shapes     Shapes `json:"shapes,omitempty"`
}

var data Data

func nextFrame() {
	data.FrameCount++
	objects = Shapes{}
	Draw()
	data.Shapes = objects
	time.Sleep(SleepDuration)
	wshub.writeJSON(data)
	if data.FrameCount == 500 {
		os.Exit(0)
	}
}
