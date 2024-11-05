package main

import (
	"math"

	"grow.graphics/gd"
	"grow.graphics/gd/gdextension"
)

type Player struct {
  gd.Class[Player, gd.Sprite2D]
  Speed float64
  AngularSpeed float64
}

func (p *Player) Ready() {
	println("Hello, world!")
	p.Speed = 400
	p.AngularSpeed = math.Pi
	
}

func (p *Player) PhysicsProcess(delta float64) {
	var radians = p.AngularSpeed * delta
	p.Super().AsNode2D().Rotate(radians)
}

/*
func (h *HelloName) Ready() {
  tmp := h.Temporary // temporary lifetime for new Godot values.
  h.Button.AsObject().Connect(tmp.StringName("pressed"), tmp.Callable(h.OnButtonPressed), 0)
}

func (h *HelloName) OnButtonPressed() {
  tmp := h.Temporary // temporary lifetime for new Godot values.
  h.Text.SetText(tmp.String("Hello " + h.Name.GetText(tmp).String()))
}
*/
func main() {
  godot, ok := gdextension.Link()
  if !ok {
    return
  }
  gd.Register[Player](godot)
}