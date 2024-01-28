package controllers

import rl "github.com/gen2brain/raylib-go/raylib"

type PaddleController interface {
  Update(d float32)
  SetPosition(pos rl.Vector2)
  GetPosition() rl.Vector2
}
