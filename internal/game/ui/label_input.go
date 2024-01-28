package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type LabelInput struct {
  Label Label
  Input *Input
}

type LabelInputOption func(li *LabelInput)

func NewLabelInput(
  l string,
  position rl.Vector2,
  width float32,
  options ...LabelInputOption,
) *LabelInput {
  label := NewLabel(l, func(l *Label) {
    l.Position = position
  })
  size := label.Size()
  input := NewInput(
    rl.NewVector2(
      position.X,
      position.Y + size.Y,
    ),
    width,
  )
  li := &LabelInput{
    Label: label,
    Input: input,
  }

  for _, option := range options {
    option(li)
  }

  return li
}

func (li *LabelInput) GetPosition() rl.Vector2 {
  x := min(li.Label.Position.X, li.Input.Position.X)
  y := min(li.Label.Position.Y, li.Input.Position.Y)
  return rl.NewVector2(x, y)
}

func (li *LabelInput) GetSize() rl.Vector2 {
  width := max(li.Label.Size().X, li.Input.GetSize().X)
  height := li.Label.Size().Y + li.Input.GetSize().Y
  return rl.NewVector2(width, height)
}

func (li *LabelInput) GetBounds() rl.Rectangle {
  p := li.GetPosition()
  s := li.GetSize()
  return rl.NewRectangle(p.X, p.Y, s.X, s.Y)
}

func (li *LabelInput) Update(d float32) {
  li.Label.Update(d)
  li.Input.Update(d)
}

func (li *LabelInput) Draw() {
  li.Label.Draw()
  li.Input.Draw()
}

func (li *LabelInput) Value() string {
  return li.Input.Value
}
