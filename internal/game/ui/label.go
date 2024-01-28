package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type Label struct {
  Text string
  Position rl.Vector2
}

type LabelOption func(l *Label)

func NewLabel(text string, options ...LabelOption) Label {
  label := Label{ Text: text }

  for _, option := range options {
    option(&label)
  }

  return label
}

func (l Label) Size() rl.Vector2 {
  font := rl.GetFontDefault()
  return rl.MeasureTextEx(font, l.Text, 32, 8)
}

func (l Label) Update(d float32) {}

func (l Label) Draw() {
  font := rl.GetFontDefault()
  rl.DrawTextEx(font, l.Text, l.Position, 32, 8, rl.White)
}
