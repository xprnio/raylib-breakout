package controllers

import rl "github.com/gen2brain/raylib-go/raylib"

type MirrorController struct {
  position rl.Vector2
  target PaddleController
}

func NewMirror(target PaddleController) *MirrorController {
  return &MirrorController{
    position: target.GetPosition(),
    target: target,
  }
}

func (c *MirrorController) Update(d float32) {
  pos := c.target.GetPosition()
  c.position.X = pos.X
}

func (c *MirrorController) GetPosition() rl.Vector2 {
  return c.position
}

func (c *MirrorController) SetPosition(p rl.Vector2) {
  c.position = p
}
