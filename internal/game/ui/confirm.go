package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ConfirmEvent struct {
  Message string
  CancelText, ConfirmText string
  OnCancel, OnConfirm func()
}

type Confirm struct {
  Position rl.Vector2
  Size rl.Vector2

  Message string
  OnCancel, OnConfirm func()

  CancelButton, ConfirmButton Button
}

func NewConfirm(message string, build func(c *Confirm)) Confirm {
  c := Confirm{Message: message}
  build(&c)

  return c
}

func (c Confirm) Update(d float32) {
  c.ConfirmButton.Update(d)
  c.CancelButton.Update(d)
}

func (c Confirm) Draw() {
  rl.DrawRectangleV(c.Position, c.Size, rl.Red)
  c.drawText()
  c.drawButtons()
}

func (c Confirm) drawButtons() {
  c.CancelButton.Draw()
  c.ConfirmButton.Draw()
}

func (c Confirm) drawText() {
  font := rl.GetFontDefault()
  fontSize := float32(32)
  fontSpacing := float32(2)

  size := rl.MeasureTextEx(font, c.Message, fontSize, fontSpacing)
  position := rl.NewVector2(
    c.Position.X + (c.Size.X - size.X) / 2,
    c.Position.Y + size.Y,
  )
  rl.DrawTextEx(font, c.Message, position, fontSize, fontSpacing, rl.White)
}

func (c Confirm) Visible() bool {
  return true
}
