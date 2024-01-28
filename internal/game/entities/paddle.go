package entities

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/xprnio/raygo/internal/game/controllers"
)

type Paddle struct {
  Controller controllers.PaddleController
  Color rl.Color
  Position rl.Vector2
  Size rl.Vector2
}

func NewPaddle(
  controller controllers.PaddleController,
  color rl.Color,
) *Paddle {
  return &Paddle{
    Controller: controller,
    Color: color,
    Position: rl.Vector2Zero(),
    Size: rl.Vector2Scale( rl.NewVector2(64, 16), 4),
  }
}

func NewLocalPaddle(color rl.Color) *Paddle {
  controller := controllers.NewLocal()

  return NewPaddle(controller, color)
}

func NewMirrorPaddle(target controllers.PaddleController, color rl.Color) *Paddle {
  controller := controllers.NewMirror(target)
  return NewPaddle(controller, color)
}

func (p *Paddle) Init() {}
func (p *Paddle) Update(d float32) {
  p.Controller.Update(d)
  p.Position = p.Controller.GetPosition()
}

func (p *Paddle) Draw() {
  rl.DrawRectangleV(p.Position, p.Size, p.Color)
}

func (p *Paddle) GetBounds() rl.Rectangle {
  return rl.NewRectangle(p.Position.X, p.Position.Y, p.Size.X, p.Size.Y)
}

func (p *Paddle) SetBounds(b rl.Rectangle) {
  pos := rl.NewVector2(b.X, b.Y)
  p.Controller.SetPosition(pos)
  p.Position = p.Controller.GetPosition()
}
