package gop5js

import "os"

var Setup = func() {}

var Draw = func() {}

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
	wshub.writeJSON(data)
	if data.FrameCount == 500 {
		os.Exit(0)
	}
}
