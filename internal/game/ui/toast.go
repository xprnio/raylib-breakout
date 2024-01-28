package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type ToastEvent string
func NewToastEvent(message string) ToastEvent {
  return ToastEvent(message)
}

type Toast struct {
  Id int64
  Message string
  Position rl.Vector2
  Size rl.Vector2
  alive float32
}

func NewToast(message string, position rl.Vector2, size rl.Vector2) *Toast {
  return &Toast{
    Message: message,
    Position: position,
    Size: size,
    alive: 5,
  }
}

func (t *Toast) Visible() bool {
  return t.alive > 0
}

func (t *Toast) Update(d float32) {
  t.alive -= d
}

func (t *Toast) Draw() {
  font := rl.GetFontDefault()
  fontSize := float32(24)
  fontSpacing := float32(2)

  size := rl.MeasureTextEx(font, t.Message, fontSize, fontSpacing)
  position := rl.Vector2Add(
    t.Position,
    rl.NewVector2(
      (t.Size.X - size.X) / 2,
      (t.Size.Y - size.Y) / 2,
    ),
  )

  rl.DrawRectangleV(t.Position, t.Size, rl.Red)
  rl.DrawTextEx(font, t.Message, position, fontSize, fontSpacing, rl.White)
}
