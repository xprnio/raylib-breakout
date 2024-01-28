package controllers

import rl "github.com/gen2brain/raylib-go/raylib"

var VelocityRight = rl.NewVector2(1, 0)
var VelocityLeft = rl.NewVector2(-1, 0)

const VelocitySpeed = float32(500);

type LocalController struct {
  position rl.Vector2
}

func NewLocal() *LocalController {
  return &LocalController{
    position: rl.Vector2Zero(),
  }
}

func (c *LocalController) Update(d float32) {
  velocity := rl.Vector2Zero()

  if rl.IsKeyDown(rl.KeyD) {
    velocity = rl.Vector2Add(velocity, VelocityRight)
  }

  if rl.IsKeyDown(rl.KeyA) {
    velocity = rl.Vector2Add(velocity, VelocityLeft)
  }

  c.position = rl.Vector2Add(
    c.position,
    rl.Vector2Scale(velocity, d * VelocitySpeed),
  )
}

func (c *LocalController) SetPosition(pos rl.Vector2) {
  c.position = pos
}

func (c *LocalController) GetPosition() rl.Vector2 {
  return c.position
}
