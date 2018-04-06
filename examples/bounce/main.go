package main

import "github.com/as27/gop5js"

type Vector struct {
	x, y float64
}

func (v *Vector) add(v2 *Vector) {
	v.x = v.x + v2.x
	v.y = v.y + v2.y
}

func (v *Vector) neg() {
	v.x = -v.x
	v.y = -v.y
}

func (v *Vector) sub(v2 *Vector) {
	v2.neg()
	v.add(v2)
}

type Mover struct {
	location     *Vector
	velocity     *Vector
	acceleration *Vector
}

func newMover(x, y float64) Mover {
	m := Mover{}
	m.location = &Vector{x, y}
	m.velocity = &Vector{0, 0}
	m.acceleration = &Vector{0, 0}
	return m
}

func (m Mover) update() {
	m.velocity.add(m.acceleration)
	m.location.add(m.velocity)
	if m.location.x <= 25 || m.location.x >= 675 {
		m.velocity.x = -m.velocity.x
	}
	if m.location.y <= 25 || m.location.y >= 575 {
		m.velocity.y = -m.velocity.y
	}
	gop5js.Ellipse(m.location.x, m.location.y, 50, 50)
}

var m1 = newMover(float64(30), float64(30))

func init() {
	wind := &Vector{0.9, 0.01}
	gravity := &Vector{0, 1}
	m1.acceleration.add(wind)
	m1.acceleration.add(gravity)
}

func draw() {
	gop5js.Background("77")
	m1.update()
}

func main() {
	gop5js.Draw = draw
	gop5js.CanvasHeight = 600
	gop5js.CanvasWidth = 700
	gop5js.Serve()
}
