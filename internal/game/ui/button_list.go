package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type ButtonList struct {
  Position rl.Vector2
  Size rl.Vector2
  Spacing float32

  Buttons []Button

  FontStyle ButtonFontStyle
  Style, HoverStyle ButtonStyle
}

func NewButtonList(position rl.Vector2, size rl.Vector2, spacing float32) *ButtonList {
  return &ButtonList{
    Position: position,
    Size: size,
    Spacing: spacing,
  }
}

func (l *ButtonList) Add(text string, onClick func()) {
  button := NewButton(text, l.NextBounds(), onClick)
  l.Buttons = append(l.Buttons, button)
}

func (l *ButtonList) NextBounds() rl.Rectangle {
  buttons := len(l.Buttons)
  offset := float32(buttons) * (l.Size.Y + l.Spacing)
  return rl.NewRectangle(
    l.Position.X,
    l.Position.Y + offset,
    l.Size.X, l.Size.Y,
  )
}

func (l *ButtonList) Update(d float32) {
  for _, b := range l.Buttons {
    b.Update(d)
  }
}

func (l *ButtonList) Draw() {
  for _, b := range l.Buttons {
    b.Draw()
  }
}
