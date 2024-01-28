package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type ButtonStyle struct {
  Background rl.Color
  Foreground rl.Color
}
type ButtonFontStyle struct {
  FontSize float32
  FontSpacing float32
}

type Button struct {
  Text string
  Bounds rl.Rectangle
  OnClick func()

  FontStyle ButtonFontStyle
  Style, HoverStyle ButtonStyle
}

func NewButton(text string, bounds rl.Rectangle, onClick func()) Button {
  return Button{
    Text: text,
    Bounds: bounds,
    OnClick: onClick,
    FontStyle: ButtonFontStyle {
      FontSize: 24,
      FontSpacing: 2,
    },
    Style: ButtonStyle{
      Background: rl.White,
      Foreground: rl.Black,
    },
    HoverStyle: ButtonStyle{
      Background: rl.Green,
      Foreground: rl.Black,
    },
  }
}

func (b Button) isHovering() bool {
  m := rl.GetMousePosition()
  return rl.CheckCollisionPointRec(m, b.Bounds)
}

func (b Button) Update(d float32) {
  if !b.isHovering() {
    return
  }

  if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
    b.OnClick()
  }
}

func (b Button) Position() rl.Vector2 {
  return rl.NewVector2(b.Bounds.X, b.Bounds.Y)
}

func (b Button) Size() rl.Vector2 {
  return rl.NewVector2(b.Bounds.Width, b.Bounds.Height)
}

func (b Button) drawText() {
  style := b.buttonStyle()

  font := rl.GetFontDefault()
  size := rl.MeasureTextEx(font, b.Text, b.FontStyle.FontSize, b.FontStyle.FontSpacing)
  position := rl.Vector2Subtract(
    rl.Vector2Add(
      rl.NewVector2(b.Bounds.X, b.Bounds.Y),
      rl.Vector2Scale(b.Size(), 0.5),
    ),
    rl.Vector2Scale(size, 0.5),
  )

  rl.DrawTextEx(
    font,
    b.Text,
    position,
    b.FontStyle.FontSize,
    b.FontStyle.FontSpacing,
    style.Foreground,
  )

  // rl.DrawRectangleLinesEx(
  //   rl.NewRectangle(position.X, position.Y, size.X, size.Y),
  //   2, rl.Green,
  // )
}

func (b Button) buttonStyle() ButtonStyle {
  if b.isHovering() {
    return b.HoverStyle
  }

  return b.Style
}

func (b Button) Draw() {
  s := b.buttonStyle()
  rl.DrawRectangleRec(b.Bounds, s.Background)

  // Debug lines
  // rl.DrawRectangleLinesEx(b.Bounds, 2, rl.Red)

  b.drawText()
}
